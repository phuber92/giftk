package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/phuber92/giftk/internal/config"
)

func main() {
	config := config.ParseConfig()

	var fileGlob = &os.Args[1]

	jpegFiles, err := filepath.Glob(*fileGlob)
	if err != nil {
		panic(err)
	}

	decodedJpegImages := []image.Image{}
	for _, jpegFile := range jpegFiles {
		fmt.Printf("Decoding jpeg: %s", jpegFile)
		reader, err := os.Open(jpegFile)
		if err != nil {
			panic(err)
		}
		defer reader.Close()
		img, err := jpeg.Decode(reader)
		if err != nil {
			panic(err)
		}
		decodedJpegImages = append(decodedJpegImages, img)
	}

	writer, err := os.OpenFile(config.OutputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	gifOptions := gif.Options{
		NumColors: 256,
	}

	for _, img := range decodedJpegImages {
		gif.Encode(writer, img, &gifOptions)
	}
}
