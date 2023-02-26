package ergast

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SendRequest(query string, data interface{}) interface{} {
	log.Println("[GET] " + query)
	resp, err := http.Get(query)

	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&data); err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	return data
}
