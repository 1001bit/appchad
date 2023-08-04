package blogchad

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/McCooll75/appchad/crypt"
)

var allFileExt = []string{"jpeg", "jpg", "png"}

func imageUpload(r *http.Request) (string, error) {
	// ParseMultipartForm parses a request body as multipart/form-data
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return "", err
	}

	file, handler, err := r.FormFile("image") // Retrieve the file from form data

	if err != nil {
		return "", err
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
		return "", nil
	}

	hex, err := crypt.RandomHex(8)
	if err != nil {
		return "", err
	}

	newName := "/assets/files/" + hex + "." + fileExt

	// This is path which we want to store the file
	f, err := os.OpenFile(newName[1:], os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return "", err
	}

	// Copy the file to the destination path
	io.Copy(f, file)

	return newName, nil
}
