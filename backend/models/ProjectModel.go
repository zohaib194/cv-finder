package model

type ProjectModel struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Highlights  []string `json:"highlights"`
	Keywords    []string `json:"keywords"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	URL         string   `json:"url"`
	Roles       []string `json:"roles"`
	Entity      string   `json:"entity"`
	Type        string   `json:"type"`
}