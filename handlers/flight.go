package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type searchParams struct {
	Origin      string
	Destination string
	DepDate     string
	RetDate     string
	Adults      int
	NonStop     bool
}

// Search will handle the search requests.
func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("There was an issue with the JSON, try again %v", err)
	}
	var s searchParams
	json.Unmarshal(body, &s)
	if err != nil {
		fmt.Printf("There was an issue with the JSON, try again %v", err)
	}
	defer r.Body.Close()

	l := setLink(r, s)
	resp := sendRequest(l)
	io.WriteString(w, resp)
}

func setLink(req *http.Request, p searchParams) string {
	const apiKey string = "API_KEY"

	link := fmt.Sprintf("https://api.sandbox.amadeus.com/v1.2/flights/low-fare-search?"+
		"apikey=%v"+
		"&origin=%v"+
		"&destination=%v"+
		"&departure_date=%v",
		apiKey,
		p.Origin,
		p.Destination,
		p.DepDate)
	fmt.Printf(link)
	return link
}

func sendRequest(link string) string {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	return bodyString
}
