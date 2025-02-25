package main

import (
	"fmt"
	"strconv"
	"strings"
)

type charateristics struct {
	race      string
	alignment string
	langauge  string
}

func getCharacteristics() charateristics {
	race := chooseRace()
	alignment := chooseAlignment()
	langauge := chooseLanguage()
	return charateristics{race, alignment, langauge}
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
	info, err := getGenericResponse("races")
	if err != nil {
		fmt.Println("Error", err)
		return ""
	}
	races := info.GenericBody
	allRaces := []string{}
	fmt.Println("Please choose a race")
	for _, item := range races {
		allRaces = append(allRaces, item.Name)
		fmt.Printf("%d) %s\n", len(allRaces), item.Name)
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
	info, err := getGenericResponse("alignments")
	if err != nil {
		fmt.Println("Error", err)
		return ""
	}
	alignments := info.GenericBody
	allAlignments := []string{}
	fmt.Println("Please chhose a alignment")
	for _, item := range alignments {
		allAlignments = append(allAlignments, item.Name)
		fmt.Printf("%d) %s\n", len(allAlignments), item.Name)
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

func chooseLanguage() string {
	info, err := getGenericResponse("languages")
	if err != nil {
		fmt.Println("Error", err)
		return ""
	}
	alignments := info.GenericBody
	allLanguages := []string{}
	fmt.Println("Please chhose a secondary language")
	for _, item := range alignments {
		allLanguages = append(allLanguages, item.Name)
		fmt.Printf("%d) %s\n", len(allLanguages), item.Name)
	}
	var languageChosen string
	fmt.Scanln(&languageChosen)
	userLanguage, err := checkChosenInput(languageChosen, allLanguages)
	if err != nil {
		fmt.Println("Error", err)
		return ""
	}
	return userLanguage
}
