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
	siteURL   = "https://baumgartner.systems"
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

	MetaTitle       string
	MetaDescription string
	CanonicalPath   string
	SiteURL         string

	Complete bool

	ArchitectureTitle string
	ArchitectureIntro string
	Architecture      []architectureStep

	WorkflowTitle string
	WorkflowIntro string
	Workflow      []architectureStep

	EvidenceTitle string
	EvidenceIntro string
	Evidence      []evidenceItem

	BoundaryTitle string
	BoundaryIntro string
	Boundaries    []boundaryItem

	PilotTitle string
	PilotIntro string
	Pilot      []pilotStep
	PilotCTA   string
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

	/* =====================================================
	   ARCA
	   ===================================================== */

	{
		Slug: "arca",
		Name: "Arca",

		MetaTitle:       "Arca — AI Evidence Infrastructure | Baumgartner Systems",
		MetaDescription: "Arca captures provenance, lineage and verification evidence across AI-assisted workflows so organisations can reconstruct what happened and why.",
		Tagline:         "Evidence workflow integrity.",

		Description: "Infrastructure for capturing, preserving and reconstructing " +
			"evidence across AI-assisted workflows.",

		Overview: "Arca creates a verifiable evidence trail around an AI-assisted " +
			"workflow. It preserves the information needed to understand what was " +
			"available, what changed and how a recorded process can later be reconstructed.",

		Layer:  "Evidence workflow",
		Status: "Enterprise readiness",
		Image:  "/assets/images/brand/arca.webp",

		Complete: true,

		ArchitectureTitle: "Evidence is a system, not a screenshot.",

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

		WorkflowTitle: "From source material to reconstructable evidence.",

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

		EvidenceTitle: "What remains after the workflow.",

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

		BoundaryTitle: "Verification is not omniscience.",

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

		PilotTitle: "Start with one workflow.",

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

		PilotCTA: "Discuss an Arca Pilot",
	},

	/* =====================================================
	   CORE
	   ===================================================== */

	{
		Slug: "core",
		Name: "Core",

		MetaTitle:       "Core — Deterministic Verification | Baumgartner Systems",
		MetaDescription: "DigiEmu Core provides deterministic identity, state, admission and verification primitives for reproducible AI and knowledge infrastructure.",
		Tagline:         "Deterministic trust foundation.",

		Description: "A deterministic foundation for identity, provenance, lineage, " +
			"admission and reproducible system state.",

		Overview: "DigiEmu Core provides the deterministic substrate beneath the wider " +
			"verification architecture. It defines how relevant artifacts, states and " +
			"admission decisions are represented so the same referentially closed inputs " +
			"can be evaluated consistently.",

		Layer:  "Verification foundation",
		Status: "Active",
		Image:  "/assets/images/brand/core.webp",

		Complete: true,

		ArchitectureTitle: "Trust begins with explicit architecture.",

		ArchitectureIntro: "Core is designed around explicit architecture boundaries. " +
			"Identity, capability, ownership, admission and verification are represented " +
			"as separate but connected parts of the system rather than implicit runtime behavior.",

		Architecture: []architectureStep{
			{
				Number: "01",
				Name:   "Identity",
				Description: "Relevant entities and artifacts are represented with stable " +
					"identities so state and evidence can refer to the same logical objects.",
			},
			{
				Number: "02",
				Name:   "Capability",
				Description: "The system records which capabilities are available and which " +
					"components are allowed to provide them.",
			},
			{
				Number: "03",
				Name:   "Ownership",
				Description: "Aggregate ownership and responsibility boundaries are made " +
					"explicit so architectural authority is not inferred implicitly.",
			},
			{
				Number: "04",
				Name:   "Admission",
				Description: "Inputs and candidate state transitions are evaluated against " +
					"declared architectural conditions before they are accepted.",
			},
			{
				Number: "05",
				Name:   "Verification",
				Description: "The resulting canonical state and admission outcome can be " +
					"checked against the same referentially closed input basis.",
			},
		},

		WorkflowTitle: "Evaluate before consequence.",

		WorkflowIntro: "In practice, Core acts as an admission and verification layer. " +
			"It evaluates whether a proposed state or transition belongs inside the " +
			"declared architecture before downstream consequences are accepted.",

		Workflow: []architectureStep{
			{
				Number:      "A",
				Name:        "Receive",
				Description: "A referentially closed input set enters the verification boundary.",
			},
			{
				Number: "B",
				Name:   "Resolve",
				Description: "Relevant identities, capabilities, ownership and architectural " +
					"constraints are resolved from the declared system state.",
			},
			{
				Number: "C",
				Name:   "Evaluate",
				Description: "The candidate state or transition is checked against admission " +
					"rules before acceptance.",
			},
			{
				Number: "D",
				Name:   "Disposition",
				Description: "Core produces a structured outcome describing whether the " +
					"candidate is admitted, rejected or otherwise classified.",
			},
			{
				Number: "E",
				Name:   "Record",
				Description: "The verification outcome and relevant state are preserved so " +
					"the decision can later be inspected or replayed.",
			},
		},

		EvidenceTitle: "What remains after admission.",

		EvidenceIntro: "Core does not merely return a yes-or-no result. The value lies " +
			"in preserving the deterministic basis and structured outcome around the " +
			"admission decision.",

		Evidence: []evidenceItem{
			{
				Name:        "Canonical identity",
				Description: "Stable references to the entities and artifacts inside the verification boundary.",
			},
			{
				Name:        "Capability state",
				Description: "Recorded information about which declared capabilities are present.",
			},
			{
				Name:        "Ownership state",
				Description: "Explicit ownership and authority relationships relevant to the architecture.",
			},
			{
				Name:        "Admission basis",
				Description: "The referentially closed input set used to evaluate the candidate state.",
			},
			{
				Name:        "Disposition",
				Description: "A structured admission outcome rather than an implicit runtime side effect.",
			},
			{
				Name:        "Replay surface",
				Description: "The preserved basis required to evaluate the same canonical verification question again.",
			},
		},

		BoundaryTitle: "Determinism has a boundary.",

		BoundaryIntro: "Core provides deterministic verification of declared architecture " +
			"and recorded state. That does not mean it can establish truth outside the " +
			"system boundary or replace governance authority.",

		Boundaries: []boundaryItem{
			{
				Text: "Core does not prove that an external real-world fact is true.",
			},
			{
				Text: "Core does not decide whether a governance policy is ethically or legally correct.",
			},
			{
				Text: "Core does not infer authority that has not been explicitly represented.",
			},
			{
				Text: "Core does not make undocumented runtime behavior part of the verification boundary.",
			},
			{
				Text: "Determinism applies to the canonical verification result for the same referentially closed inputs, not to byte-identical binaries across different toolchains.",
			},
		},

		PilotTitle: "Start with one architecture boundary.",

		PilotIntro: "A useful Core pilot starts with one architecture boundary where " +
			"admission, state integrity or reproducibility matters. The goal is to make " +
			"one verification decision explicit, deterministic and reviewable.",

		Pilot: []pilotStep{
			{
				Number: "01",
				Name:   "Choose one boundary",
				Description: "Select one architecture decision where acceptance or rejection " +
					"should be explicit and independently reviewable.",
			},
			{
				Number: "02",
				Name:   "Declare the state basis",
				Description: "Define identities, capabilities, ownership and other inputs " +
					"required to make the verification question referentially closed.",
			},
			{
				Number: "03",
				Name:   "Verify and replay",
				Description: "Run the admission decision, preserve the structured result and " +
					"confirm that the same canonical verification outcome can be reproduced.",
			},
		},

		PilotCTA: "Discuss a Core Pilot",
	},

	/* =====================================================
	   SKC
	   ===================================================== */

	{
		Slug: "skc",
		Name: "SKC",

		MetaTitle:       "SKC — Semantic Knowledge Compression | Baumgartner Systems",
		MetaDescription: "SKC explores deterministic semantic compression with reconstructable state, measurable fidelity and graph-level verification.",
		Tagline:         "Semantic knowledge compression.",

		Description: "Reconstructable semantic state designed for deterministic " +
			"compression, reconstruction and verification.",

		Overview: "SKC explores a canonical semantic state space in which knowledge can " +
			"be represented more compactly while retaining the structure required for " +
			"deterministic reconstruction and graph-level verification.",

		Layer:  "Semantic state",
		Status: "Research",
		Image:  "/assets/images/brand/skc.webp",

		Complete: true,

		ArchitectureTitle: "Compression is useful only if meaning survives.",

		ArchitectureIntro: "SKC separates semantic structure from redundant representation. " +
			"The goal is not arbitrary size reduction, but a compact semantic core that " +
			"can be reconstructed and evaluated against the original knowledge structure.",

		Architecture: []architectureStep{
			{
				Number:      "01",
				Name:        "Ingest",
				Description: "A source representation enters the semantic compression boundary.",
			},
			{
				Number:      "02",
				Name:        "Normalize",
				Description: "Relevant semantic units and relationships are normalized into a stable representation.",
			},
			{
				Number:      "03",
				Name:        "Compress",
				Description: "Redundant representation is reduced while retaining the semantic core.",
			},
			{
				Number:      "04",
				Name:        "Reconstruct",
				Description: "The compact state is expanded into a reconstruction of the original semantic structure.",
			},
			{
				Number:      "05",
				Name:        "Verify",
				Description: "Reconstruction fidelity and graph-level equivalence can be evaluated.",
			},
		},

		WorkflowTitle: "Reduce representation, preserve structure.",

		WorkflowIntro: "SKC treats compression as a reversible semantic transformation. " +
			"Each stage remains inspectable so reduction can be distinguished from semantic loss.",

		Workflow: []architectureStep{
			{
				Number:      "A",
				Name:        "Measure",
				Description: "Establish the size and structure of the original knowledge representation.",
			},
			{
				Number:      "B",
				Name:        "Extract",
				Description: "Identify the semantic core required to preserve meaningful relationships.",
			},
			{
				Number:      "C",
				Name:        "Encode",
				Description: "Represent the semantic core in a more compact deterministic form.",
			},
			{
				Number:      "D",
				Name:        "Rebuild",
				Description: "Reconstruct the expected knowledge structure from the compact representation.",
			},
			{
				Number:      "E",
				Name:        "Compare",
				Description: "Evaluate reconstruction fidelity, graph equivalence and reduction metrics.",
			},
		},

		EvidenceTitle: "Compression must leave measurable evidence.",

		EvidenceIntro: "SKC makes compression quality inspectable through explicit measurements " +
			"rather than treating a smaller output as sufficient evidence of success.",

		Evidence: []evidenceItem{
			{
				Name:        "Blueprint size",
				Description: "Size of the original representation used as the comparison basis.",
			},
			{
				Name:        "Semantic core size",
				Description: "Size of the compact semantic representation.",
			},
			{
				Name:        "Reduction ratio",
				Description: "Measured reduction between the original and compact representations.",
			},
			{
				Name:        "Reconstruction fidelity",
				Description: "Degree to which the original semantic structure can be reconstructed.",
			},
			{
				Name:        "Graph equivalence",
				Description: "Comparison of important semantic relationships before and after reconstruction.",
			},
			{
				Name:        "Deterministic result",
				Description: "A repeatable reconstruction and validation outcome for the same defined inputs.",
			},
		},

		BoundaryTitle: "Semantic compression is not semantic omniscience.",

		BoundaryIntro: "SKC can evaluate preservation within its represented semantic boundary. " +
			"It does not claim that the representation captures every possible meaning or interpretation.",

		Boundaries: []boundaryItem{
			{
				Text: "SKC does not prove that the source knowledge itself is factually correct.",
			},
			{
				Text: "SKC does not prove that every possible human interpretation has been preserved.",
			},
			{
				Text: "Compression ratios are meaningful only relative to the defined source and reconstruction basis.",
			},
			{
				Text: "Reconstruction fidelity depends on what information is represented inside the semantic boundary.",
			},
			{
				Text: "Determinism applies to the defined transformation and reconstruction process, not to unspecified external context.",
			},
		},

		PilotTitle: "Start with one semantic corpus.",

		PilotIntro: "An SKC pilot starts with a bounded body of structured or semi-structured " +
			"knowledge where reduction, reconstruction and semantic integrity can all be measured.",

		Pilot: []pilotStep{
			{
				Number:      "01",
				Name:        "Choose a corpus",
				Description: "Select one bounded knowledge set with a clearly measurable source representation.",
			},
			{
				Number:      "02",
				Name:        "Define fidelity",
				Description: "Agree which semantic relationships must survive compression and reconstruction.",
			},
			{
				Number:      "03",
				Name:        "Compress and verify",
				Description: "Measure reduction, reconstruct the semantic state and compare the result against the agreed fidelity boundary.",
			},
		},

		PilotCTA: "Discuss an SKC Pilot",
	},

	/* =====================================================
	   VSC
	   ===================================================== */

	{
		Slug: "vsc",
		Name: "VSC",

		MetaTitle:       "VSC — Verification State Compression | Baumgartner Systems",
		MetaDescription: "VSC reduces verification-relevant system state while preserving the evidence needed for reconstruction, comparison and reproducibility.",
		Tagline:         "Verification state compression.",

		Description: "A verification-oriented approach to reducing system state while " +
			"preserving the information required for reconstruction and reproducibility.",

		Overview: "VSC investigates how operational or model-related state can be reduced " +
			"without discarding the information needed to verify what state was used, " +
			"how it changed and whether an equivalent verification question can be replayed.",

		Layer:  "State optimisation",
		Status: "Research",
		Image:  "/assets/images/brand/vsc.webp",

		Complete: true,

		ArchitectureTitle: "Optimisation should not erase the verification surface.",

		ArchitectureIntro: "VSC treats state reduction as a verification problem. " +
			"Compression is valuable only when the retained state still supports the " +
			"reconstruction, comparison and reproducibility requirements of the workflow.",

		Architecture: []architectureStep{
			{
				Number:      "01",
				Name:        "Observe",
				Description: "Establish the original system state and the verification information it contains.",
			},
			{
				Number:      "02",
				Name:        "Classify",
				Description: "Separate verification-relevant state from redundant or replaceable representation.",
			},
			{
				Number:      "03",
				Name:        "Compress",
				Description: "Reduce state representation within the declared optimisation boundary.",
			},
			{
				Number:      "04",
				Name:        "Restore",
				Description: "Reconstruct the state required for the intended verification operation.",
			},
			{
				Number:      "05",
				Name:        "Compare",
				Description: "Evaluate whether the retained and restored state supports the expected verification outcome.",
			},
		},

		WorkflowTitle: "Preserve what verification actually needs.",

		WorkflowIntro: "VSC is intended to sit between expensive state representation and " +
			"verification-sensitive execution. It asks which information must remain " +
			"available for reproducibility instead of assuming all state must remain unchanged.",

		Workflow: []architectureStep{
			{
				Number:      "A",
				Name:        "Baseline",
				Description: "Record the original state, parameters and verification target.",
			},
			{
				Number:      "B",
				Name:        "Reduce",
				Description: "Apply state reduction while keeping the defined verification requirements explicit.",
			},
			{
				Number:      "C",
				Name:        "Execute",
				Description: "Use the reduced state inside the intended workload or verification path.",
			},
			{
				Number:      "D",
				Name:        "Recover",
				Description: "Restore or reconstruct the state needed for comparison.",
			},
			{
				Number:      "E",
				Name:        "Verify",
				Description: "Compare the resulting verification surface against the baseline requirements.",
			},
		},

		EvidenceTitle: "Optimisation needs a reproducible comparison basis.",

		EvidenceIntro: "A VSC result should make it possible to distinguish resource reduction " +
			"from loss of verification-relevant state.",

		Evidence: []evidenceItem{
			{
				Name:        "Baseline state",
				Description: "The recorded state before optimisation.",
			},
			{
				Name:        "Compression boundary",
				Description: "A definition of which state may be reduced and which state must remain available.",
			},
			{
				Name:        "Resource delta",
				Description: "Measured change in the relevant memory or state footprint.",
			},
			{
				Name:        "Restored state",
				Description: "The state reconstructed for verification or comparison.",
			},
			{
				Name:        "Verification outcome",
				Description: "The observed result after reduction and reconstruction.",
			},
			{
				Name:        "Replay basis",
				Description: "The parameters and retained evidence required to repeat the comparison.",
			},
		},

		BoundaryTitle: "Lower resource use is not automatically equivalent state.",

		BoundaryIntro: "VSC does not treat resource reduction itself as evidence of correctness. " +
			"Any equivalence claim is limited to the defined verification surface and experiment boundary.",

		Boundaries: []boundaryItem{
			{
				Text: "VSC does not prove that every internal model state is unchanged.",
			},
			{
				Text: "VSC does not claim semantic equivalence outside the verification criteria defined for the experiment.",
			},
			{
				Text: "Resource reduction does not by itself demonstrate correctness or reproducibility.",
			},
			{
				Text: "Results depend on the declared model, runtime, parameters and state boundary.",
			},
			{
				Text: "Verification claims must remain scoped to the evidence captured by the experiment.",
			},
		},

		PilotTitle: "Start with one constrained workload.",

		PilotIntro: "A VSC pilot should begin with one reproducible workload where the baseline " +
			"resource footprint and verification requirements can both be measured precisely.",

		Pilot: []pilotStep{
			{
				Number:      "01",
				Name:        "Establish baseline",
				Description: "Record the workload, parameters, state footprint and expected verification surface.",
			},
			{
				Number:      "02",
				Name:        "Apply reduction",
				Description: "Introduce VSC within a clearly defined and isolated optimisation boundary.",
			},
			{
				Number:      "03",
				Name:        "Compare results",
				Description: "Measure resource reduction and verify whether the defined reconstruction and replay criteria remain satisfied.",
			},
		},

		PilotCTA: "Discuss a VSC Pilot",
	},

	/* =====================================================
	   FOAM
	   ===================================================== */

	{
		Slug: "foam",
		Name: "Foam",

		MetaTitle:       "Foam — Memory State Optimisation | Baumgartner Systems",
		MetaDescription: "Foam explores memory-state optimisation for efficient, reconstructable and verification-aware AI infrastructure.",
		Tagline:         "Memory state optimisation.",

		Description: "Experimental memory-state optimisation for efficient, reconstructable " +
			"and verification-aware AI infrastructure.",

		Overview: "Foam explores how memory state can be represented more efficiently " +
			"while preserving the information needed by higher-level reconstruction and " +
			"verification processes.",

		Layer:  "Memory",
		Status: "Research",
		Image:  "/assets/images/brand/foam.webp",

		Complete: true,

		ArchitectureTitle: "Memory is infrastructure, not an unlimited resource.",

		ArchitectureIntro: "Foam investigates memory optimisation as a structured state problem. " +
			"Instead of treating memory reduction as an opaque runtime trick, relevant " +
			"state transitions and recovery requirements remain explicit.",

		Architecture: []architectureStep{
			{
				Number:      "01",
				Name:        "Profile",
				Description: "Observe the memory state and identify the structures responsible for the active footprint.",
			},
			{
				Number:      "02",
				Name:        "Separate",
				Description: "Distinguish active, reconstructable and redundant memory representation.",
			},
			{
				Number:      "03",
				Name:        "Optimise",
				Description: "Reduce or reorganise the memory representation within the defined boundary.",
			},
			{
				Number:      "04",
				Name:        "Recover",
				Description: "Restore required state when downstream computation or verification needs it.",
			},
			{
				Number:      "05",
				Name:        "Measure",
				Description: "Compare footprint, recoverability and verification-relevant behavior against the baseline.",
			},
		},

		WorkflowTitle: "Use less memory without making state invisible.",

		WorkflowIntro: "Foam is intended for controlled environments where memory pressure " +
			"matters but state must remain understandable enough to support reconstruction " +
			"and higher-level verification.",

		Workflow: []architectureStep{
			{
				Number:      "A",
				Name:        "Profile",
				Description: "Capture a reproducible baseline of the workload and its memory behavior.",
			},
			{
				Number:      "B",
				Name:        "Map",
				Description: "Identify which memory structures are essential, reconstructable or redundant.",
			},
			{
				Number:      "C",
				Name:        "Optimise",
				Description: "Apply the selected memory-state optimisation strategy.",
			},
			{
				Number:      "D",
				Name:        "Restore",
				Description: "Recover state when the workload requires information that is no longer fully resident.",
			},
			{
				Number:      "E",
				Name:        "Evaluate",
				Description: "Measure memory reduction together with recovery and verification behavior.",
			},
		},

		EvidenceTitle: "Memory reduction should remain measurable.",

		EvidenceIntro: "Foam experiments should preserve enough evidence to show what memory " +
			"was reduced, what was recoverable and what trade-offs were introduced.",

		Evidence: []evidenceItem{
			{
				Name:        "Baseline footprint",
				Description: "Memory use measured before optimisation.",
			},
			{
				Name:        "State map",
				Description: "Classification of relevant memory structures inside the experiment boundary.",
			},
			{
				Name:        "Optimised footprint",
				Description: "Memory use measured after the optimisation step.",
			},
			{
				Name:        "Recovery behavior",
				Description: "Evidence describing whether required state can be reconstructed when needed.",
			},
			{
				Name:        "Resource delta",
				Description: "Measured reduction or redistribution of the active memory footprint.",
			},
			{
				Name:        "Verification basis",
				Description: "The baseline and retained evidence needed to compare experiment outcomes.",
			},
		},

		BoundaryTitle: "Memory efficiency is not proof of behavioral equivalence.",

		BoundaryIntro: "Foam is experimental memory infrastructure. Any claim about preservation " +
			"or equivalence must remain limited to the state and measurements actually evaluated.",

		Boundaries: []boundaryItem{
			{
				Text: "Foam does not prove that every internal runtime state remains identical.",
			},
			{
				Text: "Reduced memory use does not automatically imply equivalent model behavior.",
			},
			{
				Text: "Recoverability applies only to state represented inside the defined optimisation boundary.",
			},
			{
				Text: "Performance results depend on workload, runtime, model and system configuration.",
			},
			{
				Text: "Experimental results must not be generalized beyond the measured evidence.",
			},
		},

		PilotTitle: "Start with one measurable memory problem.",

		PilotIntro: "A Foam experiment should begin with a reproducible workload whose memory " +
			"footprint can be measured before and after optimisation without changing the " +
			"verification question midway through the experiment.",

		Pilot: []pilotStep{
			{
				Number:      "01",
				Name:        "Measure baseline",
				Description: "Record the workload, memory footprint and state requirements before optimisation.",
			},
			{
				Number:      "02",
				Name:        "Define recovery boundary",
				Description: "Identify which state must remain resident, reconstructable or independently verifiable.",
			},
			{
				Number:      "03",
				Name:        "Optimise and compare",
				Description: "Apply the experiment, measure resource change and evaluate recovery against the baseline.",
			},
		},

		PilotCTA: "Discuss a Foam Experiment",
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
		"<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">\n" +
		"  <url><loc>" + siteURL + "/</loc></url>\n"

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

