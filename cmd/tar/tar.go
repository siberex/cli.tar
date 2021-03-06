package main

import (
	"archive/tar"
	"flag"
	"io"
	"log"
	"os"
	"path"
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
		// Skip dir handling, Bazel should not produce such inputs.
		// Also, CLI intentionally should produce flat tarballs (same as rules_pkg)
		return
	}

	absPath, err := filepath.Abs(relPath)
	if err != nil {
		log.Fatal(err)
	}

	destPath := replacePath + fStat.Name()

	hdr := &tar.Header{
		Typeflag: tar.TypeReg,
		Name:     destPath,
		Size:     fStat.Size(),
		Mode:     0555,
		ModTime:  PortableMtime,
	}

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

func writeDirHeader(tw *tar.Writer, path string) error {
	if path == "" {
		return nil
	}

	return tw.WriteHeader(&tar.Header{
		Typeflag: tar.TypeDir,
		Name:     path,
		Size:     0,
		Mode:     0755,
		ModTime:  PortableMtime,
	})
}

func main() {
	directoryArg := flag.String("dir", "", "Target directory")
	flag.Parse()

	args := flag.Args()
	if len(args) < 3 {
		log.Fatal("tar: no files specified\nUsage:\n  tar [--dir /target/path] output.tar input1 [input2 ...]")
	}

	outPath, _ := filepath.Abs(args[0])

	extension := filepath.Ext(outPath)
	if strings.ToLower(extension) != ".tar" {
		outPath += ".tar"
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

	replacePath := "./"
	if *directoryArg != "" {
		replacePath += strings.Trim(
			path.Clean(*directoryArg),
			"./",
		)
		if replacePath[len(replacePath)-1] != '/' {
			replacePath += "/"
		}
	}

	// Emulate mkdir -p /very/long/path for tarball target directory with tar headers
	subdirs := strings.SplitAfter(replacePath, "/")
	subdirs = subdirs[:len(subdirs)-1]
	incPath := ""
	for _, subdir := range subdirs {
		incPath += subdir
		err := writeDirHeader(tw, incPath)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Interpret arguments as input files to archive
	for _, relPath := range args[1:] {
		addToArchive(tw, relPath, replacePath)
	}

}
