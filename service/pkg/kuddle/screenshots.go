package kuddle

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"golang.design/x/clipboard"
	"net/http"
)

func (app *App) RequestScreenshot() string {
	image := clipboard.Read(clipboard.FmtImage)
	if image == nil {
		return ""
	}
	return imgToBase64(*bytes.NewBuffer(image))
}

// img to b64 cc: https://github.com/polds/imgbase64/blob/master/images.go

func encode(bin []byte) []byte {
	e64 := base64.StdEncoding

	maxEncLen := e64.EncodedLen(len(bin))
	encBuf := make([]byte, maxEncLen)

	e64.Encode(encBuf, bin)
	return encBuf
}

func format(enc []byte, mime string) string {
	switch mime {
	case "image/gif", "image/jpeg", "image/pjpeg", "image/png", "image/tiff":
		return fmt.Sprintf("data:%s;base64,%s", mime, enc)
	default:
	}

	return fmt.Sprintf("data:image/png;base64,%s", enc)
}

func imgToBase64(buf bytes.Buffer) string {
	enc := encode(buf.Bytes())
	mime := http.DetectContentType(buf.Bytes())

	return format(enc, mime)
}
