package main

import (
	"fmt"
	"strings"
	"time"
)

/*
Program to convert 12 Hour time format to 24 Hour format.
Example:
"10:27:50 AM" to "10:27:50"
"10:27:50 PM" to "22:27:50"

"12:00:00 PM" -> "12:00:00"
"12:00:00 AM" -> "00:00:00"
*/

// using the standard time library func
func convertTo24HourFormat(timeIn string) {

	layoutLongForm := "03:04:05 PM" // standard time as per time lib
	layoutDesiredFormat := "15:04:05"

	t, err := time.Parse(layoutLongForm, timeIn)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t.Format(layoutDesiredFormat))
}

func convertTo24HourFormat2(timeIn string) {

	ss := strings.Split(timeIn, " ")
	timeSplit := strings.Split(ss[0], ":")

	hh := int((timeSplit[0][0]-'0')*10 + (timeSplit[0][1]-'0')%10)

	if strings.Contains(timeIn, "AM") {
		if hh == 12 {
			hh = 0
			timeSplit[0] = fmt.Sprintf("%02d", hh)
		}
	} else { // PM
		if hh < 12 {
			hh += 12
			timeSplit[0] = fmt.Sprintf("%02d", hh)
		}
	}

	fmt.Printf("%s:%s:%s\n", timeSplit[0], timeSplit[1], timeSplit[2])
}
