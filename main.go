package main

import (
	"flag"
	"image/gif"
	"image/png"
	"os"
)

func main() {
	var gifFile = flag.String("gif", "", "gif input file")
	var pngFile = flag.String("png", "", "png output file")
	flag.Parse()

	reader, err := os.Open(*gifFile)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	img, err := gif.Decode(reader)
	if err != nil {
		panic(err)
	}

	writer, err := os.Create(*pngFile)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	err = png.Encode(writer, img)
	if err != nil {
		panic(err)
	}
}
