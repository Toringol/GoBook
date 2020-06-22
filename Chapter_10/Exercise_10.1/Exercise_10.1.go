/*
* Расширьте программу jpeg так, чтобы она преобразовывала
* любой поддерживаемый входной формат в любой выходной
* с использованием функции image.Decode для определения
* входного формата и флага командной строки для выбора
* выходного формата.
 */

package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var (
	outputFormat = flag.String("outputFormat", "png", "value for output format convertion")
)

func main() {
	flag.Parse()
	if err := convert(os.Stdin, os.Stdout, *outputFormat); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func convert(in io.Reader, out io.Writer, outputFormat string) error {
	img, _, err := image.Decode(in)
	if err != nil {
		return err
	}
	switch outputFormat {
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, nil)
	case "jpeg", "jpg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	default:
		return fmt.Errorf("unknown format: %s", outputFormat)
	}
}
