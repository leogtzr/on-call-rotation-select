package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())

	/*
	rotation := Shift()
	for _, shift := range rotation {
		fmt.Println(shift.String())
	}
	*/

		rotation := Shift()

		mx := normalizeHolidayBasedOnCurrentYear(buildMEXHolidays())
		usa := normalizeHolidayBasedOnCurrentYear(buildUSAHolidays())

		for _, shift := range rotation {

			if is, holiday := IsHolidayWithinShiftEstrict(mx, shift.Date); is && shift.Location == MEX {
				fmt.Println("\t", shift.String(), " <====> ", holiday)
			}

			if is, holiday := IsHolidayWithinShiftEstrict(usa, shift.Date); is && shift.Location == USA {
				fmt.Println("\t", shift.String(), " <====> ", holiday)
			}

			fmt.Println(shift.String())

		}

}
