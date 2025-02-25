package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func getAbilitiesScores() (map[string]int, error) {
	return chooseAbilities()
}

func chooseAbilities() (map[string]int, error) {
	resp, err := getGenericResponse("ability-scores")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	abilites := resp.GenericBody
	fmt.Println("Would you like to \n1) Roll your stats \n2) Manually input them")
	var choice string
	fmt.Scanln(&choice)
	numChoice, err := strconv.Atoi(choice)
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}

	var allAbilites map[string]int
	if numChoice == 1 {
		allAbilites = RollStats(abilites)
	} else {
		allAbilites = ManualStats(abilites)
	}
	return allAbilites, nil
}

func RollStats(abilities []GenericName) map[string]int {
	allAbilites := make(map[string]int)
	for _, item := range abilities {
		rnd := 3 + rand.Intn(14)
		allAbilites[item.Name] = rnd
		fmt.Printf("%s: %d\n", item.Name, rnd)
	}
	return allAbilites
}

func ManualStats(abilites []GenericName) map[string]int {
	allAbilites := make(map[string]int)
	for _, item := range abilites {
		fmt.Printf("Please input what value you want for %s: ", item.Name)
		var value int
		fmt.Scanln(&value)
		allAbilites[item.Name] = value
	}
	return allAbilites
}
