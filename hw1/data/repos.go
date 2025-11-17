package data

import (
	"math"
	"sort"
	"time"

	"github.com/Peshka564/Go-Course-HWs/hw1/httpclient"
)

type Repo struct {
	Name       string
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ForksCount int       `json:"forks_count"`
}

func GetRepos(username string, httpClient *httpclient.HttpClient) ([]Repo, error) {
	var repos []Repo
	reposRoute := "/users/" + username + "/repos"
	err := httpClient.Get(reposRoute, &repos)
	if err != nil {
		return nil, err
	}

	return repos, nil
}

func GetTotalForks(repos []Repo) int {
	totalNumForks := 0
	for _, repo := range repos {
		totalNumForks += repo.ForksCount
	}
	return totalNumForks
}

type UserActivity struct {
	Year        int
	NumActivity int
}

func aggregateUserActivityByYear(years []int) []UserActivity {
	activityMap := make(map[int]int)
	minYear := math.MaxInt
	maxYear := -1
	for _, year := range years {
		activityMap[year]++
		if year < minYear {
			minYear = year
		}
		if year > maxYear {
			maxYear = year
		}
	}

	activities := make([]UserActivity, 0)

	for year := minYear; year <= maxYear; year++ {
		activities = append(activities, UserActivity{year, activityMap[year]})
	}

	sort.Slice(activities, func(i int, j int) bool {
		return activities[i].Year > activities[j].Year
	})

	return activities
}

func GetUserActivityByYear(repos []Repo) (createdActivity []UserActivity, updatedActivity []UserActivity) {
	createdYears := make([]int, 0)
	updatedYears := make([]int, 0)
	for _, repo := range repos {
		createdYears = append(createdYears, repo.CreatedAt.Year())
		updatedYears = append(updatedYears, repo.UpdatedAt.Year())
	}

	createdActivity = aggregateUserActivityByYear(createdYears)
	updatedActivity = aggregateUserActivityByYear(updatedYears)

	return
}
