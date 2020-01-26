package phase

import (
	"encoding/json"
	"html/template"
	"log"
	"math"
	"net/http"
	"os"
)

type saveFile struct {
	PTests []testResults `json:"PromptTests"`
	FTests []testResults `json:"FreestyleTests"`
}

type testResults struct {
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

func SaveToFile(date, user string, wrds []string, msCount int, t float64) { // decode if one exists, group tests by freestyle and prompt-based, order by characters per minute, encode to new file
	wordCount := len(wrds)
	runeCount := getByteCount(wrds)

	saves := saveFile{}
	filename := "save.json"

	file, er := os.Open(filename)
	if er != nil {
		log.Println(er)
	} else {
		er = json.NewDecoder(file).Decode(&saves)
		if er != nil {
			log.Println(er)
		} else {
			file.Close()
		}
	}
	var newSave testResults
	if msCount > -1 { // prompted test
		newSave = testResults{
			Date:   date,
			User:   user,
			Words:  wordCount,
			Runes:  runeCount,
			Missed: msCount,
			Time:   math.Round(t*100) / 100,
			Wpm:    math.Round(float64(wordCount)/t*6000) / 100,
			Awpm:   math.Round(float64(wordCount-msCount)/t*6000) / 100,
			Cpm:    math.Round(float64(runeCount)/t*6000) / 100,
			Acpm:   math.Round((float64(runeCount)-(float64(msCount)*4.7))/t*6000) / 100, // fun fact, the average length of an english word is 4.7 characters. Haven't decided how to weight characters missed
		}
		if len(saves.PTests) > 0 {
			for i, sv := range saves.PTests {
				if sv.Acpm >= newSave.Acpm {
					f := saves.PTests[0:i]
					l := make([]testResults, len(saves.PTests[i:len(saves.PTests)]))
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
			saves.PTests = make([]testResults, 1)
			saves.PTests[0] = newSave
		}
	} else {
		newSave = testResults{
			Date:   date,
			User:   user,
			Words:  wordCount,
			Runes:  runeCount,
			Missed: 0,
			Time:   math.Round(t*100) / 100,
			Wpm:    math.Round(float64(wordCount)/t*6000) / 100,
			Awpm:   0,
			Cpm:    math.Round(float64(runeCount)/t*6000) / 100,
			Acpm:   0,
		}
		if len(saves.FTests) > 0 {
			for i, sv := range saves.FTests {
				if sv.Cpm >= newSave.Cpm {
					f := saves.FTests[0:i]
					l := make([]testResults, len(saves.FTests[i:len(saves.FTests)]))
					copy(l, saves.FTests[i:len(saves.FTests)])
					f = append(f, newSave)
					f = append(f, l...)
					saves.FTests = f
					break
				} else if i == len(saves.FTests)-1 {
					f := append(saves.FTests, newSave)
					saves.FTests = f
					break
				}
			}
		} else {
			saves.FTests = make([]testResults, 1)
			saves.FTests[0] = newSave
		}
	}
	file, _ = os.Create(filename)
	en := json.NewEncoder(file)
	en.SetIndent("", "  ")
	er = en.Encode(saves)
	if er != nil {
		log.Println(er)
	}
	doHttpStuff(saves)
}

func doHttpStuff(saves saveFile) {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("../../web/tables.html")
		t.Execute(res, saves)
	})
	
	http.ListenAndServe(":8080", nil)
}
