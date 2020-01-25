package phase

import "os"

import "encoding/json"

import "log"

type saveFile struct {
	PromptSaves    []testStats `json:"PromptTests"`
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

	exstSave := saveFile{}
	filename := "save.json"

	file, er := os.Open(filename)
	if er != nil {
		log.Println(er)
		file, _ = os.Create(filename)
	}
	defer file.Close()
	er = json.NewDecoder(file).Decode(&exstSave)
	if er != nil {
		log.Println(er)
	}

	var newSave testStats
	if msCount > -1 { // prompted test
		newSave = testStats{
			Date:   "placeholder",
			User:   "user",
			Words:  wordCount,
			Runes:  runeCount,
			Missed: msCount,
			Time:   t,
			Wpm:    (float64(wordCount) / t * 60),
			Awpm:   (float64(wordCount-msCount) / t * 60),
			Cpm:    (float64(runeCount) / t * 60),
			Acpm:   ((float64(runeCount) - (float64(msCount) * 4.7)) / t * 60), // fun fact, the average length of an english word is 4.7 characters. Haven't decided how to weight characters missed
		}
		if len(exstSave.PromptSaves) > 0 {
			for i, sv := range exstSave.PromptSaves {
				if sv.Acpm >= newSave.Acpm {

					k := exstSave.PromptSaves[i:len(exstSave.PromptSaves)]
					exstSave.PromptSaves = append(exstSave.PromptSaves[0:i], newSave)
					exstSave.PromptSaves = append(exstSave.PromptSaves, k...)
					break
				} else if i == len(exstSave.PromptSaves)-1 {
					exstSave.PromptSaves = append(exstSave.PromptSaves, newSave)
					break
				}
			}
		} else {
			exstSave.PromptSaves = make([]testStats, 1)
			exstSave.PromptSaves[0] = newSave
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

	er = json.NewEncoder(file).Encode(exstSave)
	if er != nil {
		log.Println(er)
	}

}
