package kuddle

import (
	"bytes"
	"errors"
	"github.com/ShindouMihou/siopao/siopao"
	"github.com/portainer/libcrypto"
	"os"
	"path/filepath"
)

func (app *App) Register(key string) error {
	keyAsByte := []byte(key)

	authFilePath := filepath.Join(GetKuddleDir(), ".store", ".auth")
	authFile := siopao.Open(authFilePath)

	hasRegistered := app.HasRegistered()
	if hasRegistered {
		return errors.New("cannot register a new key")

	}

	encrypted, err := libcrypto.Encrypt([]byte("true"), keyAsByte)
	if err != nil {
		return err
	}
	if err := authFile.Write(encrypted); err != nil {
		return err
	}

	return nil
}

func (app *App) HasRegistered() bool {
	authFilePath := filepath.Join(GetKuddleDir(), ".store", ".auth")
	if _, err := os.Open(authFilePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	} else {
		return true
	}
}

func (app *App) Authenticate(key string) (bool, error) {
	keyAsByte := []byte(key)

	authFilePath := filepath.Join(GetKuddleDir(), ".store", ".auth")
	authFile := siopao.Open(authFilePath)

	hasRegistered := app.HasRegistered()
	if !hasRegistered {
		return false, errors.New("no key has been registered")
	}

	fileBytes, err := authFile.Bytes()
	if err != nil {
		return false, err
	}
	result, err := libcrypto.Decrypt(fileBytes, keyAsByte)
	if err != nil {
		if err.Error() == "cipher: message authentication failed" {
			return false, nil
		}
		return false, err
	}
	return bytes.EqualFold(result, []byte("true")), nil
}
