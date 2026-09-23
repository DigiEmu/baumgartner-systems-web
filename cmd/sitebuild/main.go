package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	publicDir = "public"
	staticDir = "static"
	siteURL   = "https://baumgartner.systems"
)

func main() {
	if err := build(); err != nil {
		fmt.Fprintln(os.Stderr, "build failed:", err)
		os.Exit(1)
	}

	fmt.Println("Built public/")
}

func build() error {
	if err := recreateDir(publicDir); err != nil {
		return fmt.Errorf(
			"prepare public directory: %w",
			err,
		)
	}

	if err := buildStaticPages(); err != nil {
		return err
	}

	if err := buildTechnologyPages(); err != nil {
		return err
	}

	if err := buildArticlePages(); err != nil {
		return err
	}

	if err := buildArcaDemo(); err != nil {
		return err
	}

	if err := buildSEOFiles(); err != nil {
		return err
	}

	if err := copyDir(
		staticDir,
		filepath.Join(
			publicDir,
			"assets",
		),
	); err != nil {
		return fmt.Errorf(
			"copy static assets: %w",
			err,
		)
	}

	return nil
}
