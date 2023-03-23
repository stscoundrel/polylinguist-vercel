package service

import (
	"os"

	"github.com/stscoundrel/polylinguist"
	"github.com/stscoundrel/polylinguist/stats"
)

func GetStats() ([]stats.LanguageStatistic, error) {
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")

	settings := stats.Settings{
		IgnoredLanguages: []string{"SCSS", "CSS", "ASL", "HTML"},
		IgnoredRepos: []string{
			"old-norwegian-dictionary",
			"old-norwegian-dictionary-rs",
			"old-norwegian-dictionary-go",
		},
		AliasedLanguages: []stats.LanguageAlias{
			{
				Language: "C",
				Alias:    "C/C++",
			},
			{
				Language: "C++",
				Alias:    "C/C++",
			},
		},
	}

	// Fetches all languages from all repos.
	stats, err := polylinguist.GetTopLanguagesWithSettings("stscoundrel", accessToken, settings)

	return stats, err
}
