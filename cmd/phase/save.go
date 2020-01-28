package phase

import (
	"bytes"
	"log"
	"encoding/json"
	"net/http"
)

// post save data for typetestd to process.
type TestUpload struct {
	Date   string   `json:"Date"`
	User   string   `json:"User"`
	Words  []string `json:"Words"`
	Missed int      `json:"MissedCount"`
	Time   float64  `json:"Time"`
}

func SendSave(d, u string, wrds []string, msCount int, t float64) {
	up := TestUpload{
		Date:   d,
		User:   u,
		Words:  wrds,
		Missed: msCount,
		Time:   t,
	}

	js, err := json.Marshal(up)
	if err != nil {
		log.Println("[Save]: " + err.Error())
		return
	}
	
	res, err := http.Post("http://localhost:8080/Upload", "application/json", bytes.NewBuffer(js))
	if err != nil {
		log.Println("[Save]: " + err.Error())
	} else if res.StatusCode != 200 {
		log.Println("[Save]]: " + res.Status)
	}
}
