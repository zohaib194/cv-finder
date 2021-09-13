package model

type WorkModel struct {
	Name       string   `json:"name"`
	Position   string   `json:"position"`
	URL        string   `json:"url"`
	StartDate  string   `json:"startDate"`
	EndDate    string   `json:"endDate"`
	Summary    string   `json:"summary"`
	Highlights []string `json:"highlights"`
}