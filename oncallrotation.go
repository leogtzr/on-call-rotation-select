package main

import (
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())

	//dt := currentDate()
	//fmt.Println(dt)

	//dt = dt.AddDate(0, 0, 7)
	//fmt.Println(dt)

	// Some code ...
	//o := OnCallPerson{"Leo", MEX}
	//fmt.Println(o)

	//initialRotationDate := initialRotationDate()
	//fmt.Println(initialRotationDate)

	//fmt.Println(buildUSAHolidays())
	//fmt.Println(buildMEXHolidays())

	//fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	/*
		team := nonRandomizedTeam()
		for i := range team {
			j := rand.Intn(i + 1)
			team[i], team[j] = team[j], team[i]
		}

		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~>")

		for _, p := range team {
			fmt.Println(p.Name, p.Location)
		}

		fmt.Println(int(time.Sunday), time.Sunday)
		fmt.Println(int(time.Monday), time.Monday)
		fmt.Println(int(time.Tuesday), time.Tuesday)
		fmt.Println(int(time.Wednesday), time.Wednesday)
		fmt.Println(int(time.Thursday), time.Thursday)
		fmt.Println(int(time.Friday), time.Friday)
		fmt.Println(int(time.Saturday), time.Saturday)
		fmt.Println(int(time.Sunday), time.Sunday)
	*/
}
