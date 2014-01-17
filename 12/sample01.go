package main

import (
	"archive/tar"
	"compress/gzip"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	NoCompression()
	BestSpeed()
	BestCompression()
	DefaultCompression()
}

func NoCompression() {
	compress("output/NoCompression.tar.gz", gzip.NoCompression)
}

func BestSpeed() {
	compress("output/BestSpeed.tar.gz", gzip.BestSpeed)
}

func BestCompression() {
	compress("output/BestCompression.tar.gz", gzip.BestCompression)
}

func DefaultCompression() {
	compress("output/DefaultCompression.tar.gz", gzip.DefaultCompression)
}

func compress(outfile string, level int) {
	var file *os.File
	var err error
	var writer *gzip.Writer
	var body []byte

	if file, err = os.Create(outfile); err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	if writer, err = gzip.NewWriterLevel(file, level); err != nil {
		log.Fatalln(err)
	}
	defer writer.Close()

	tw := tar.NewWriter(writer)
	defer tw.Close()

	root := "c:/golang/go/src/pkg"

	err = filepath.Walk(root,
		func(path string, fi os.FileInfo, er error) error {

			if fi.IsDir() {
				return nil
			}

			if body, err = ioutil.ReadFile(path); err != nil {
				log.Fatalln(err)
			}

			rel, err := filepath.Rel(root, path)

			if err != nil {
				log.Fatalln(err)
			}

			if body != nil {
				hdr := &tar.Header{
					Name: rel,
					Size: int64(len(body)),
				}
				if err := tw.WriteHeader(hdr); err != nil {
					println(err)
				}
				if _, err := tw.Write(body); err != nil {
					println(err)
				}
			}

			return nil
		})
}
