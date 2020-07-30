// serveGBFS project main.go
package main

import (
	"encoding/gob"
	"encoding/json"
	_ "fmt"
	"log"
	_ "math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

func saveCache() {
	dataFile, err := os.Create("cache.gob")

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// serialize the data
	dataEncoder := gob.NewEncoder(dataFile)
	dataEncoder.Encode(c.Items())

	dataFile.Close()
}

func getCache() (map[string]cache.Item, error) {
	var data map[string]cache.Item

	// open data file
	dataFile, err := os.Open("cache.gob")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&data)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	dataFile.Close()
	return data, nil

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Nothing to see here", http.StatusForbidden)
	return
}

func autodiscoverHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	now := time.Now() // current local time
	sec := now.Unix()

	profile := GbfsMain{int(sec), 0, "2.0", System{}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func gbfsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	now := time.Now() // current local time
	sec := now.Unix()

	id := strings.TrimPrefix(r.URL.RequestURI(), "/gbfs/")

	log.Println("Hello..." + id)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	//c.Set("gbfs_versions", rand.Intn(99-1)+1, cache.NoExpiration)

	profile := GbfsMain{int(sec), 0, "2.0", System{}}
	c.Set(id, profile, cache.NoExpiration)

	foo, found := c.Get(id)
	if found {
		log.Println(foo)
	} else {
		log.Println("Not found in cache")
	}

	js, err := json.Marshal(foo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	saveCache()
}

var c *cache.Cache

func main() {

	data, err := getCache()
	if err != nil {
		log.Println("Cache not found")
		c = cache.New(5*time.Minute, 10*time.Minute)
	} else {
		c = cache.NewFrom(5*time.Minute, 10*time.Minute, data)
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/gbfs.json", autodiscoverHandler)
	http.HandleFunc("/gbfs/gbfs_versions", gbfsHandler)
	http.HandleFunc("/gbfs/system_information", gbfsHandler)
	http.HandleFunc("/gbfs/station_information", gbfsHandler)
	http.HandleFunc("/gbfs/station_status", gbfsHandler)
	http.HandleFunc("/gbfs/free_bike_status", gbfsHandler)
	http.HandleFunc("/gbfs/system_hours", gbfsHandler)
	http.HandleFunc("/gbfs/system_calendar", gbfsHandler)
	http.HandleFunc("/gbfs/system_regions", gbfsHandler)
	http.HandleFunc("/gbfs/system_pricing_plans", gbfsHandler)
	http.HandleFunc("/gbfs/system_alerts", gbfsHandler)

	log.Print("GBFS server started")
	http.ListenAndServe(":"+port, nil)
}
