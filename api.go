package main

import "encoding/json"
import "net/http"
import "time"

func cat_fact() string{
	client:= &http.Client{Timeout: 10*time.Second}
	resp, err := client.Get("http://catfact.ninja/fact")
	if err!= nil{
		return "Failed to fetch the fact from site"
	}
	defer resp.Body.Close()

	var result struct {
		Fact string `json:"fact"`
	}
	
	if err :=json.NewDecoder(resp.Body).Decode(&result); err!=nil{
		return "failed"
	}
	return result.Fact
}
