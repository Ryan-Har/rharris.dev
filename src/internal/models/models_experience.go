package models

import (
	_ "embed"
	"encoding/json"
	"log/slog"
)

type ProfessionalRole struct {
	Title   string   `json:"title"`
	Company string   `json:"company"`
	Dates   string   `json:"dates"`
	Content []string `json:"content"`
}

//go:embed experience.json
var experienceData []byte

// LoadExperienceData loads JSON data from the embedded file content
func LoadExperienceData() []ProfessionalRole {
	// Parse JSON data into []Experience
	var roles []ProfessionalRole
	err := json.Unmarshal(experienceData, &roles)
	if err != nil {
		slog.Error("error unmarshalling experience into JSON")
		return nil
	}
	return roles
}
