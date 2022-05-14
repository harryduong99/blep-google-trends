package schedule

import (
	"strconv"

	"github.com/harryduong99/google-trends/trendclient"
	"github.com/robfig/cron"
)

func InitCron(numOfMinutes int, listKeywords []string, storeDBOption bool) {
	c := cron.New()
	c.AddFunc("*/"+strconv.Itoa(numOfMinutes*60)+" * * * * ", func() {
		trendclient.GetInterestOverTimeByHour(listKeywords)
		trendclient.GetInterestOverTimeByMinute(listKeywords, storeDBOption, false)
	})

	c.Start()
}
