package phase

import "os"

import "encoding/json"
import "math"
import "log"

type saveFile struct {
	PTests    []testStats `json:"PromptTests"`
	FreestyleSaves []testStats `json:"FreestyleTests"`
}

type testStats struct {
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

func SaveToFile(wrds []string, msCount int, t float64) { // decode if one exists, group tests by freestyle and prompt-based, order by characters per minute, encode to file
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
	var newSave testStats
	if msCount > -1 { // prompted test
		newSave = testStats{
			Date:   "placeholder",
			User:   "user",
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
					log.Println("sv acpm >= newsave acpm")
					f := saves.PTests[0:i]
					log.Println(f)
					l := make([]testStats, len(saves.PTests[i:len(saves.PTests)]))
					copy(l, saves.PTests[i:len(saves.PTests)])
					log.Println(l)
					f = append(f, newSave)
					log.Println(f)
					f = append(f, l...)
					log.Println(f)
					saves.PTests = f
					break
				} else if i == len(saves.PTests)-1 {
					log.Println("i == len")
					f := append(saves.PTests, newSave)
					log.Println(f)
					saves.PTests = f 
					break
				}
			}
		} else {
			log.Println("make PTests")
			saves.PTests = make([]testStats, 1)
			saves.PTests[0] = newSave
		}
	} else {
		newSave = testStats{
			Date:   "placeholder",
			User:   "user",
			Words:  wordCount,
			Runes:  runeCount,
			Missed: 0,
			Time:   t,
			Wpm:    (float64(wordCount) / t * 60),
			Awpm:   0,
			Cpm:    (float64(runeCount) / t * 60),
			Acpm:   0, // fun fact, the average length of an english word is 4.7 characters. Haven't decided how to weight characters missed
		}
	}
	file, _ = os.Create(filename)
	en := json.NewEncoder(file)
	en.SetIndent("", "  ")
	er = en.Encode(saves)
	if er != nil {
		log.Println(er)
	}

}
