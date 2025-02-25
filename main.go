package main

import "fmt"

type Character struct {
	characterName string
	class         string
	abilites      map[string]int
	proficiencies []string
	Weapons       []string
	alignment     string
	race          string
}

func main() {
	class, err := getclass()
	if err != nil {
		fmt.Println(err)
		return
	}
	weapons := getEquipment(class.className)
	charateristics := getCharacteristics()
	abilites, err := getAbilitiesScores()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Please type your characters name")
	var characterName string
	fmt.Scanln(&characterName)
	mychar := Character{characterName, class.className, abilites, class.proficencychoices, weapons, charateristics.alignment, charateristics.race}
	fmt.Println(mychar)
}
