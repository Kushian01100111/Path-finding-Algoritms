package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type OSMResponse struct {
	Elements []struct {
		Type     string  `json:"type"`
		ID       int64   `json:"id"`
		Lat      float64 `json:"lat,omitempty"`
		Lon      float64 `json:"lon,omitempty"`
		Nodes    []int64 `json:"nodes,omitempty"`
		Geometry []struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"geometry,omitempty"`
	} `json:"elements"`
}

func (app *application) fetchInterpreter(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var data OSMResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	app.OSM = &data
	return nil
}
