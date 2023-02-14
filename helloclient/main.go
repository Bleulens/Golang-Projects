package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Fragrance string `json:"fragrance"`
	House     string
	Category  string
	Rating    string
	Longevity string
}

func main() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get("http://localhost:3333")
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	// req, err := http.NewRequest("Get", "http://localhost:3333", nil)
	// if err != nil {
	// 	fmt.Printf("Error %s", err)
	// }
	// req.Header.Add("Accept", `application/json`)
	// resp1, err := c.Do(req)
	// if err != nil {
	// 	fmt.Printf("error %s", err)
	// 	return
	// }
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response Response
	json.Unmarshal(body, &response)
	// var data map[string]interface{}
	// response, err := json.Unmarshal(body, &data);  err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Printf("%+v\n", response)

	//fmt.Printf("Body : %s", body)

	//for key, val := range Response {
	//}

	//  for i, p := range Response {
	// 	fmt.Println("Fragrance:", p.Fragrance)
	// 	fmt.Println("House:", p.House)
	// 	fmt.Println("Category:", p.Category)
	// 	fmt.Println("Rating:", p.Rating)
	// 	fmt.Println("Longevity:", p.Longevity)
	// 	fmt.Println()

	// }
	//}

}
