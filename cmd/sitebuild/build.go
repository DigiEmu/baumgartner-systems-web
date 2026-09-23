package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func buildStaticPages() error {
	for _, p := range staticPages {
		target :=
			filepath.Join(
				publicDir,
				p.Target,
			)

		if err := copyFile(
			p.Source,
			target,
		); err != nil {
			return fmt.Errorf(
				"build page %s -> %s: %w",
				p.Source,
				target,
				err,
			)
		}
	}

	return nil
}

func buildTechnologyPages() error {
	tmpl, err :=
		template.ParseFiles(
			"templates/technology.html",
		)

	if err != nil {
		return fmt.Errorf(
			"parse technology template: %w",
			err,
		)
	}

	for _, technology := range technologies {
		technology.SiteURL = siteURL
		technology.CanonicalPath = "/" + technology.Slug + "/"

		if technology.MetaTitle == "" {
			technology.MetaTitle = technology.Name + " — Baumgartner Systems"
		}

		if technology.MetaDescription == "" {
			technology.MetaDescription = technology.Description
		}

		target :=
			filepath.Join(
				publicDir,
				technology.Slug,
				"index.html",
			)

		if err :=
			renderTemplate(
				tmpl,
				target,
				technology,
			); err != nil {

			return fmt.Errorf(
				"build technology page %s: %w",
				technology.Slug,
				err,
			)
		}
	}

	return nil
}

func buildArticlePages() error {
	indexTemplate, err := template.ParseFiles(
		"templates/articles.html",
	)

	if err != nil {
		return fmt.Errorf(
			"parse articles index template: %w",
			err,
		)
	}

	indexTarget := filepath.Join(
		publicDir,
		"articles",
		"index.html",
	)

	if err := renderTemplate(
		indexTemplate,
		indexTarget,
		articles,
	); err != nil {
		return fmt.Errorf(
			"build articles index: %w",
			err,
		)
	}

	articleTemplate, err := template.ParseFiles(
		"templates/article.html",
	)

	if err != nil {
		return fmt.Errorf(
			"parse article template: %w",
			err,
		)
	}

	for _, article := range articles {
		article.SiteURL = siteURL
		article.CanonicalPath =
			"/articles/" + article.Slug + "/"

		target := filepath.Join(
			publicDir,
			"articles",
			article.Slug,
			"index.html",
		)

		if err := renderTemplate(
			articleTemplate,
			target,
			article,
		); err != nil {
			return fmt.Errorf(
				"build article %s: %w",
				article.Slug,
				err,
			)
		}
	}

	return nil
}
func buildArcaDemo() error {
	tmpl, err :=
		template.ParseFiles(
			"templates/arca-demo.html",
		)

	if err != nil {
		return fmt.Errorf(
			"parse Arca demo template: %w",
			err,
		)
	}

	target :=
		filepath.Join(
			publicDir,
			"arca",
			"demo",
			"index.html",
		)

	if err :=
		renderTemplate(
			tmpl,
			target,
			nil,
		); err != nil {

		return fmt.Errorf(
			"build Arca demo: %w",
			err,
		)
	}

	return nil
}

func renderTemplate(
	tmpl *template.Template,
	target string,
	data any,
) error {
	if err :=
		os.MkdirAll(
			filepath.Dir(target),
			0o755,
		); err != nil {

		return err
	}

	out, err :=
		os.Create(target)

	if err != nil {
		return err
	}

	defer func() {
		_ = out.Close()
	}()

	if err :=
		tmpl.Execute(
			out,
			data,
		); err != nil {

		return err
	}

	return out.Close()
}
