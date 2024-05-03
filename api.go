package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func catFact() string {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("http://catfact.ninja/fact")
	if err != nil {
		return "Failed to fetch the fact from site"
	}
	defer resp.Body.Close()

	var result struct {
		Fact string `json:"fact"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "Failed to decode fact"
	}
	return result.Fact
}
