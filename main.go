package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"log"

	"github.com/harryduong99/google-trends/databasedriver"
	"github.com/harryduong99/google-trends/schedule"
	"github.com/harryduong99/google-trends/trendclient"
	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
	databasedriver.Mongo.ConnectDatabase()
}

func main() {
	getInput()
	trendclient.GetInterestOverTimeByMinute(listKeywords, storeDBOption, true)
	schedule.InitCron(numOfMinutes, listKeywords, storeDBOption)

	select {}
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

var (
	listKeywords        []string
	defaultNumOfMinutes int = 10
	numOfMinutes        int
	storeDBOption       bool
)

func getInput() {
	getKeywords()
	getTimeToFetch()
	getStoreDatabaseOption()
}

func getKeywords() {
	fmt.Printf("Enter the keywords (for example: luna harry): ")
	var keywords string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		keywords = scanner.Text()
	}
	listKeywords = strings.Fields(keywords)
}

func getTimeToFetch() {
	fmt.Printf("Fetch data every (minutes) (default is 10): ")
	fmt.Scanln(&numOfMinutes)
	if numOfMinutes == 0 {
		numOfMinutes = defaultNumOfMinutes
	}
}

func getStoreDatabaseOption() {
	var option string
	fmt.Printf("Store result to database (1 for yes and 0 for no): ")
	fmt.Scanln(&option)

	storeDBOption = option == "1"
}
