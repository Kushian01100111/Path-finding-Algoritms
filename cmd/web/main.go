package main

import "fmt"

type application struct {
	OSM *OSMResponse
}

func main() {
	app := &application{}

	app.fetchInterpreter("https://overpass-api.de/api/interpreter?data=[out:json];way[\"highway\"](-34.615,-58.445,-34.600,-58.430);out%20geom;")

	fmt.Printf("Found %d elements\n", len(app.OSM.Elements))
	for _, elem := range app.OSM.Elements {
		fmt.Printf("Way ID: %d, Points: %d\n", elem.ID, len(elem.Geometry))
	}

}
