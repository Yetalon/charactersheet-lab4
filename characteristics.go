package main

import (
	"fmt"
	"strconv"
	"strings"
)

type charateristics struct {
	race      string
	alignment string
}

func getCharacteristics() charateristics {
	race := chooseRace()
	alignment := chooseAlignment()
	return charateristics{race, alignment}
}

func checkChosenInput(input string, weapons []string) (string, error) {
	num, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return "", err
	}
	if num > len(weapons) {
		return "", fmt.Errorf("invalid input")
	}
	choice := weapons[num-1]
	return choice, nil
}

func chooseRace() string {
	info, err := getraces()
	if err != nil {
		fmt.Println("Error", err)
		return ""
	}
	races := info.GenericBody
	allRaces := []string{}
	fmt.Println("Please choose a race")
	for i, item := range races {
		allRaces = append(allRaces, item.Name)
		fmt.Printf("%d) %s\n", i+1, item.Name)
	}
	var raceChosen string
	fmt.Scanln(&raceChosen)
	usersRace, err := checkChosenInput(raceChosen, allRaces)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return usersRace
}

func chooseAlignment() string {
	info, err := getAlignments()
	if err != nil {
		fmt.Println("Error", err)
		return ""
	}
	alignments := info.GenericBody
	allAlignments := []string{}
	fmt.Println("Please chhose a alignment")
	for i, item := range alignments {
		allAlignments = append(allAlignments, item.Name)
		fmt.Printf("%d) %s\n", i+1, item.Name)
	}
	var alignmentChosen string
	fmt.Scanln(&alignmentChosen)
	userAlignment, err := checkChosenInput(alignmentChosen, allAlignments)
	if err != nil {
		fmt.Println("Error", err)
		return ""
	}
	return userAlignment
}
