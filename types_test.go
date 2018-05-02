package main

import (
	"fmt"
	"testing"
)

func Test_teamShiftsOccurrencesCount(t *testing.T) {
	occurrences := teamShiftsOccurrencesCount(nonRandomizedTeam())
	if occurrences == nil {
		t.Error("occurrences is nil ... ")
	}
	for _, v := range occurrences {
		if v != 0 {
			t.Fatal("Initial values should be zero ... ")
		}
	}
}

func Test_getRandomTeamMember(t *testing.T) {
	team := nonRandomizedTeam()
	teamShiftCounts := teamShiftsOccurrencesCount(team)
	randomTeamMember := getRandomTeamMember(teamShiftCounts)
	fmt.Println(randomTeamMember)
}

func Test_getRandomTeamWithLocation(t *testing.T) {
	team := nonRandomizedTeam()
	teamShiftCounts := teamShiftsOccurrencesCount(team)
	randomTeamMember := getRandomTeamWithLocation(teamShiftCounts, USA)
	if randomTeamMember.Location != USA {
		t.Error("Error, expecting USA team member.")
	}
}
