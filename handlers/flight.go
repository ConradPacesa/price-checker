package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	// Importen this package to pull environment variables from .env
	_ "github.com/joho/godotenv/autoload"
)

type searchParams struct {
	Origin      string
	Destination string
	DepDate     string
	RetDate     string
	Adults      int
	Children    int
	Infants     int
	NonStop     bool
}

// Search will handle the search requests.
func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
	params := getReqParams(r)
	l := setLink(r, params)
	resp := sendRequest(l)
	io.WriteString(w, resp)
}

func getReqParams(req *http.Request) searchParams {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("There was an issue with the JSON, try again %v", err)
	}
	var s searchParams
	json.Unmarshal(body, &s)
	if err != nil {
		fmt.Printf("There was an issue with the JSON, try again %v", err)
	}
	return s
}

func setLink(req *http.Request, p searchParams) string {
	apiKey := os.Getenv("API_KEY")
	var rdate string

	if p.RetDate != "" {
		rdate = "return_date"
	}

	link := fmt.Sprintf("https://api.sandbox.amadeus.com/v1.2/flights/low-fare-search?"+
		"apikey=%v"+
		"&origin=%v"+
		"&destination=%v"+
		"&departure_date=%v"+
		"&%v=%v"+
		"&adults=%v"+
		"&children=%v"+
		"&infants=%v"+
		"&non_stop=%v",
		apiKey,
		p.Origin,
		p.Destination,
		p.DepDate,
		rdate,
		p.RetDate,
		p.Adults,
		p.Children,
		p.Infants,
		p.NonStop)
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
