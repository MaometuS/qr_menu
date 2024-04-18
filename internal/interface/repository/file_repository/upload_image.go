package file_repository

import (
	"errors"
	"github.com/gofrs/uuid"
	"github.com/nfnt/resize"
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

func isFileImage(file multipart.File) error {
	fileHeader := make([]byte, 512)
	_, err := file.Read(fileHeader)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	mime := http.DetectContentType(fileHeader)
	if !contains([]string{"image/jpeg", "image/png", "image/bmp", "image/gif"}, mime) {
		return errors.New("invalid mime type " + mime)
	}

	return nil
}

func fileToImage(file multipart.File) (image.Image, error) {
	imgSrc, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	newImg := image.NewRGBA(imgSrc.Bounds())
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), imgSrc, imgSrc.Bounds().Min, draw.Over)
	resImg := resize.Thumbnail(1920, 1080, newImg, resize.Lanczos3)

	return resImg, nil
}

func writeImageToFile(filename, directory string, img image.Image) error {
	err := os.MkdirAll(directory, 0700)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	var opt jpeg.Options
	opt.Quality = 80

	err = jpeg.Encode(f, img, &opt)
	if err != nil {
		return err
	}

	defer f.Close()

	return nil
}

func uploadImage(file multipart.File, directory string) (string, error) {
	err := isFileImage(file)
	if err != nil {
		return "", err
	}

	name := buildFileName() + ".jpg"

	fname := path.Join(directory, name)

	img, err := fileToImage(file)
	if err != nil {
		return "", err
	}

	err = writeImageToFile(fname, directory, img)
	if err != nil {
		return "", err
	}

	return name, nil
}
