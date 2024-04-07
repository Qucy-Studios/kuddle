package kuddle

import (
	"encoding/base64"
	"errors"
	"github.com/ShindouMihou/siopao/siopao"
	"mime"
	"path/filepath"
	"strings"
)
import "github.com/dchest/uniuri"
import "github.com/imroc/req/v3"

type ImageDownloadResult struct {
	ContentType string
	ImagePath   string
	Extension   string
}

func (app *App) Download(link string) (*ImageDownloadResult, error) {
	tempImagePath := filepath.Join(GetKuddleDir(), ".temp", uniuri.NewLen(128))
	response, err := req.Get(link)
	if err != nil {
		return nil, err
	}
	contentType := strings.ToLower(response.GetContentType())
	if !strings.HasPrefix(contentType, "image/") {
		return nil, errors.New("link must redirect to an image")
	}
	file := siopao.Open(tempImagePath)
	if err := file.Write(response.Body); err != nil {
		return nil, err
	}

	exts, _ := mime.ExtensionsByType(response.GetContentType())
	ext := ""
	if len(exts) > 0 {
		ext = exts[0]
		if ext == ".jpe" {
			ext = ".jpeg"
		}
	}

	return &ImageDownloadResult{
		ContentType: response.GetContentType(),
		ImagePath:   tempImagePath,
		Extension:   ext,
	}, nil
}

func (app *App) DownloadBase64(b64 string) (string, error) {
	tempImagePath := filepath.Join(GetKuddleDir(), ".temp", uniuri.NewLen(128))
	b64 = strings.SplitN(b64, ",", 2)[1]
	dec, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "", err
	}
	file := siopao.Open(tempImagePath)
	if err := file.Write(dec); err != nil {
		return "", err
	}
	return tempImagePath, nil
}

func (app *App) CleanUp() {
	tempImagesPath := filepath.Join(GetKuddleDir(), ".temp")
	_ = siopao.Open(tempImagesPath).DeleteRecursively()
}
