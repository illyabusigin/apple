package xcassets

import (
	"fmt"
	"image"
	_ "image/jpeg" // support for JPG images
	_ "image/png"  // support for PNG images
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"path/filepath"

	_ "golang.org/x/image/bmp" // support for BMP images

	"os"
)

type assetLoader struct {
	source AssetSource

	width, height, maxScaleFactor int
}

func (l *assetLoader) Key(source AssetSource) string {
	if path := source.file; path != "" {
		return path
	}

	if url := source.url; url != "" {
		return url
	}

	return ""
}

func (l *assetLoader) Load(source AssetSource) (image.Image, error) {
	if path := source.file; path != "" {
		img, err := l.loadImageFromFile(path)
		return img, err
	}

	if url := l.source.url; url != "" {
		img, err := l.loadImageFromURL(url)
		return img, err
	}

	return nil, fmt.Errorf("No image source specified")
}

func (l *assetLoader) Validate() error {
	if path := l.source.file; path != "" {
		return l.validateFile(path)
	}

	if url := l.source.url; url != "" {
		return l.validateURL(url)
	}

	return nil
}

func (l *assetLoader) loadImageFromFile(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "png"
	imageData, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode image: %w", err)
	}

	return imageData, nil
}

func (l *assetLoader) loadImageFromURL(url string) (image.Image, error) {
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	fileName := path.Base(url)
	path := fmt.Sprintf("*_%v", fileName)

	file, err := ioutil.TempFile("", path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return nil, err
	}

	return l.loadImageFromFile(path)
}

func (l *assetLoader) validateURL(url string) error {
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	fileName := path.Base(url)

	file, err := ioutil.TempFile("", fmt.Sprintf("*_%v", fileName))
	if err != nil {
		return err
	}

	defer os.Remove(file.Name())
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	if err = l.validateFile(file.Name()); err != nil {
		return err
	}

	return nil
}

func (l *assetLoader) validateFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()
	image, _, err := image.DecodeConfig(file)
	if err != nil {
		return err
	}

	return l.source.validateImage(image, filepath.Base(file.Name()))
}
