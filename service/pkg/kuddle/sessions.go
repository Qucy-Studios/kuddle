package kuddle

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ShindouMihou/go-little-utils/slices"
	"github.com/ShindouMihou/siopao/siopao"
	"github.com/dchest/uniuri"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/portainer/libcrypto"
	"path/filepath"
	"strings"
	"sync"
)

type Session struct {
	Name            string `json:"name"`
	AccessKeyId     string `json:"access_key_id"`
	Endpoint        string `json:"endpoint"`
	SecretAccessKey string `json:"secret_access_key"`
	Token           string `json:"token"`
	UseSSL          bool   `json:"use_ssl"`
	Bucket          string `json:"bucket"`
	Domain          string `json:"domain"`
	client          *minio.Client
}

type EphemeralSettings struct {
	UrlLength uint8
	Folder    string
}

func (app *App) CreateSession(encryptionKey string, session Session) error {
	keyAsByte := []byte(encryptionKey)

	session.client = nil
	sessionAsJson, err := json.Marshal(session)
	if err != nil {
		return err
	}
	encrypted, err := libcrypto.Encrypt(sessionAsJson, keyAsByte)
	if err != nil {
		return err
	}
	sessionPath := filepath.Join(GetKuddleDir(), "sessions", session.Name+".kddles")
	if err := siopao.Open(sessionPath).Write(encrypted); err != nil {
		return err
	}
	return nil
}

func (app *App) EditSession(encryptionKey string, originalSessionName string, session Session) error {
	// create new session first, to be sure that the edit happened.
	if err := app.CreateSession(encryptionKey, session); err != nil {
		return err
	}

	if session.Name != originalSessionName {
		// Delete original session if the name has also changed.
		// This happens last to ensure that we do not delete anything before the updated session is made.
		sessionPath := filepath.Join(GetKuddleDir(), "sessions", originalSessionName+".kddles")
		_ = siopao.Open(sessionPath).Delete()
	}

	return nil
}

func (app *App) GetSession(name string) *Session {
	if len(app.sessions) < 1 {
		panic("no sessions found")
	}
	session, ok := app.sessions[name]
	if !ok {
		return nil
	}
	return &session
}

func (app *App) Test(session Session) (bool, error) {
	err := session.Initialize()
	if err != nil {
		return false, err
	}
	for obj := range session.client.ListObjects(context.TODO(), session.Bucket, minio.ListObjectsOptions{MaxKeys: 1}) {
		println(obj.Key)
		return true, nil
	}
	return true, nil
}

func (session *Session) Initialize() error {
	endpoint := session.Endpoint
	endpoint = strings.TrimPrefix(endpoint, "https://")
	endpoint = strings.TrimPrefix(endpoint, "http://")
	endpoint = strings.TrimPrefix(endpoint, "www.")
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(session.AccessKeyId, session.SecretAccessKey, session.Token),
		Secure: session.UseSSL,
	})
	if err != nil {
		return err
	}

	session.client = client
	return nil
}

func (app *App) Upload(session Session, path string, ext string, contentType string, settings EphemeralSettings) (string, error) {
	if err := session.Initialize(); err != nil {
		return "", err
	}
	file := uniuri.NewLen(int(settings.UrlLength))
	if ext != "" {
		file += "."
		file += ext
	}
	if settings.Folder != "" {
		file = filepath.Join(settings.Folder, file)
	}
	_, err := session.client.FPutObject(context.TODO(), session.Bucket, file, path, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", err
	}
	return session.Domain + "/" + file, nil
}

func (app *App) LocateSessions(key string) ([]string, error) {
	keyAsByte := []byte(key)
	sessionsPath := filepath.Join(GetKuddleDir(), "sessions")

	var wg sync.WaitGroup
	var sessions []Session

	dir := siopao.Open(sessionsPath)
	_ = dir.MkdirParent()
	if err := dir.Recurse(false, func(file *siopao.File) {
		if isDir, err := file.IsDir(); isDir || err != nil {
			return
		}

		if filepath.Ext(file.Path()) != ".kddles" {
			return
		}

		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()

			bytes, err := file.Bytes()
			if err != nil {
				fmt.Println("sessions: failed to read session ", file.Path(), " for  ", err)
				return
			}
			sessionRaw, err := libcrypto.Decrypt(bytes, keyAsByte)
			if err != nil {
				fmt.Println("sessions: invalid decryption key for session ", file.Path(), " with error ", err)
				return
			}
			var session Session
			if err := json.Unmarshal(sessionRaw, &session); err != nil {
				fmt.Println("sessions: failed to unmarshal session ", file.Path(), " with error ", err)
				return
			}

			sessions = append(sessions, session)
		}()
	}); err != nil {
		return nil, err
	}
	wg.Wait()
	app.sessions = make(map[string]Session)
	for _, session := range sessions {
		app.sessions[session.Name] = session
	}
	return slices.Map(sessions, func(v Session) string {
		return v.Name
	}), nil
}

func (app *App) DeleteSession(sessionName string) {
	sessionPath := filepath.Join(GetKuddleDir(), "sessions", sessionName+".kddles")
	_ = siopao.Open(sessionPath).Delete()
}
