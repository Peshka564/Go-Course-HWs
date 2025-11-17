package main

import (
	"bufio"
	"log"
	"math"
	"os"

	"github.com/Peshka564/Go-Course-HWs/hw1/data"
	"github.com/Peshka564/Go-Course-HWs/hw1/errors"
	"github.com/Peshka564/Go-Course-HWs/hw1/formatter"
	"github.com/Peshka564/Go-Course-HWs/hw1/httpclient"
	"github.com/joho/godotenv"
)

const maxRepos int = 10

func readUsernames() ([]string, error) {
	args := os.Args
	// The first argument is the .exe filename
	if len(args) == 1 {
		return nil, errors.FileNotProvidedError{}
	}

	filepath := args[1]
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	usernames := make([]string, 0)
	for scanner.Scan() {
		usernames = append(usernames, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return usernames, nil
}

func main() {
	usernames, err := readUsernames()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(usernames)

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	apiToken := os.Getenv("GITHUB_API_TOKEN")

	var httpClient httpclient.HttpClient
	httpClient.Init(apiToken)

	for _, username := range usernames {
		user, err := data.GetUser(username, &httpClient)
		if err != nil {
			log.Fatalf("Could not get user info for username %s because %v", username, err)
			continue
		}
		repos, err := data.GetRepos(username, &httpClient)
		if err != nil {
			log.Fatalf("Could not get user repos for username %s because %v", username, err)
			continue
		}

		repoToLanguage := make(map[string]data.LanguagesForRepo)
		// maxRepos so we don't do too many requests
		for _, repo := range repos[:int(math.Min(float64(maxRepos), float64(len(repos))))] {
			languages, err := data.GetLanguages(username, repo.Name, &httpClient)
			if err != nil {
				log.Fatalf("Could not get repo languages for username %s and repo %s because %v", username, repo.Name, err)
				continue
			}
			repoToLanguage[repo.Name] = languages
		}
		formatter.FormatAndPrintData(user, repos, repoToLanguage)
	}
}
