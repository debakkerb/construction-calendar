package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)

var (
	tranches = [...]int{60, 135, 150, 165, 180, 278, 350, 410, 480}
	restDays = map[time.Time]bool{}
	debugLogLevel = false
)

const DATE_FORMAT = "Monday, 02/01/2006"

func init() {
	bankHolidayFileLocation := flag.String("bankholiday", "bankholidays.txt", "Location of the file that contains bank holidays")
	restDaysFileLocation := flag.String("restdays", "restdays.txt", "Location of the file that contains restdays")
	debugEnabled := flag.String("debug", "false", "Enable DEBUG logging")
	flag.Parse()

	readSpecialDates(*bankHolidayFileLocation)
	readSpecialDates(*restDaysFileLocation)

	var err error
	debugLogLevel, err = strconv.ParseBool(*debugEnabled)
	if err != nil {
		log.Fatalf("Error while parsing debug logging: %v\n", err)
	}
}

func main() {
	startDate := time.Date(2022, 10, 3, 0, 0, 0, 0, time.UTC)

	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	for counter, tranch := range tranches {
		temp := addDays(&startDate, tranch)
		fmt.Fprintf(writer, "Tranch %d: %s\n", counter + 1, temp.Format(DATE_FORMAT))
	}

	writer.Flush()
}

func readSpecialDates(fileLocation string) error {
	file, err := os.Open(fileLocation)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		date, err := time.Parse("02/01/2006", scanner.Text())
		if err != nil {
			return err
		}

		restDays[date] = true
	}

	return nil
}

func addDays(startDate *time.Time, numDays int) *time.Time {
	tempDate := *startDate
	dayCounter := 0

	for dayCounter < numDays {
		if tempDate.Weekday() != time.Saturday && tempDate.Weekday() != time.Sunday && !restDays[tempDate] {
			dayCounter++

			if debugLogLevel {
				fmt.Printf("Counter: %d - %s\n", dayCounter, tempDate.Format(DATE_FORMAT))
			}
		}

		if dayCounter < numDays {
			if debugLogLevel && tempDate.Weekday() == time.Friday {
				fmt.Println()
			}
			tempDate = tempDate.Add(time.Hour * 24)
		}
	}

	return &tempDate
}
