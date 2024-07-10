package awesomesearchqueries

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// ====== NOTE: Update https://github.com/projectdiscovery/cvemap.git before planning any breaking changes ====

// Custom type that can be either a string or a slice of strings
type StringOrSlice struct {
	Value []string
}

// UnmarshalJSON customizes the unmarshaling of the JSON
func (s *StringOrSlice) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as a string slice first
	var sliceValue []string
	if err := json.Unmarshal(data, &sliceValue); err == nil {
		s.Value = sliceValue
		return nil
	}

	// If slice unmarshaling fails, try as a single string
	var stringValue string
	if err := json.Unmarshal(data, &stringValue); err == nil {
		s.Value = []string{stringValue}
		return nil
	}

	// If both fail, return an error
	return fmt.Errorf("value is neither a string nor a slice of strings")
}

type Query struct {
	Name    string        `json:"name"`
	Vendor  StringOrSlice `json:"vendor"`
	Type    string        `json:"type"`
	Engines []Engine      `json:"engines"`
}

type Engine struct {
	Platform string   `json:"platform"`
	Queries  []string `json:"queries"`
}

// LoadShodanQueries only loads queries that are for shodan
func LoadShodanQueries(filePath string) ([]Query, error) {
	var queries []Query
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&queries)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to load shodan queries from %s", filePath))
	}
	filteredQueries := []Query{}
	for _, query := range queries {
		eng := []Engine{}
		for _, engine := range query.Engines {
			if engine.Platform == "shodan" {
				eng = append(eng, engine)
			}
		}
		if len(eng) > 0 {
			query.Engines = eng
			filteredQueries = append(filteredQueries, query)
		}
	}
	return filteredQueries, nil
}
