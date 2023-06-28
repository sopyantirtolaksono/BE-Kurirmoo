package storage

import (
	"bytes"
	"errors"
	"kurirmoo"
	"net/http"
)

var (
	oneByte = 1048576
	oneMb   = (oneByte * 1)
)

func validateUploadImage(rt *kurirmoo.Runtime, bytesClientFile []byte) error {
	mimeType := http.DetectContentType(bytesClientFile)
	if _, ok := MimetypeImageAllowed[mimeType]; !ok {
		err := rt.SetError(http.StatusUnsupportedMediaType, errors.New("unsupported file format image, should be one of [png, jpg, jpeg]").Error())
		return err
	}

	buff := bytes.NewBuffer(bytesClientFile)
	if buff.Len() > (oneMb * 2) {
		err := rt.SetError(http.StatusUnprocessableEntity, errors.New("file size is too large than 2 MB").Error())
		return err
	}

	return nil
}
