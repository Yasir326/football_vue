package main

import (
	"football_teams_viewer/server/models/models.go"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello World!\n")
	}

	http.HandleFunc("/", helloHandler)
	log.Println("Listening for requests at PORT 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

}

func geTeams() {
	url := "https://api.football-data.org/v2/competitions/2021/standings?standingType=HOME"
	apiToken := ""

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("x-Auth-Token", apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		table := models.Table{}
		// Convert response body to string

	}
}
