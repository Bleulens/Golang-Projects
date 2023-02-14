package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const keyServerAddr = "serverAddr"

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	type Frag struct {
		Fragrance string `json:"fragrance"`
		House     string `json:"house"`
		Category  string `json:"category"`
		Rating    string `json:"rating"`
		Longevity string `json:"longevity"`
	}

	frag := Frag{Fragrance: "Blue de Chanel",
		House:     "Chanel",
		Category:  "Blue Fragrance",
		Rating:    "5 stars out of 5",
		Longevity: "12 plus hours"}
	// resp := map[string]interface{}{
	// 	"Fragrance": "Blue de Chanel",
	// 	"House":     "Chanel",
	// 	"Category":  "Blue Fragrance",
	// 	"Rating":    "5 stars out of 5",
	// 	"Longevity": "12 plus hours",
	//}
	// resp := make(map[string]string)
	// resp["message"] = "Status Created"
	jsonResp, err := json.Marshal(frag)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return

}
func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "Hello world!\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	log.Print("listening...")

	http.ListenAndServe(":3333", mux)
}
