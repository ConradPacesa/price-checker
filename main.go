package main

import (
	"net/http"

	"github.com/ConradPacesa/price-checker/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/search", handlers.Search)
	http.ListenAndServe(":80", nil)
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	url := setLink(r)

// 	responses := sendRequest(url)

// 	io.WriteString(w, responses)
// }

// func setLink(req *http.Request) string {
// 	const apiKey string = "JL2hRPTaG4oCAh08AxhFGg5XOmxvPAAT"

// 	fmt.Println(req.Header)

// 	var keys []string
// 	for k := range req.Header {
// 		keys = append(keys, k)
// 	}

// 	origin := strings.Join(req.Header["Org"], "")
// 	fmt.Println(origin)
// 	destination := strings.Join(req.Header["Dest"], "")
// 	fmt.Println(destination)

// 	link := fmt.Sprintf("https://api.sandbox.amadeus.com/v1.2/flights/low-fare-search?apikey=%v&origin=%v&destination=%v&departure_date=2017-12-25", apiKey, origin, destination)
// 	fmt.Printf(link)
// 	return link
// }

// func sendRequest(link string) string {
// 	resp, err := http.Get(link)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer resp.Body.Close()
// 	bodyBytes, _ := ioutil.ReadAll(resp.Body)
// 	bodyString := string(bodyBytes)

// 	return bodyString
// }
