package main

import (
	"io"
	"os"
	"path/filepath"
)

func writeTextFile(
	target string,
	content string,
) error {
	if err := os.MkdirAll(
		filepath.Dir(target),
		0o755,
	); err != nil {
		return err
	}

	return os.WriteFile(
		target,
		[]byte(content),
		0o644,
	)
}

func recreateDir(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}

	return os.MkdirAll(
		path,
		0o755,
	)
}

func copyDir(
	src string,
	dst string,
) error {
	return filepath.Walk(
		src,
		func(
			path string,
			info os.FileInfo,
			err error,
		) error {

			if err != nil {
				return err
			}

			rel, err :=
				filepath.Rel(
					src,
					path,
				)

			if err != nil {
				return err
			}

			target :=
				filepath.Join(
					dst,
					rel,
				)

			if info.IsDir() {
				return os.MkdirAll(
					target,
					info.Mode(),
				)
			}

			return copyFile(
				path,
				target,
			)
		},
	)
}

func copyFile(
	src string,
	dst string,
) error {
	in, err :=
		os.Open(src)

	if err != nil {
		return err
	}

	defer in.Close()

	if err :=
		os.MkdirAll(
			filepath.Dir(dst),
			0o755,
		); err != nil {

		return err
	}

	out, err :=
		os.Create(dst)

	if err != nil {
		return err
	}

	defer func() {
		_ = out.Close()
	}()

	if _, err :=
		io.Copy(
			out,
			in,
		); err != nil {

		return err
	}

	return out.Close()
}
