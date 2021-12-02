package main

import (
	"archive/tar"
	"bytes"
	"log"
	"os"
	"path/filepath"
)

func addToArchive(tw *tar.Writer, outFile *os.File, relPath string) {
	fStat, err := os.Stat(relPath)
	if err != nil {
		log.Fatal(err)
	}

	if fStat.IsDir() {
		// FIXME: dir handling?
	}

	absPath, err := filepath.Abs(relPath)
	if err != nil {
		log.Fatal(err)
	}

	hdr, err := tar.FileInfoHeader(fStat, absPath)
	if err != nil {
		log.Fatal(err)
	}

	if err = tw.WriteHeader(hdr); err != nil {
		log.Fatal(err)
	}

	//FIXME tw.Write()

}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("tar: no files specified\nUsage:\n  tar output.tar input1 [input2 ...]")
	}

	outPath, _ := filepath.Abs(os.Args[1])

	extension := filepath.Ext(outPath)
	if extension != ".tar" {
		log.Fatalf("tar: only .tar output supported, %s provided", os.Args[1])
	}

	outFile, err := os.Create(outPath)
	if err != nil {
		log.Fatalf("tar: could not write to the output file %s", outPath)
	}

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	defer func(tw *tar.Writer) {
		_ = tw.Close()
	}(tw)

	// Interpret arguments as input files to archive
	for _, relPath := range os.Args[2:] {
		addToArchive(tw, outFile, relPath)
	}

}
