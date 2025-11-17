package formatter

import (
	"os"
	"strconv"

	"github.com/Peshka564/Go-Course-HWs/hw1/data"
	"github.com/olekukonko/tablewriter"
)

func FormatAndPrintData(user *data.UserData, repos []data.Repo, repoToLanguage map[string]data.LanguagesForRepo) {
	userTable := tablewriter.NewWriter(os.Stdout)
	userTable.Header([]string{"Username", "Num Repos", "Num Followers", "Total Number of Forks"})
	userTable.Append([]string{user.Login, strconv.Itoa(user.PublicRepos), strconv.Itoa(user.Followers), strconv.Itoa(data.GetTotalForks(repos))})
	userTable.Render()

	languagesTable := tablewriter.NewWriter(os.Stdout)
	languagesTable.Header([]string{"Language", "Num Bytes Total"})
	rankedLanguages := data.AggregateLanguageData(repoToLanguage)
	for _, ranking := range rankedLanguages {
		languagesTable.Append([]string{ranking.Name, strconv.Itoa(ranking.Freq)})
	}
	languagesTable.Render()

	createdActivity, updatedActivity := data.GetUserActivityByYear(repos)

	reposCreatedTable := tablewriter.NewWriter(os.Stdout)
	reposCreatedTable.Header([]string{"Year", "Repositories Created"})
	for _, activity := range createdActivity {
		reposCreatedTable.Append([]string{strconv.Itoa(activity.Year), strconv.Itoa(activity.NumActivity)})
	}
	reposCreatedTable.Render()

	reposUpdatedTable := tablewriter.NewWriter(os.Stdout)
	reposUpdatedTable.Header([]string{"Year", "Repositories Updated"})
	for _, activity := range updatedActivity {
		reposUpdatedTable.Append([]string{strconv.Itoa(activity.Year), strconv.Itoa(activity.NumActivity)})
	}
	reposUpdatedTable.Render()
}
