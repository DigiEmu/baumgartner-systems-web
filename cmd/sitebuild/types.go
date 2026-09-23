package main

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
	Cluster         string
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
