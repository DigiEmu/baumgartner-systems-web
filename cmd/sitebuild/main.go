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

type articleSection struct {
	Title      string
	Paragraphs []string
	Points     []string
}

type articleLink struct {
	Title string
	Slug  string
}

type articlePage struct {
	Slug            string
	Title           string
	MetaTitle       string
	MetaDescription string
	Description     string
	Published       string
	Modified        string
	ReadingTime     string
	CanonicalPath   string
	SiteURL         string
	Related         []articleLink
	Sections        []articleSection
}

var articles = []articlePage{
	{
		Slug: "ai-evidence-infrastructure",

		Title: "What Is AI Evidence Infrastructure?",

		MetaTitle: "What Is AI Evidence Infrastructure? | Baumgartner Systems",

		MetaDescription: "A technical introduction to AI evidence infrastructure: provenance, lineage, system state, transformations, verification boundaries and reconstructable AI-assisted decision trails.",

		Description: "AI evidence infrastructure preserves the information required to reconstruct, inspect and verify how an AI-assisted process moved from source material to an outcome.",

		Published:   "2026-09-13",
		Modified:    "2026-09-13",
		ReadingTime: "8 min read",

		Sections: []articleSection{
			{
				Title: "AI outputs are not evidence trails",
				Paragraphs: []string{
					"An AI system can produce a useful answer while leaving an organisation unable to reconstruct how that answer became part of a decision.",
					"A final output normally does not contain the complete evidence basis: source identity, document versions, model or agent identity, configuration, transformations, human review and the state of the workflow at the time the result was produced.",
					"AI evidence infrastructure treats those elements as part of the operational system rather than as information to reconstruct manually after an incident or audit.",
				},
			},
			{
				Title: "The evidence chain",
				Paragraphs: []string{
					"A useful verification model follows the chain through which information changes role and meaning.",
				},
				Points: []string{
					"Source evidence establishes what entered the process.",
					"Identity records which human, model, agent or system participated.",
					"Lineage records how information moved between stages.",
					"State captures the relevant configuration and environment.",
					"Provenance records origin, custody and transformations.",
					"Verification establishes whether the recorded trail can be independently inspected or replayed.",
				},
			},
			{
				Title: "Why ordinary audit logs are not enough",
				Paragraphs: []string{
					"Traditional logs are useful, but logging and evidence preservation are not the same problem. A log can record that an event happened without preserving the exact information needed to understand what the event meant.",
					"For verification, the important question is not simply whether an action was recorded. The question is whether the recorded material is sufficient to reconstruct the verification question later.",
				},
			},
			{
				Title: "Meaning can change while data remains intact",
				Paragraphs: []string{
					"Information often passes through several representations: evidence, record, summary, AI analysis, decision and later review.",
					"Each representation may be technically valid while context, qualification or emphasis changes. That means integrity cannot be reduced to checking whether a file or database record remained byte-identical.",
					"A verification architecture therefore needs to preserve both the recorded artefacts and the relationships between transformations.",
				},
			},
			{
				Title: "Verification requires an explicit boundary",
				Paragraphs: []string{
					"Evidence infrastructure does not prove everything. A useful verification system declares what information belongs inside the verification boundary and what remains external.",
					"For example, a system may verify the recorded provenance and transformation history of an AI-assisted decision without proving that every external real-world statement in the source material is objectively true.",
				},
				Points: []string{
					"Do not infer authority that was never represented.",
					"Do not treat undocumented runtime behaviour as verified state.",
					"Do not confuse provenance with factual truth.",
					"Do not claim reproducibility outside the declared inputs and environment.",
				},
			},
			{
				Title: "What a reconstructable decision trail contains",
				Paragraphs: []string{
					"The exact implementation depends on the workflow, but reconstructability generally requires more than storing the final prompt and response.",
				},
				Points: []string{
					"source identifiers and versions",
					"participating human and machine identities",
					"relevant configuration and parameters",
					"ordered transformations",
					"provenance and chain-of-custody records",
					"structured verification outcomes",
					"a declared replay or review basis",
				},
			},
			{
				Title: "Start with one boundary",
				Paragraphs: []string{
					"An organisation does not need to make every AI system globally verifiable at once. A practical starting point is one workflow where the cost of losing evidence, context or reconstructability is meaningful.",
					"Define the evidence basis, make the transformations explicit, preserve the resulting state and test whether an independent reviewer can reconstruct what happened.",
				},
			},
		},
	},
	{
		Slug: "why-ai-audit-logs-are-not-enough",

		Title: "Why AI Audit Logs Are Not Enough",

		MetaTitle: "Why AI Audit Logs Are Not Enough | Baumgartner Systems",

		MetaDescription: "AI audit logs record events, but often fail to preserve the evidence, provenance, system state and transformations required to reconstruct an AI-assisted decision.",

		Description: "Audit logs can show that an event occurred. Verification requires enough evidence to reconstruct what happened, which inputs were available, what changed and under which system state.",

		Published:   "2026-09-13",
		Modified:    "2026-09-13",
		ReadingTime: "11 min read",

		Sections: []articleSection{
			{
				Title: "Audit logs are event records, not evidence systems",
				Paragraphs: []string{
					"Logging is essential operational infrastructure. It can record requests, responses, timestamps, errors, user actions and system events. But an event record is not automatically a sufficient evidence record.",
					"An AI-assisted decision may depend on documents, retrieved context, model configuration, intermediate transformations, human review and external systems. A log entry may show that a model was called without preserving enough information to reconstruct the question the model actually received or the evidence available at that moment.",
					"This distinction matters when the objective changes from monitoring a running system to explaining a past decision.",
				},
			},
			{
				Title: "A timestamp does not preserve context",
				Paragraphs: []string{
					"A timestamp answers when something happened. It does not necessarily answer what information was available, which version was used or what assumptions were active.",
					"If source material, retrieval state, policies or prompt templates later change, the timestamp alone does not recover the original decision context.",
					"Reconstructability therefore requires stable references to the relevant evidence basis rather than only chronological records of execution.",
				},
			},
			{
				Title: "Logs rarely capture the complete lineage",
				Paragraphs: []string{
					"AI workflows increasingly consist of multiple stages rather than one model call. Data may be collected, normalised, retrieved, summarised, classified, transformed and reviewed before reaching a final decision.",
					"Individual components may each produce valid logs while the relationship between those logs remains ambiguous. Verification needs to know how one state became the next.",
				},
				Points: []string{
					"Which source artefact entered the workflow?",
					"Which version or representation was used?",
					"Which transformation produced the next artefact?",
					"Which human or machine identity performed that transformation?",
					"Which result became the input to the following stage?",
					"Which final artefact informed the recorded decision?",
				},
			},
			{
				Title: "System state is part of the evidence",
				Paragraphs: []string{
					"The same apparent request can produce different outcomes when model versions, parameters, retrieval indexes, policies or runtime configuration change.",
					"A useful verification record therefore needs an explicit state basis. The exact state required depends on the verification question, but important configuration cannot remain implicit.",
					"This does not mean preserving every internal byte of a system. It means identifying which state is materially relevant to the result and retaining enough information to evaluate the same verification question later.",
				},
			},
			{
				Title: "Transformations can change meaning without looking like failures",
				Paragraphs: []string{
					"One of the most difficult problems in evidence-sensitive workflows is that information can remain technically valid while its meaning changes.",
					"A source document can become a record. The record can become a summary. The summary can become AI context. The AI analysis can become a recommendation. The recommendation can become a decision.",
					"Every stage may execute successfully. No checksum needs to fail and no database record needs to be corrupted. Yet qualification, uncertainty, emphasis or contextual relationships can disappear along the way.",
					"Traditional logs are good at showing successful execution. They are much weaker at showing whether the meaning relevant to a decision survived a sequence of transformations.",
				},
			},
			{
				Title: "Provenance is different from logging",
				Paragraphs: []string{
					"Logging describes events. Provenance describes origin and transformation history.",
					"A provenance-aware system links an artefact to the material from which it was derived and records the transitions that produced it.",
					"Chain of custody extends this by making control and transfer explicit. In workflows involving several organisations, services or reviewers, knowing who controlled an artefact can be as important as knowing when it was processed.",
				},
			},
			{
				Title: "Verification requires a declared boundary",
				Paragraphs: []string{
					"No audit system can prove everything. Verification becomes meaningful only when the system declares what is inside the verification boundary.",
					"A recorded evidence trail might verify that a particular source produced a declared state, that a transformation occurred under a defined configuration and that the resulting artefact remained intact.",
					"It does not automatically prove that the original source described reality correctly or that the final decision was ethically, legally or professionally correct.",
					"Explicit boundaries strengthen verification claims because they distinguish demonstrated properties from assumptions.",
				},
			},
			{
				Title: "What stronger AI evidence infrastructure should preserve",
				Paragraphs: []string{
					"Not every workflow requires the same evidence model. But systems designed for reconstructability typically need several layers beyond ordinary application logs.",
				},
				Points: []string{
					"stable source and artefact identities",
					"version information",
					"human, model, agent and service identities",
					"ordered lineage between workflow stages",
					"material configuration and system state",
					"recorded transformations",
					"provenance and custody information",
					"integrity evidence or tamper detection",
					"structured verification outcomes",
					"a defined basis for replay or independent review",
				},
			},
			{
				Title: "When ordinary logging is enough",
				Paragraphs: []string{
					"Not every AI workflow needs a full evidence infrastructure. Operational logs may be completely adequate for debugging, performance monitoring, availability analysis and many low-consequence automation tasks.",
					"The need for stronger evidence increases when an organisation must later explain why an outcome occurred, demonstrate which information influenced it, distinguish human from machine actions or review whether a decision was based on the expected state.",
					"The question is therefore not whether logs are useful. They are. The question is whether they preserve enough evidence for the verification obligation the organisation actually has.",
				},
			},
			{
				Title: "Can an independent reviewer reconstruct the decision?",
				Paragraphs: []string{
					"A practical test is to imagine that the original operators are unavailable six months later.",
					"Give an independent reviewer the retained records and ask them to identify the relevant source evidence, participating identities, transformations, system state and final decision path.",
					"If the reviewer can see that events occurred but cannot determine how the evidence became the outcome, the organisation has logging but not yet a reconstructable evidence trail.",
				},
			},
			{
				Title: "Start with one workflow",
				Paragraphs: []string{
					"The practical path is not to replace every existing log system. Start with one workflow where evidence loss or ambiguity would matter.",
					"Map the chain from source evidence to final outcome. Identify which transitions are already recorded, which state remains implicit and which transformations cannot currently be reconstructed.",
					"Existing logs can then become one input to a broader evidence architecture rather than being asked to provide guarantees they were never designed to provide.",
				},
			},
		},
	},
	{
		Slug: "provenance-traceability-verification-difference",

		Title: "Provenance, Traceability and Verification: What’s the Difference?",

		MetaTitle: "Provenance, Traceability and Verification: What’s the Difference? | Baumgartner Systems",

		MetaDescription: "A technical comparison of provenance, traceability and verification in AI-assisted workflows, including their boundaries, overlaps and role in reconstructable decision trails.",

		Description: "Provenance, traceability and verification are related but distinct. Together they form a stronger evidence model for reconstructing and independently reviewing AI-assisted workflows.",

		Published:   "2026-09-13",
		Modified:    "2026-09-13",
		ReadingTime: "10 min read",

		Related: []articleLink{
			{
				Title: "What Is AI Evidence Infrastructure?",
				Slug:  "ai-evidence-infrastructure",
			},
			{
				Title: "Why AI Audit Logs Are Not Enough",
				Slug:  "why-ai-audit-logs-are-not-enough",
			},
		},

		Sections: []articleSection{
			{
				Title: "Three related concepts, three different questions",
				Paragraphs: []string{
					"Provenance, traceability and verification are often used together, but they answer different questions.",
					"Provenance asks where an artefact came from. Traceability asks how it moved through a process. Verification asks whether a recorded claim, state or transformation can be independently checked.",
					"The concepts overlap because they all contribute to reconstructability, but treating them as interchangeable weakens the architecture.",
				},
			},
			{
				Title: "Provenance: where did this come from?",
				Paragraphs: []string{
					"Provenance describes origin and derivation. In an AI-assisted workflow, provenance can identify which document, dataset, model output or human contribution produced a later artefact.",
					"A provenance record is especially useful when the same information passes through several representations. It allows a reviewer to move backward from a result to the material from which that result was derived.",
				},
				Points: []string{
					"source identity",
					"source version",
					"derivation relationship",
					"creator or producing system",
					"transformation history",
					"custody or control where relevant",
				},
			},
			{
				Title: "Traceability: how did it move through the workflow?",
				Paragraphs: []string{
					"Traceability connects events and artefacts across a process. It is concerned with the path an item followed rather than only its origin.",
					"In a multi-stage AI workflow, traceability can show that a source document became a normalized representation, then retrieval context, then a model input, then an analysis and finally a decision record.",
					"Good traceability makes the sequence visible. It should be possible to identify which output from one stage became the input to the next.",
				},
			},
			{
				Title: "Verification: can the recorded claim be checked?",
				Paragraphs: []string{
					"Verification goes beyond recording origin or movement. It asks whether a specific property can be independently evaluated against a declared basis.",
					"For example, a verification result may show that a recorded artefact matches an expected digest, that a required identity was present, that a transformation followed a declared rule or that the same referentially closed inputs produce the same canonical verification outcome.",
					"Verification is therefore always scoped. It proves something about a defined boundary rather than establishing universal truth.",
				},
			},
			{
				Title: "Why the terms are often confused",
				Paragraphs: []string{
					"The three concepts frequently appear in the same architecture because they depend on overlapping records.",
					"A provenance system may contain trace information. A trace may include integrity checks. A verification system may consume provenance records. But the existence of one does not automatically imply the others.",
					"Confusion usually appears when a system records metadata and then makes a stronger claim than the metadata actually supports.",
				},
			},
			{
				Title: "Why provenance alone is not enough",
				Paragraphs: []string{
					"Knowing where an artefact came from does not necessarily explain every intermediate step that changed it.",
					"A final summary may correctly identify its source document while omitting which extraction, filtering or compression step altered the representation.",
					"Provenance is necessary for understanding origin, but reconstructability also depends on the sequence and state of the transformations.",
				},
			},
			{
				Title: "Why traceability alone is not enough",
				Paragraphs: []string{
					"A trace can show the path through a workflow without proving the integrity or correctness of the recorded transitions.",
					"A system might record that artefact A became artefact B, but unless the transformation basis is preserved, a reviewer may not be able to determine whether that transition was valid or reproducible.",
					"Traceability provides visibility. Verification provides an independent basis for checking selected properties of the trace.",
				},
			},
			{
				Title: "Why verification without provenance is weak",
				Paragraphs: []string{
					"A verification result is only as meaningful as the inputs and identities it refers to.",
					"If a system verifies that a file has a particular digest but cannot establish which source artefact that file represents, the integrity claim may be technically correct while operationally ambiguous.",
					"Provenance gives verification results context. It connects the verified state to the evidence chain that matters.",
				},
			},
			{
				Title: "A combined evidence model",
				Paragraphs: []string{
					"A stronger architecture treats provenance, traceability and verification as separate but connected layers.",
				},
				Points: []string{
					"Provenance establishes origin and derivation.",
					"Traceability establishes ordered movement through the workflow.",
					"State records the configuration and context materially relevant to each step.",
					"Verification checks declared properties against the preserved evidence basis.",
					"Boundary declarations define what the verification claim does and does not cover.",
				},
			},
			{
				Title: "A practical AI-assisted workflow example",
				Paragraphs: []string{
					"Consider a workflow in which a source document is ingested, summarized by an AI model, reviewed by a human and used in a later decision.",
					"Provenance identifies the original document and the summary derived from it. Traceability records the ordered path from ingestion through summarization and review to the decision record. Verification checks selected properties such as source identity, integrity, required review steps or reproducibility of a defined verification outcome.",
					"Together, these layers make it possible to reconstruct more than the final answer. They make the transition from evidence to outcome inspectable.",
				},
			},
			{
				Title: "What to capture technically",
				Paragraphs: []string{
					"The exact schema depends on the workflow, but several elements recur across evidence-sensitive systems.",
				},
				Points: []string{
					"stable identifiers for source and derived artefacts",
					"versions or immutable references",
					"human, model, agent and service identities",
					"ordered parent-child or input-output relationships",
					"material configuration and state",
					"timestamps where chronology matters",
					"integrity evidence",
					"structured verification results",
					"explicit verification boundaries",
				},
			},
			{
				Title: "Boundary conditions matter",
				Paragraphs: []string{
					"A provenance record does not prove that a source statement is factually true. A trace does not prove that every transformation preserved meaning. A verification result does not prove properties that were never represented inside the verification boundary.",
					"Reliable systems make these limits explicit rather than allowing metadata to imply stronger guarantees than it can support.",
				},
			},
			{
				Title: "When each concept matters most",
				Paragraphs: []string{
					"Provenance matters most when origin, derivation and custody need to be established. Traceability matters most when a workflow spans several stages, systems or actors. Verification matters most when an organisation needs to independently evaluate whether a declared condition was satisfied.",
					"In evidence-sensitive AI workflows, the strongest architecture usually combines all three.",
				},
			},
		},
	},
	{
		Slug: "what-is-a-verification-boundary",

		Title: "What Is a Verification Boundary?",

		MetaTitle: "What Is a Verification Boundary? | Baumgartner Systems",

		MetaDescription: "A technical explanation of verification boundaries in AI-assisted systems: what belongs inside the boundary, what remains external and why verification claims must stay scoped.",

		Description: "A verification boundary defines which identities, inputs, state, transformations and claims are actually covered by a verification result.",

		Published:   "2026-09-13",
		Modified:    "2026-09-13",
		ReadingTime: "10 min read",

		Related: []articleLink{
			{
				Title: "What Is AI Evidence Infrastructure?",
				Slug:  "ai-evidence-infrastructure",
			},
			{
				Title: "Provenance, Traceability and Verification: What’s the Difference?",
				Slug:  "provenance-traceability-verification-difference",
			},
			{
				Title: "Why AI Audit Logs Are Not Enough",
				Slug:  "why-ai-audit-logs-are-not-enough",
			},
		},

		Sections: []articleSection{
			{
				Title: "Verification is always about something specific",
				Paragraphs: []string{
					"A verification result is meaningful only when it is clear what was actually verified.",
					"In AI-assisted systems, this can include source identity, system state, configuration, provenance, workflow transitions, required approvals or the integrity of a recorded artefact.",
					"A verification boundary defines the set of inputs, identities, state and rules that belong to that claim.",
				},
			},
			{
				Title: "The boundary separates verified state from assumed context",
				Paragraphs: []string{
					"Every system depends on information that may exist outside the verification process. External facts, undocumented human assumptions, hidden runtime behaviour or third-party services may influence an outcome without being represented inside the evidence model.",
					"A sound verification architecture distinguishes these external dependencies from the material it can actually inspect and evaluate.",
					"Without this distinction, a narrow technical check can easily be mistaken for a much broader guarantee.",
				},
			},
			{
				Title: "What can belong inside a verification boundary",
				Paragraphs: []string{
					"The exact contents depend on the verification question. A useful boundary includes only the information required to make that question sufficiently explicit and reviewable.",
				},
				Points: []string{
					"source and artefact identities",
					"versions or immutable references",
					"human, model, agent and service identities",
					"relevant configuration and parameters",
					"declared system state",
					"ordered transformations",
					"provenance and lineage records",
					"required approvals or authority relationships",
					"integrity evidence",
					"the rule or condition being evaluated",
				},
			},
			{
				Title: "What usually remains outside the boundary",
				Paragraphs: []string{
					"A verification system should not silently claim authority over information it cannot observe or represent.",
				},
				Points: []string{
					"objective truth of external real-world statements",
					"undocumented human reasoning",
					"hidden model internals that were not captured",
					"authority that was never explicitly represented",
					"runtime behaviour outside the recorded system state",
					"ethical or legal correctness unless encoded as an explicit rule set",
				},
			},
			{
				Title: "Why verification boundaries matter in AI systems",
				Paragraphs: []string{
					"AI-assisted workflows often contain opaque or probabilistic components. That makes precise scoping especially important.",
					"A system may be able to verify that a specific model version received a specific evidence package under a declared configuration. It may also verify that the resulting artefact entered a later workflow stage unchanged.",
					"That does not mean the system has proven that the model's reasoning was correct or that the source material described reality accurately.",
				},
			},
			{
				Title: "A simple example",
				Paragraphs: []string{
					"Imagine an AI-assisted review process in which a document is uploaded, summarized, reviewed by a human and used to support a later decision.",
					"The verification boundary might include the original document identity, the exact source version, the summarization configuration, the generated summary, the reviewer identity and the final decision record.",
					"Within that boundary, the system may be able to verify which source was used, how the summary entered the workflow and whether the expected review step occurred.",
					"It still cannot automatically prove that every statement in the original document was true or that the reviewer reached the best possible decision.",
				},
			},
			{
				Title: "Referential closure makes verification stronger",
				Paragraphs: []string{
					"A verification question becomes stronger when every reference required to evaluate it is explicit inside the declared state basis.",
					"If a rule depends on a user identity, policy version, capability, ownership relationship or configuration value, those references should be represented rather than inferred from external context.",
					"This is sometimes described as making the verification input referentially closed: the verifier should not need to guess which external object or authority a reference was intended to mean.",
				},
			},
			{
				Title: "Boundaries should be declared before the result",
				Paragraphs: []string{
					"A weak architecture defines the meaning of a verification result after seeing the outcome. A stronger architecture declares the scope first.",
					"The verification question, required inputs, relevant state and expected condition should be explicit before evaluation.",
					"That makes the resulting record easier to review, replay and challenge.",
				},
			},
			{
				Title: "A verification boundary is not a security perimeter",
				Paragraphs: []string{
					"The term boundary can sound similar to a network or security perimeter, but the concepts are different.",
					"A security boundary controls access or trust between system components. A verification boundary defines which evidence and state support a particular verification claim.",
					"The two may overlap, but neither substitutes for the other.",
				},
			},
			{
				Title: "Verification claims should match the boundary",
				Paragraphs: []string{
					"The language of a verification result should never be broader than the evidence it evaluates.",
					"If the system checks integrity, say that integrity was verified. If it checks that a required identity and capability were present, say that admission conditions were satisfied.",
					"Avoid turning a bounded technical result into claims such as true, safe, compliant or correct unless those properties were explicitly represented and evaluated.",
				},
			},
			{
				Title: "How to define a boundary in practice",
				Paragraphs: []string{
					"A practical boundary definition starts with a single question rather than an entire platform.",
				},
				Points: []string{
					"State exactly what must be verified.",
					"Identify every input required to evaluate that condition.",
					"Represent identities and authority explicitly.",
					"Record the state and configuration that materially affect the result.",
					"Declare which external assumptions remain outside the system.",
					"Produce a structured result that records both the outcome and its scope.",
				},
			},
			{
				Title: "Boundaries make verification more credible",
				Paragraphs: []string{
					"Limiting a verification claim does not weaken it. In many cases, it makes the claim more credible.",
					"A system that clearly states what it can verify, what evidence supports the result and what remains outside its scope is easier to audit than one that relies on broad trust language.",
					"Good verification infrastructure is not omniscient. It is explicit about its evidence, state and limits.",
				},
			},
		},
	},
	{
		Slug: "from-evidence-to-decision-where-meaning-can-change",

		Title: "From Evidence to Decision: Where Meaning Can Change",

		MetaTitle: "From Evidence to Decision: Where Meaning Can Change | Baumgartner Systems",

		MetaDescription: "A technical look at semantic drift in AI-assisted decision chains: how evidence becomes records, summaries, AI analysis, decisions and later review.",

		Description: "Information can remain technically intact while meaning changes across summaries, transformations, AI analysis and decision records. This article examines where those shifts occur and how to make them inspectable.",

		Published:   "2026-09-13",
		Modified:    "2026-09-13",
		ReadingTime: "11 min read",

		Related: []articleLink{
			{
				Title: "What Is AI Evidence Infrastructure?",
				Slug:  "ai-evidence-infrastructure",
			},
			{
				Title: "What Is a Verification Boundary?",
				Slug:  "what-is-a-verification-boundary",
			},
			{
				Title: "Provenance, Traceability and Verification: What’s the Difference?",
				Slug:  "provenance-traceability-verification-difference",
			},
		},

		Sections: []articleSection{
			{
				Title: "Data integrity does not guarantee meaning integrity",
				Paragraphs: []string{
					"A document can remain unmodified while its meaning changes in the way it is represented, summarized or interpreted downstream.",
					"Checksums, signatures and immutable storage are valuable because they show that recorded bytes remained intact. But many decision failures occur without any corruption of the underlying data.",
					"The change happens in representation: which details are selected, which context is omitted, how uncertainty is expressed and which interpretation becomes dominant.",
				},
			},
			{
				Title: "The chain from evidence to decision",
				Paragraphs: []string{
					"A typical evidence-sensitive workflow contains several representational transitions rather than one direct step from source to outcome.",
				},
				Points: []string{
					"Evidence",
					"Record",
					"Summary",
					"AI analysis",
					"Decision",
					"Review",
				},
			},
			{
				Title: "Evidence: the original basis",
				Paragraphs: []string{
					"Evidence is the source material available to the process: documents, measurements, records, statements, observations or other artefacts.",
					"At this stage, meaning is still closely tied to the original context. The evidence may contain ambiguity, contradiction, uncertainty or qualifications that later stages compress or reorganize.",
					"A reliable evidence architecture therefore starts by preserving identity, version and provenance before interpretation begins.",
				},
			},
			{
				Title: "Record: structure already introduces selection",
				Paragraphs: []string{
					"Turning evidence into a structured record is already a transformation.",
					"Fields are selected. Categories are assigned. Free text may be normalized. Some information becomes first-class data while other information remains in notes or disappears entirely.",
					"Nothing needs to be technically incorrect for this step to change what later systems can see.",
				},
			},
			{
				Title: "Summary: compression creates semantic risk",
				Paragraphs: []string{
					"Summaries are useful precisely because they remove information. That also makes them one of the most important semantic boundaries in a decision chain.",
					"A summary may preserve the main conclusion while losing uncertainty, chronology, dissenting evidence, exceptions or conditions attached to the source.",
					"The result can be factually consistent with the original material while supporting a different interpretation.",
				},
			},
			{
				Title: "AI analysis: interpretation becomes explicit",
				Paragraphs: []string{
					"When AI analyses a record or summary, the system moves from representation toward interpretation.",
					"Model behaviour may depend on prompt structure, retrieved context, configuration, model version and system state. The same source evidence can therefore support different analyses under different conditions.",
					"For reconstructability, the analysis should remain linked to the exact evidence basis and relevant configuration that produced it.",
				},
			},
			{
				Title: "Decision: interpretation becomes consequence",
				Paragraphs: []string{
					"A decision converts analysis into an operational consequence.",
					"At this stage, the important question is not only what the AI produced, but how that output was used, weighted or overridden by human judgment and policy.",
					"A strong decision trail distinguishes source evidence, machine analysis, human interpretation and the final authority responsible for the outcome.",
				},
			},
			{
				Title: "Review: later reconstruction can introduce another interpretation",
				Paragraphs: []string{
					"Review is often treated as a neutral reconstruction of what happened. In practice, reviewers may see a different subset of records, summaries or system outputs than the original decision-makers.",
					"If the original evidence basis and transformation chain are unavailable, the review may evaluate the record of the decision rather than the information that actually produced it.",
					"That creates a second-order semantic risk: the review can become an interpretation of an interpretation.",
				},
			},
			{
				Title: "Where semantic drift can occur",
				Paragraphs: []string{
					"Semantic drift does not require malicious alteration. It can emerge naturally whenever information changes representation.",
				},
				Points: []string{
					"omission of contextual detail",
					"loss of uncertainty or qualification",
					"reordering of evidence",
					"normalization into categories",
					"summarization",
					"translation",
					"model interpretation",
					"human interpretation",
					"policy mapping",
					"later reconstruction from incomplete records",
				},
			},
			{
				Title: "Why hashes alone cannot detect semantic drift",
				Paragraphs: []string{
					"A cryptographic hash can prove that a particular artefact has not changed. It cannot prove that a later summary preserved the important meaning of the earlier artefact.",
					"If both the source and summary are individually intact, their hashes may both verify successfully while the relationship between them remains semantically weak.",
					"This is why evidence integrity and transformation verification should be treated as related but separate problems.",
				},
			},
			{
				Title: "Preserve transitions, not only artefacts",
				Paragraphs: []string{
					"A reconstructable system records more than a collection of documents. It preserves the relationships between them.",
					"For every material transition, the system should be able to identify the input, output, actor or process, relevant configuration and the type of transformation that occurred.",
					"That makes it possible to inspect where meaning may have changed rather than merely confirming that individual files exist.",
				},
			},
			{
				Title: "Semantic preservation is not absolute",
				Paragraphs: []string{
					"No technical system can guarantee preservation of every possible human interpretation.",
					"The practical objective is narrower: define which semantic properties matter to the workflow and preserve enough evidence to evaluate whether those properties survived a transformation.",
					"For one workflow, chronology may be critical. For another, uncertainty or source attribution may matter more.",
					"The verification boundary should state which properties are being evaluated rather than claiming universal semantic equivalence.",
				},
			},
			{
				Title: "A useful technical model",
				Paragraphs: []string{
					"A stronger evidence chain treats each transition as an explicit object of review.",
				},
				Points: []string{
					"identify the source artefact",
					"record the derived artefact",
					"record the transformation type",
					"record the participating identity or system",
					"capture materially relevant configuration",
					"preserve provenance and lineage",
					"record integrity evidence",
					"declare which semantic properties matter",
					"produce a structured verification outcome",
				},
			},
			{
				Title: "Meaning should be reviewable at the point of change",
				Paragraphs: []string{
					"The most useful place to detect semantic change is the transition where it occurs, not months later at the end of the process.",
					"If a summary removes a qualification, that relationship should be inspectable between source and summary. If AI analysis introduces an interpretation, the evidence basis and configuration should remain visible at that stage.",
					"This reduces the need to infer meaning retrospectively from incomplete downstream records.",
				},
			},
			{
				Title: "Start with one transformation chain",
				Paragraphs: []string{
					"A practical assessment can begin with one real workflow and map every transition from source evidence to final decision.",
					"At each step, ask what information entered, what representation left, what changed, who or what performed the transformation and whether the relationship can later be reconstructed.",
					"The result is often more revealing than examining the final AI output alone.",
				},
			},
		},
	},
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

	{
		Source: "templates/pilot.html",
		Target: "pilot/index.html",
	},

	{
		Source: "templates/museum.html",
		Target: "museum/index.html",
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
