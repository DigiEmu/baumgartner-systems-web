package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func buildSEOFiles() error {
	robots := "User-agent: *\n" +
		"Allow: /\n\n" +
		"Sitemap: " + siteURL + "/sitemap.xml\n"

	if err := writeTextFile(
		filepath.Join(
			publicDir,
			"robots.txt",
		),
		robots,
	); err != nil {
		return fmt.Errorf(
			"build robots.txt: %w",
			err,
		)
	}

	sitemap := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
		"<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">\n"

	for _, p := range staticPages {
		if p.Target == "404.html" {
			continue
		}

		path := strings.TrimSuffix(p.Target, "index.html")
		path = strings.TrimSuffix(path, "/")
		var loc string
		if path == "" {
			loc = siteURL + "/"
		} else {
			loc = siteURL + "/" + path + "/"
		}

		sitemap +=
			"  <url><loc>" +
				loc +
				"</loc></url>\n"
	}

	for _, technology := range technologies {
		sitemap +=
			"  <url><loc>" +
				siteURL +
				"/" +
				technology.Slug +
				"/</loc></url>\n"
	}

	sitemap +=
		"  <url><loc>" +
			siteURL +
			"/articles/</loc></url>\n"

	for _, article := range articles {
		sitemap +=
			"  <url><loc>" +
				siteURL +
				"/articles/" +
				article.Slug +
				"/</loc></url>\n"
	}

	sitemap +=
		"  <url><loc>" +
			siteURL +
			"/arca/demo/</loc></url>\n"

	sitemap += "</urlset>\n"

	if err := writeTextFile(
		filepath.Join(
			publicDir,
			"sitemap.xml",
		),
		sitemap,
	); err != nil {
		return fmt.Errorf(
			"build sitemap.xml: %w",
			err,
		)
	}

	return nil
}
