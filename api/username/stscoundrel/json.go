package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"polylinguistvercel/service"
)

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	// Fetches all languages from all repos.
	stats, err := service.GetStats()

	if err != nil {
		fmt.Fprintf(w, "Could not fetch languages for stscoundrel")
		return
	}

	jsonResult, _ := json.Marshal(stats)

	fmt.Fprintf(w, string(jsonResult))
}
