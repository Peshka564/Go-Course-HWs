package data

import (
	"sort"

	"github.com/Peshka564/Go-Course-HWs/hw1/httpclient"
)

type LanguagesForRepo = map[string]int

func GetLanguages(username string, reponame string, httpClient *httpclient.HttpClient) (LanguagesForRepo, error) {
	var languages LanguagesForRepo
	languagesRoute := "/repos/" + username + "/" + reponame + "/languages"
	err := httpClient.Get(languagesRoute, &languages)
	if err != nil {
		return nil, err
	}

	return languages, nil
}

type LanguageData struct {
	Name string
	Freq int
}

func AggregateLanguageData(repoToLanguage map[string]LanguagesForRepo) []LanguageData {
	languagesFreq := make(LanguagesForRepo)
	for _, languagesForRepo := range repoToLanguage {
		for language, codeBytes := range languagesForRepo {
			languagesFreq[language] += codeBytes
		}
	}

	ranking := make([]LanguageData, 0, len(languagesFreq))

	for language, freq := range languagesFreq {
		ranking = append(ranking, LanguageData{language, freq})
	}

	sort.Slice(ranking, func(i int, j int) bool {
		return ranking[i].Freq > ranking[j].Freq
	})
	return ranking
}
