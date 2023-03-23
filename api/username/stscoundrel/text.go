package handler

import (
	"fmt"
	"net/http"
	"polylinguistvercel/service"
)

func TextHandler(w http.ResponseWriter, r *http.Request) {
	// Fetches all languages from all repos.
	stats, err := service.GetStats()

	if err != nil {
		fmt.Fprintf(w, "Could not fetch languages for stscoundrel")
		return
	}

	for index, language := range stats {
		fmt.Fprintf(w, "%d. %s - %f - %s \n", index+1, language.Name, language.Percentage, language.Color)
	}
}
