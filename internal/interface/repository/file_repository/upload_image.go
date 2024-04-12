package file_repository

import (
	"errors"
	"github.com/gofrs/uuid"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
)

func contains(arr []string, item string) bool {
	for _, el := range arr {
		if el == item {
			return true
		}
	}

	return false
}

func buildFileName() string {
	id, _ := uuid.NewV4()
	return strings.ReplaceAll(id.String(), "-", "_")
}

func uploadImage(file multipart.File, directory string) (string, error) {
	fileHeader := make([]byte, 512)
	_, err := file.Read(fileHeader)
	if err != nil {
		return "", err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return "", err
	}

	mime := http.DetectContentType(fileHeader)

	err = os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return "", err
	}

	fname := path.Join(directory, buildFileName()) + ".jpg"
	if !contains([]string{"image/jpeg", "image/png", "image/bmp", "image/gif"}, mime) {
		return "", errors.New("invalid mime type " + mime)
	}

	imgSrc, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	newImg := image.NewRGBA(imgSrc.Bounds())
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), imgSrc, imgSrc.Bounds().Min, draw.Over)
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}

	var opt jpeg.Options
	opt.Quality = 80

	err = jpeg.Encode(f, newImg, &opt)
	if err != nil {
		return "", err
	}

	defer f.Close()
	return "/" + fname, nil
}
