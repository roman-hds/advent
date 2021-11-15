package handlers

import (
	"advent/version"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// home is a simple HTTP handler function which writes a response.
func home(buildTime, commit, release string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		info := struct {
			BuildTime string `json:"buildTime"`
			Commit    string `json:"commit"`
			Release   string `json:"release"`
		}{
			version.BuildTime, version.Commit, version.Release,
		}

		fmt.Fprint(w, "Hello! Your request was processed.")

		body, err := json.Marshal(info)
		if err != nil {
			log.Printf("Could not encode info data: %v", err)
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}
