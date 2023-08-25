package images

import (
	"io"
	"net/http"
	"os"
	"strings"
)

var allFileExt = []string{"jpeg", "jpg", "png"}

const filePath = "/images/files/"

func ImageUpload(r *http.Request, id string) error {
	// ParseMultipartForm parses a request body as multipart/form-data
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return err
	}

	file, handler, err := r.FormFile("image") // Retrieve the file from form data

	if err != nil {
		return err
	}
	defer file.Close() // Close the file when we finish

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

	// This is path which we want to store the file
	f, err := os.OpenFile(filePath[1:]+id, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return err
	}
	defer f.Close()

	// Copy the file to the destination path
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
