package main

import "fmt"

type Character struct {
	characterName     string
	class             string
	abilities         map[string]int
	proficiencies     []string
	equipment         []string
	alignment         string
	race              string
	secondaryLanguage string
}

func main() {
	class, err := getclass()
	if err != nil {
		fmt.Println(err)
		return
	}
	weapons := getEquipment(class.className)
	charateristics := getCharacteristics()
	abilities, err := getAbilitiesScores()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Please type your characters name")
	var characterName string
	fmt.Scanln(&characterName)
	mychar := Character{characterName, class.className, abilities, class.proficencychoices, weapons, charateristics.alignment, charateristics.race, charateristics.langauge}
	fmt.Printf(`
		Character Name:  		%s
		Class:           		%s
		Abilities:       		%+v
		Proficiencies:   		%v
		Equipment:       		%v
		Alignment:       		%s
		Race:            		%s
		Secondary Language:		%v
	`, mychar.characterName, mychar.class, mychar.abilities, mychar.proficiencies, mychar.equipment, mychar.alignment, mychar.race, mychar.secondaryLanguage)
}
