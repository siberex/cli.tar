package main

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var PortableMtime = time.Unix(946684800, 0) // 2000-01-01 00:00:00.000 UTC

func addToArchive(tw *tar.Writer, relPath string, replacePath string) {
	fStat, err := os.Stat(relPath)
	if err != nil {
		log.Fatal(err)
	}

	if fStat.IsDir() {
		// TODO: dir handling?
	}

	absPath, err := filepath.Abs(relPath)
	if err != nil {
		log.Fatal(err)
	}

	hdr, err := tar.FileInfoHeader(fStat, absPath)
	if err != nil {
		log.Fatal(err)
	}

	//destPath := path.Clean(relPath)
	destPath := fStat.Name()

	if replacePath != "" {
		destPath = "./" + strings.TrimLeft(replacePath, "/") + "/" + destPath
	}
	hdr.Name = destPath

	// Portability
	hdr.Uid = 0
	hdr.Uname = "0"
	hdr.Gid = 0
	hdr.Gname = "0"
	hdr.ModTime = PortableMtime

	srcFile, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(srcFile *os.File) {
		_ = srcFile.Close()
	}(srcFile)

	if err = tw.WriteHeader(hdr); err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(tw, srcFile)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("tar: no files specified\nUsage:\n  tar output.tar input1 [input2 ...]")
	}

	// FIXME --directory arg

	outPath, _ := filepath.Abs(os.Args[1])

	extension := filepath.Ext(outPath)
	if extension != ".tar" {
		log.Fatalf("tar: only .tar output supported, %s provided", os.Args[1])
	}

	outFile, err := os.Create(outPath)
	if err != nil {
		log.Fatalf("tar: could not write to the output file %s", outPath)
	}
	defer func(outFile *os.File) {
		_ = outFile.Close()
	}(outFile)

	tw := tar.NewWriter(outFile)
	defer func(tw *tar.Writer) {
		_ = tw.Close()
	}(tw)

	// Interpret arguments as input files to archive
	for _, relPath := range os.Args[2:] {
		addToArchive(tw, relPath, "")
	}

}
