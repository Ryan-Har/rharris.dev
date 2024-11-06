package models

import (
	_ "embed"
	"encoding/json"
	"log/slog"
)

type Project struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Content     []string `json:"content"` //each string is a new paragraph
	Assets      []Asset  `json:"assets"`
	Tags        []string `json:"tags"`
	Links       Links    `json:"links"`
	Features    []string `json:"features"`
}

type Asset struct {
	Type        string `json:"type"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

type Links struct {
	GitHub        string `json:"github"`
	Live          string `json:"live,omitempty"`
	Documentation string `json:"documentation,omitempty"`
}

//go:embed projects.json
var projectData []byte

// LoadExperienceData loads JSON data from the embedded file content
func LoadProjectData() []Project {
	// Parse JSON data into []Experience
	var projects []Project
	err := json.Unmarshal(projectData, &projects)
	if err != nil {
		slog.Error("error unmarshalling projects into JSON", "error", err)
		return nil
	}
	return projects
}

func (p *Project) IsFeatured() bool {
	for _, tag := range p.Tags {
		if tag == "featured" {
			return true
		}
	}
	return false
}
