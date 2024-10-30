package awesomesearchqueries

import (
	"embed"
)

//go:embed QUERIES.json wordpress/wp-plugins.txt wordpress/wp-themes.txt
var embeddedFiles embed.FS

// GetQueries returns the content of QUERIES.json file
func GetQueries() ([]byte, error) {
	return embeddedFiles.ReadFile("QUERIES.json")
}

// GetWordPressPlugins returns the content of wordpress/wp-plugins.txt file
func GetWordPressPlugins() ([]byte, error) {
	return embeddedFiles.ReadFile("wordpress/wp-plugins.txt")
}

// GetWordPressThemes returns the content of wordpress/wp-themes.txt file
func GetWordPressThemes() ([]byte, error) {
	return embeddedFiles.ReadFile("wordpress/wp-themes.txt")
}
