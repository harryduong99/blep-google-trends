package trendclient

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/groovili/gogtrends"
	"github.com/harryduong99/google-trends/models"
	"github.com/harryduong99/google-trends/repository/keyword_repository"
	"github.com/harryduong99/google-trends/repository/trend_repository"
	"github.com/pkg/errors"
)

const (
	locUS  = "US"
	catAll = "all"
	langEn = "EN"
)

func GetInterestOverTimeByMinute(listKeywords []string, storeDBOption bool, isFirstTime bool) {
	log.Println("Explore Search by minute:")
	for _, keyword := range listKeywords {
		log.Println("Explore trends for keyword:", keyword)

		overTimeHour := fetching(keyword, storeDBOption, "now 7-d")
		overTime := fetching(keyword, storeDBOption, "now 4-H")
		resultHour := handleItems(overTimeHour)
		resultMinute := handleItems(overTime)

		if !isFirstTime {
			resultMinute = calculate(resultHour, resultMinute)
		}
		printResult(keyword, resultMinute)
		if storeDBOption {
			storeData(keyword, resultMinute)
		}
	}
}

func printResult(keyword string, resultMinute map[string]int) {
	for time, value := range resultMinute {
		fmt.Printf("%v: %#v => %#v \n", keyword, time, value)
	}

}
func calculate(resultHour map[string]int, resultMinute map[string]int) map[string]int {
	var hour string
	var splited []string
	var targetHour string
	for time, value := range resultMinute {
		splited = strings.Split(time, " ")
		hour = strings.Split(splited[4], ":")[0]
		splited[4] = hour + ":00"
		targetHour = strings.Join(splited[:], " ")
		resultMinute[time] = resultHour[targetHour] * value / 100
	}

	return resultMinute
}

func GetInterestOverTimeByHour(listKeywords []string) {
	log.Println("Explore Search by hour:")
	for _, keyword := range listKeywords {
		go fetching(keyword, false, "now 7-d")
	}
}

func fetching(keyword string, storeDBOption bool, timeRange string) []*gogtrends.Timeline {
	ctx := context.Background()

	// get widgets for Golang keyword in programming category
	explore, err := gogtrends.Explore(ctx, &gogtrends.ExploreRequest{
		ComparisonItems: []*gogtrends.ComparisonItem{
			{
				Keyword: keyword,
				Geo:     locUS,
				Time:    timeRange,
			},
		},
		Category: 0, // All category
		Property: "",
	}, langEn)

	handleError(err, "Failed to explore widgets")

	log.Println("Interest over time:")
	overTime, err := gogtrends.InterestOverTime(ctx, explore[0], langEn)
	handleError(err, "Failed in call interest over time")

	return overTime
}

// func storeData(keyword string, overTime []*gogtrends.Timeline) {
func storeData(keyword string, resultMinute map[string]int) {
	var keywordModel models.Keyword
	keywordModel.Keyword = keyword
	keyword_repository.KeywordRepository.Store(keywordModel)
	// ref := reflect.ValueOf(overTime)

	var trend models.Trend
	// for i := 0; i < ref.Len(); i++ {
	for time, value := range resultMinute {
		trend.KeywordID = keyword_repository.KeywordRepository.GetKeyword(keyword).ID
		trend.Time = time
		trend.Value = value
		trend_repository.TrendRepository.Store(trend)
	}

}

func handleError(err error, errMsg string) {
	if err != nil {
		log.Fatal(errors.Wrap(err, errMsg))
	}
}

func handleItems(items interface{}) map[string]int {
	result := make(map[string]int)

	var time string
	var value int
	ref := reflect.ValueOf(items)

	if ref.Kind() != reflect.Slice {
		log.Fatalf("Failed to print %s. It's not a slice type.", ref.Kind())
	}

	for i := 0; i < ref.Len(); i++ {
		time = fmt.Sprint(reflect.Indirect(ref.Index(i)).FieldByIndex([]int{1}).Interface())
		value, _ = strconv.Atoi(fmt.Sprint(reflect.Indirect(ref.Index(i)).FieldByIndex([]int{5}).Index(0).Interface()))
		result[time] = value
	}

	return result
}
