package main

import "fmt"

func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "New York", "North America"
	case "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Uknown", "Uknown"
	}

	return region, continent
}

func main() {

	region, continent := location("LA")

	fmt.Println("Region : ", region, "Continent : ", continent)
}
