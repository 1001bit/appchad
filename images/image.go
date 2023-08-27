package images

import (
	"io"
	"net/http"
	"os"
	"strings"
)

var allFileExt = []string{"jpeg", "jpg", "png"}

const filePath = "/images/files/"

// load image from request to a folder
func ImageUpload(r *http.Request, id string) error {
	// ParseMultipartForm parses a request body as multipart/form-data
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return err
	}

	// get file from request
	file, handler, err := r.FormFile("image")

	if err != nil {
		return err
	}
	defer file.Close()

	// file extension (no non-image)
	fileExt := append(strings.Split(handler.Filename, "."), "")[1]
	var compatible bool = false

	for _, e := range allFileExt {
		if e == fileExt {
			compatible = true
			break
		}
	}

	if !compatible {
		return nil
	}

	// create file on path
	f, err := os.OpenFile(filePath[1:]+id, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return err
	}
	defer f.Close()

	// Copy the request file to a new created file in folder
	io.Copy(f, file)

	return nil
}

func ImageDelete(id string) error {
	err := os.Remove(filePath[1:] + id)
	if err != nil {
		return err
	}
	return nil
}
