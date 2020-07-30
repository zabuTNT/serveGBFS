package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func get() (string, error) {
	url := "https://mobility.api.opendatahub.bz.it/v2/tree/BikesharingStation?limit=200&distinct=true"
	feedClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 5 secs
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "OpenMove-GBFS")

	res, getErr := feedClient.Do(req)
	if getErr != nil {
		return "", getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return "", readErr
	}

	feedTT1 := make(map[string]interface{})
	jsonErr := json.Unmarshal(body, &feedTT1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	log.Println(feedTT1)
	return "", nil
}
