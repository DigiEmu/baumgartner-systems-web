package main

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
)

const (
	publicDir = "public"
	staticDir = "static"
)

type page struct {
	Source string
	Target string
}

type architectureStep struct {
	Number      string
	Name        string
	Description string
}

type evidenceItem struct {
	Name        string
	Description string
}

type boundaryItem struct {
	Text string
}

type pilotStep struct {
	Number      string
	Name        string
	Description string
}

type technologyPage struct {
	Slug        string
	Name        string
	Tagline     string
	Description string
	Overview    string
	Layer       string
	Status      string
	Image       string

	Complete bool

	ArchitectureIntro string
	Architecture      []architectureStep

	WorkflowIntro string
	Workflow      []architectureStep

	EvidenceIntro string
	Evidence      []evidenceItem

	BoundaryIntro string
	Boundaries    []boundaryItem

	PilotIntro string
	Pilot      []pilotStep
}

var staticPages = []page{
	{
		Source: "templates/index.html",
		Target: "index.html",
	},
	{
		Source: "templates/404.html",
		Target: "404.html",
	},
}

var technologies = []technologyPage{

	{
		Slug:    "arca",
		Name:    "Arca",
		Tagline: "Evidence workflow integrity.",

		Description: "Infrastructure for capturing, preserving and reconstructing " +
			"evidence across AI-assisted workflows.",

		Overview: "Arca creates a verifiable evidence trail around an AI-assisted " +
			"workflow. It preserves the information needed to understand what was " +
			"available, what changed and how a recorded process can later be reconstructed.",

		Layer:  "Evidence workflow",
		Status: "Enterprise readiness",
		Image:  "/assets/images/brand/arca.webp",

		Complete: true,

		ArchitectureIntro: "Arca separates source material, transformations, provenance " +
			"and verification into explicit stages rather than treating an AI-assisted " +
			"result as an isolated output.",

		Architecture: []architectureStep{
			{
				Number: "01",
				Name:   "Source intake",
				Description: "Documents, repository material and other permitted inputs " +
					"enter the workflow as identifiable source material.",
			},
			{
				Number: "02",
				Name:   "Identity",
				Description: "Relevant artifacts receive stable identities so later " +
					"evidence can refer to the same recorded objects.",
			},
			{
				Number: "03",
				Name:   "Provenance",
				Description: "Relationships between sources, transformations and generated " +
					"artifacts are preserved as part of the evidence trail.",
			},
			{
				Number: "04",
				Name:   "Build",
				Description: "Arca produces controlled outputs from the recorded source " +
					"basis while retaining the information needed for reconstruction.",
			},
			{
				Number: "05",
				Name:   "Verification",
				Description: "Recorded evidence can be checked against expected state, " +
					"lineage and integrity information.",
			},
		},

		WorkflowIntro: "In practice, Arca sits around an existing workflow. It does not " +
			"need to replace the organisation's business process in order to make that " +
			"process more traceable.",

		Workflow: []architectureStep{
			{
				Number: "A",
				Name:   "Collect",
				Description: "Bring together the source material relevant to one defined " +
					"decision or documentation workflow.",
			},
			{
				Number: "B",
				Name:   "Record",
				Description: "Capture the evidence basis, source identities and relevant " +
					"relationships before publication or handoff.",
			},
			{
				Number: "C",
				Name:   "Produce",
				Description: "Generate the required document, package or evidence output " +
					"from the recorded source basis.",
			},
			{
				Number: "D",
				Name:   "Verify",
				Description: "Check whether the resulting evidence package still corresponds " +
					"to the recorded source and provenance state.",
			},
			{
				Number: "E",
				Name:   "Reconstruct",
				Description: "Return later to the evidence trail and understand what the " +
					"workflow contained at the relevant point in time.",
			},
		},

		EvidenceIntro: "The purpose of Arca is not simply to generate another document. " +
			"It is to preserve the evidence required to inspect and reconstruct how that " +
			"document or workflow came into existence.",

		Evidence: []evidenceItem{
			{
				Name:        "Source basis",
				Description: "The material that was actually available to the workflow.",
			},
			{
				Name:        "Artifact identity",
				Description: "Stable identification of relevant recorded artifacts.",
			},
			{
				Name:        "Lineage",
				Description: "Relationships between source material and resulting outputs.",
			},
			{
				Name:        "Provenance",
				Description: "Recorded information about where evidence came from and how it moved.",
			},
			{
				Name:        "Build state",
				Description: "The recorded state associated with a generated evidence package.",
			},
			{
				Name:        "Verification result",
				Description: "A structured result describing what was checked and what was observed.",
			},
		},

		BoundaryIntro: "Arca makes recorded evidence inspectable. That is deliberately " +
			"different from claiming that a software system can prove every aspect of truth, " +
			"reasoning or governance.",

		Boundaries: []boundaryItem{
			{
				Text: "Arca does not prove the objective truth of an AI-generated statement.",
			},
			{
				Text: "Arca does not prove the correctness of a model's internal reasoning.",
			},
			{
				Text: "Arca does not determine whether a governance decision is ethically correct.",
			},
			{
				Text: "Arca cannot reconstruct evidence that was never recorded.",
			},
			{
				Text: "Verification applies to the recorded evidence boundary and the claims that boundary supports.",
			},
		},

		PilotIntro: "The smallest useful Arca engagement is one real workflow with a " +
			"clear evidence problem. We define the boundary first, then verify whether Arca " +
			"can make that process materially more reconstructable.",

		Pilot: []pilotStep{
			{
				Number: "01",
				Name:   "Choose one workflow",
				Description: "Select a process where AI-assisted work, documentation or " +
					"decision support creates a genuine traceability requirement.",
			},
			{
				Number: "02",
				Name:   "Define the evidence boundary",
				Description: "Agree what must be captured, what must remain outside scope " +
					"and what a successful verification result should demonstrate.",
			},
			{
				Number: "03",
				Name:   "Run the pilot",
				Description: "Apply Arca to the workflow and produce a concrete evidence " +
					"package that can be reviewed with the organisation.",
			},
		},
	},
	{
		Slug:    "core",
		Name:    "Core",
		Tagline: "Trusted digital foundation.",
		Description: "A deterministic foundation for identity, provenance, lineage " +
			"and reproducible system state.",
		Overview: "DigiEmu Core provides the deterministic substrate beneath the wider " +
			"verification architecture.",
		Layer:  "Foundation",
		Status: "Active",
		Image:  "/assets/images/brand/core.webp",
	},
	{
		Slug:    "skc",
		Name:    "SKC",
		Tagline: "Semantic knowledge compression.",
		Description: "Reconstructable semantic state designed for deterministic " +
			"compression and verification.",
		Overview: "SKC explores compact representations of semantic state that retain " +
			"the information required for deterministic reconstruction.",
		Layer:  "Semantic state",
		Status: "Research",
		Image:  "/assets/images/brand/skc.webp",
	},
	{
		Slug:    "vsc",
		Name:    "VSC",
		Tagline: "Verification state compression.",
		Description: "A verification-oriented approach to state compression for " +
			"resource-constrained AI systems.",
		Overview: "VSC investigates how system state can be reduced while preserving " +
			"the information required for verification and reproducibility.",
		Layer:  "State optimisation",
		Status: "Research",
		Image:  "/assets/images/brand/vsc.webp",
	},
	{
		Slug:    "foam",
		Name:    "Foam",
		Tagline: "Memory state optimisation.",
		Description: "Experimental memory-state optimisation for efficient and " +
			"reconstructable AI infrastructure.",
		Overview: "Foam explores memory-state optimisation techniques intended to " +
			"reduce resource usage while preserving important system state.",
		Layer:  "Memory",
		Status: "Research",
		Image:  "/assets/images/brand/foam.webp",
	},
}

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
