package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Print("Enter a time in 12-hour format (e.g. 07:05:45PM): ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		militaryTime, err := TimeConversion(input)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("%s on a 12-hour clock is %s on a 24-hour clock\n", input, militaryTime)
		}
	}
}

// func TimeConversion(s string) (string, error) {
// 	t, err := time.Parse("3:04:05PM", s)
// 	if err != nil {
// 		return "", err
// 	}

// 	return t.Format("15:04:05"), nil
// }

func TimeConversion(s string) (string, error) {
	// Split the time string into hours, minutes, seconds, and AM/PM
	timeParts := strings.Split(strings.ToUpper(s), ":")
	if len(timeParts) != 3 {
		return "", fmt.Errorf("invalid time string: %s", s)
	}

	hours, err := strconv.Atoi(timeParts[0])
	if err != nil || hours < 1 || hours > 12 {
		return "", fmt.Errorf("invalid hours: %s", timeParts[0])
	}

	minutes, err := strconv.Atoi(timeParts[1])
	if err != nil || minutes < 0 || minutes > 59 {
		return "", fmt.Errorf("invalid minutes: %s", timeParts[1])
	}

	timeSeconds := strings.Split(timeParts[2], "PM")
	if len(timeSeconds) == 1 {
		timeSeconds = strings.Split(timeParts[2], "AM")
	}

	seconds, err := strconv.Atoi(strings.TrimSpace(timeSeconds[0]))
	if err != nil || seconds < 0 || seconds > 59 {
		return "", fmt.Errorf("invalid seconds: %s", timeSeconds[0])
	}

	ampm := strings.TrimSpace(timeParts[2])[len(timeSeconds[0]):]
	if ampm != "AM" && ampm != "PM" {
		return "", fmt.Errorf("invalid AM/PM indicator: %s", ampm)
	}

	// Convert hours to military time
	if ampm == "PM" && hours != 12 {
		hours += 12
	} else if ampm == "AM" && hours == 12 {
		hours = 0
	}

	// Format the time string in military time
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds), nil
}
