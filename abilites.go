package main

import "fmt"

func getAbilitiesScores() {

}

func chooseAbilities() (map[string]int, error) {
	resp, err := getAbilities()
	if err != nil {
		fmt.Println("Error:", err)
		return []string{}, err
	}
	abilites := resp.GenericBody
	abilitesList := 
	for i, item := range abilites {
		abilitesList = append(abilitesList, item.Name)
		fmt.Printf("%d) %s\n", i+1, item.Name)
	}

}
