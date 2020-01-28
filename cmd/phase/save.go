package phase

import (
	"encoding/json"
	"log"
	"math"
	"os"
)

type SaveFile struct {
	PTests []testResults `json:"PromptTests"`
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

const savefilename string = "../save.json"

func AppendSave(date, user string, wrds []string, msCount int, t float64) {
	wordCount := len(wrds)
	runeCount := getByteCount(wrds)

	saves, er := GetSave()
	if er != nil {
		log.Println(er)
	} else {

		var newSave testResults
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

		er = setSave(saves)
		if er != nil {
			log.Println(er)
		}
	}
}

func setSave(saves *SaveFile) error {
	file, _ := os.Create(savefilename)
	en := json.NewEncoder(file)

	en.SetIndent("", "  ")
	er := en.Encode(*saves)
	if er != nil {
		return er
	}
	return nil
}

func GetSave() (*SaveFile, error) {
	saves := SaveFile{}

	file, er := os.Open(savefilename)
	if er != nil {
		return nil, er
	}

	er = json.NewDecoder(file).Decode(&saves)
	if er != nil {
		return nil, er
	}

	return &saves, nil
}
