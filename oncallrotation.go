package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())

	rotation := onCallShift()

	mx := normalizeHolidayBasedOnCurrentYear(buildMEXHolidays())
	usa := normalizeHolidayBasedOnCurrentYear(buildUSAHolidays())

	for _, shift := range rotation {

		if is, holiday := IsHolidayWithinShiftEstrict(mx, shift.Date); is && shift.Location == MEX {
			fmt.Println(shift, " <====> ", holiday)
		}

		if is, holiday := IsHolidayWithinShiftEstrict(usa, shift.Date); is && shift.Location == USA {
			fmt.Println(shift, " <====> ", holiday)
		}

		fmt.Println("Shift: ", shift)

	}

	fmt.Println("Fin ... ")

}
