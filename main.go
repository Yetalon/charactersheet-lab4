package main

import "fmt"

type Character struct {
	characterName string
	class         string
	//abilites      []int
	proficiencies []string
	Weapons       []string
	alignment     string
	race          string
}

func main() {
	class := getclass()
	weapons := getWeapons(class.className)
	charateristics := getCharacteristics()
	fmt.Println("Please type your characters name")
	var characterName string
	fmt.Scanln(&characterName)
	mychar := Character{characterName, class.className, class.proficencychoices, weapons, charateristics.alignment, charateristics.race}
	fmt.Println(mychar)
}
