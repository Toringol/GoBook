/*
* Определите обобщенную функцию для чтения архива, способную
* читать ZIP-файлы (archive/zip) и POSIX tar-файлы (archive/tar).
* Воспользуйтесь механизмом регистрации, аналогичным описанному выше,
* чтобы поддержка каждого формата файла могла быть добавлена с помощью
* пустого импорта.
 */

package main

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

type FileHeader struct {
	Name string
	Size uint64
}

var (
	archiveFile = flag.String("archiveFile", "", "archiveFile what you want to read")
)

func main() {
	flag.Parse()

	file, err := os.Open(*archiveFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var headers []FileHeader

	switch {
	case strings.HasSuffix(*archiveFile, ".tar"):
		headers, err = ListTar(file)
		if err != nil {
			log.Fatal(err)
		}
	case strings.HasSuffix(*archiveFile, ".zip"):
		headers, err = ListZip(file)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("No such format")
	}

	fmt.Println(sprintFileHeaders(headers))
}

func ListTar(f *os.File) ([]FileHeader, error) {
	var headers []FileHeader

	// Open the tar archive for reading.
	tr := tar.NewReader(f)

	// Iterate through the files in the archive.
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}
		if err != nil {
			return nil, err
		}
		headers = append(headers, FileHeader{
			Name: hdr.Name,
			Size: uint64(hdr.Size),
		})
	}
	return headers, nil
}

func ListZip(f *os.File) ([]FileHeader, error) {
	var headers []FileHeader

	// Open a zip archive for reading.
	r, err := zip.OpenReader(f.Name())
	if err != nil {
		return nil, err
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		headers = append(headers, FileHeader{
			Name: f.Name,
			Size: f.UncompressedSize64,
		})
	}
	return headers, nil
}

func sprintFileHeaders(headers []FileHeader) string {
	var b bytes.Buffer
	namelen, sizelen := longestLength(headers)
	for _, header := range headers {
		fmt.Fprintf(&b, "% -*s %*d\n", namelen, header.Name, sizelen, header.Size)
	}
	return b.String()
}

func longestLength(headers []FileHeader) (name int, size int) {
	for _, header := range headers {
		n := utf8.RuneCountInString(header.Name)
		if name < n {
			name = n
		}
		s := len(fmt.Sprintf("%d", header.Size))
		if size < s {
			size = s
		}
	}
	return
}
