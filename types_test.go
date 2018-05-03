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

func Test_averageNumberOfShiftsPerPerson(t *testing.T) {
	team := Team{
		OnCallPerson{Name: "Shxin", Location: USA},
		OnCallPerson{Name: "Abhi", Location: USA},
		OnCallPerson{Name: "Paulina", Location: USA},
		OnCallPerson{Name: "Brodr", Location: USA},
	}
	teamShiftCounts := teamShiftsOccurrencesCount(team)

	maxNumOfRotations := maxNumberOfRotations(WeeksPerYear, team)
	fmt.Println("Here: ", maxNumOfRotations)

	averageNumberOfRotations := averageNumberOfShiftsPerPerson(teamShiftCounts, maxNumOfRotations)
	fmt.Println(averageNumberOfRotations)

}

func Test_smallest(t *testing.T) {
	team := Team{
		OnCallPerson{Name: "Shxin", Location: USA},
		OnCallPerson{Name: "Abhi", Location: USA},
		OnCallPerson{Name: "Paulina", Location: USA},
		OnCallPerson{Name: "Brodr", Location: USA},
	}
	teamShiftCounts := teamShiftsOccurrencesCount(team)
	teamShiftCounts[team[0]] = 3
	teamShiftCounts[team[1]] = 56
	teamShiftCounts[team[2]] = 5
	teamShiftCounts[team[3]] = 4556

	small, _ := smallest(teamShiftCounts)
	expected := 3
	if small != expected {
		t.Errorf("Got: %d, expected: %d", small, expected)
	}

}

func Test_everybodyHadSameShifts(t *testing.T) {
	team := Team{
		OnCallPerson{Name: "Shxin", Location: USA},
		OnCallPerson{Name: "Abhi", Location: USA},
		OnCallPerson{Name: "Paulina", Location: USA},
		OnCallPerson{Name: "Brodr", Location: USA},
	}
	teamShiftCounts := teamShiftsOccurrencesCount(team)
	teamShiftCounts[team[0]] = 2
	teamShiftCounts[team[1]] = 2
	teamShiftCounts[team[2]] = 2
	teamShiftCounts[team[3]] = 2

	fmt.Println(teamShiftCounts)

	small, _ := smallest(teamShiftCounts)

	if ok, _ := everybodyHadSameShifts(teamShiftCounts, small); !ok {
		t.Error("Expecting true ... ")
	}
}
