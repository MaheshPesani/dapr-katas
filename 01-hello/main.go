package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type incomingEvent struct {
	Key string `json:"key"`
	Value int `json:"value"`
}

func main() {
	http.HandleFunc("/greeting",
		func(w http.ResponseWriter, r *http.Request) {
			resp,:= http.Get("http://localhost:8089/v1.0/state/statestore/mystate")
			defer resp.Body.Close()
			body, := ioutil.ReadAll(resp.Body)
			strVal := string(body)
			count := 0
			if strVal !="" {
				count, _ = strconv.Atoi(strVal)
				count++
			}

			stateObj := []stateData{stateData{Key: "mystate", Value: count}}
			stateData, _ := json.Marshal(stateObj)
			resp, = http.Post("http://localhost:8089/v1.0/state/statestore",
				"application/json", bytes.NewBuffer(stateData))
			if count == 1 {
				fmt.Fprint(w, "I've greeted you " +
					strconv.Itoa(count) + " time.")
			} else {
				fmt.Fprint(w, "I've greeted you " +
					strconv.Itoa(count) + " times.")
			}
		})
	log.Fatal(http.ListenAndServe(":8088", nil))
}
