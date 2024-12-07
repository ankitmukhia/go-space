package main

import "fmt"

type Key struct {
	Path, Country string
}	
	
func main() {
	hits := make(map[Key]int)

	addVisit(hits, "/doc/", "au")  
	addVisit(hits, "/doc/", "us") 
	addVisit(hits, "/home/", "vn")
	addVisit(hits, "/doc/", "au")  
	addVisit(hits, "/home/", "vn")

	printVisits(hits)
}

func addVisit(hits map[Key]int, path, country string) {
	key := Key{Path: path, Country: country}
	/*
	* key = { "/doc/", "au" }
	* hits[key]++ Increment the visit count for given key
	* count is increasing with ++ 
	* maps in go are refrence types means any change we make to this hits, will be reflected to
	* original map, such that we don't have to explicitly return.
	*/
	hits[key]++
}

func printVisits(hits map[Key]int){
	for key, count := range hits {
		fmt.Printf("Page: %s, Country: %s, Visits: %d\n", key.Path, key.Country, count)
	}
}
