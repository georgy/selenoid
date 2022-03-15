package clipboard

import (
	"encoding/json"
	"log"
	"net/http"
)
import "github.com/atotto/clipboard"

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		log.Printf("[-] [CLIPBOARD] Executing request for clipboard content")
		clip, err := clipboard.ReadAll()

		if err != nil {
			log.Printf("[-] [CLIPBOARD] [ERROR] [Failed to get clipboard: %v]", err)
			err = json.NewEncoder(w).Encode(struct {
				Status int `json:"status"`
				Clip  string `json:"clipboard"`
				Error string  `json:"error"`
			}{Status: 13, Clip: clip, Error: err.Error()})
		} else {
			log.Printf("[-] [CLIPBOARD] [Got text from clipboard: %.*s]", 50, clip)
			err = json.NewEncoder(w).Encode(struct {
				Status int `json:"status"`
				Clip  string `json:"clipboard"`
				Error *string  `json:"error"`
			}{Status: 0, Clip: clip, Error: nil})
		}

		if err != nil {
			log.Printf("[-] [INIT] [Failed to encode clipboard reply to JSON: %v]", err)
		}
	}
}
