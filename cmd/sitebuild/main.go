package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	if err := os.RemoveAll("public"); err != nil {
		fail("delete public: %v", err)
	}
	if err := os.MkdirAll("public", 0755); err != nil {
		fail("create public: %v", err)
	}
	if err := copyFile("templates/index.html", "public/index.html"); err != nil {
		fail("copy index.html: %v", err)
	}
	if err := copyFile("templates/404.html", "public/404.html"); err != nil {
		fail("copy 404.html: %v", err)
	}
	if err := copyDir("static", "public/assets"); err != nil {
		fail("copy static: %v", err)
	}
	fmt.Println("Built public/")
}

func fail(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
	os.Exit(1)
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		target := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			return err
		}
		return copyFile(path, target)
	})
}
