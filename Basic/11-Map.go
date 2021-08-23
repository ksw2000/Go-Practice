package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	// we can create a map by call make()
	var m = make(map[string]Vertex)

	// we can also give values when creating map
	var n = map[string]Vertex{
		"Google": {37.42202, -122.08408},
		"NCHU":   {24.12281, 120.67616},
	}

	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])
	fmt.Println(n["Google"])
	fmt.Println(n["NCHU"])

	// insert value
	m["ZSGH"] = Vertex{24.778289, 120.988108}

	// check value
	if elem, ok := m["TFG"]; !ok {
		fmt.Println("TFG is not in the map")
	} else {
		fmt.Println("TFG is at", elem)
	}

	if elem, ok := m["ZSGH"]; !ok {
		fmt.Println("ZSGH is not in the map")
	} else {
		fmt.Println("ZSGH is at", elem)
	}
}
