package main

import (
	"encoding/json"
	"fmt"
	"github.com/JosephZoeller/GoTyping/cmd/phase"
	"html/template"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type SaveFile struct {
	PTests []TestResults `json:"PromptTests"`
}

type TestResults struct {
	Date   string  `json:"Date"`
	User   string  `json:"User"`
	Words  int     `json:"WordCount"`
	Runes  int     `json:"RuneCount"`
	Missed int     `json:"MissedCount"`
	Time   float64 `json:"Time"`
	Wpm    float64 `json:"WordsPerMinute"`
	Awpm   float64 `json:"AdjustedWordsPerMinute"`
	Cpm    float64 `json:"CharactersPerMinute"`
	Acpm   float64 `json:"AdjustedCharactersPerMinute"`
}

const savefilename string = "save.json"

func main() {
	hostSave()
}

func loadFile() (*SaveFile, error) {
	saves := SaveFile{}

	file, er := os.Open(savefilename)
	if er != nil {
		return &saves, er
	} else {
		defer file.Close()
	}

	er = json.NewDecoder(file).Decode(&saves)
	if er != nil {
		return &saves, er
	}

	return &saves, nil
}

func saveFile(s *SaveFile) error {
	file, er := os.Create(savefilename)
	if er != nil {
		return er
	}
	defer file.Close()

	en := json.NewEncoder(file)
	en.SetIndent("", "  ")
	er = en.Encode(*s)
	if er != nil {
		return er
	}
	return nil
}

func hostSave() {
	http.HandleFunc("/Display", func(res http.ResponseWriter, req *http.Request) {
		saves, er := loadFile()
		if er != nil {
			log.Println(er)
		}

		log.Println("Displaying Content")
		t, _ := template.ParseFiles("./web/tables.html")
		t.Execute(res, *saves)
	})

	http.HandleFunc("/Upload", func(res http.ResponseWriter, req *http.Request) {
		up := phase.TestUpload{}

		log.Println("Upload Requested")
		if req.Method != "POST" {
			log.Println("Upload Rejected")
		} else {
			err := json.NewDecoder(req.Body).Decode(&up)

			if err != nil {
				log.Println(err)
			} else {
				AppendSave(up)
			}
		}
	})

	errorChan := make(chan error)
	fmt.Println("Listening on port 8080 (http)...")
	go func() {
		errorChan <- http.ListenAndServe(":8080", nil)
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	for {
		select {
		case err := <-errorChan:
			if err != nil {
				log.Fatalln(err)
			}

		case sig := <-signalChan:
			log.Println("shutting down: ", sig)
			os.Exit(0)
		}
	}

}

func AppendSave(tup phase.TestUpload) {
	wordCount := len(tup.Words)
	runeCount := phase.GetByteCount(tup.Words)

	saves, er := loadFile()
	if er != nil {
		log.Println(er)
	}

	var newSave TestResults
	newSave = TestResults{
		Date:   tup.Date,
		User:   tup.User,
		Words:  wordCount,
		Runes:  runeCount,
		Missed: tup.Missed,
		Time:   math.Round(tup.Time*100) / 100,
		Wpm:    math.Round(float64(wordCount)/tup.Time*6000) / 100,
		Awpm:   math.Round(float64(wordCount-tup.Missed)/tup.Time*6000) / 100,
		Cpm:    math.Round(float64(runeCount)/tup.Time*6000) / 100,
		Acpm:   math.Round((float64(runeCount)-(float64(tup.Missed)*4.7))/tup.Time*6000) / 100, // fun fact, the average length of an english word is 4.7 characters. Haven't decided how to weight characters missed
	}

	if len(saves.PTests) > 0 {
		for i, sv := range saves.PTests {
			if sv.Acpm >= newSave.Acpm {
				f := saves.PTests[0:i]
				l := make([]TestResults, len(saves.PTests[i:len(saves.PTests)]))
				copy(l, saves.PTests[i:len(saves.PTests)])
				f = append(f, newSave)
				f = append(f, l...)
				saves.PTests = f
				break
			} else if i == len(saves.PTests)-1 {
				f := append(saves.PTests, newSave)
				saves.PTests = f
				break
			}
		}
	} else {
		saves.PTests = make([]TestResults, 1)
		saves.PTests[0] = newSave
	}

	er = saveFile(saves)
	if er != nil {
		log.Println(er)
	} else {
		log.Println("Content Updated")
	}

}
