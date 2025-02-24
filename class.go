package main

import (
	"fmt"
	"strconv"
	"strings"
)

type userclass struct {
	className         string
	proficencychoices []string
}

type ProficiencyOption struct {
	Item struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"item"`
}

type ProficiencyFrom struct {
	Options []ProficiencyOption `json:"options"`
}

type ProficiencyChoice struct {
	From ProficiencyFrom `json:"from"`
}

func getclass() (userclass, error) {
	class := chooseclass()
	if class == "" {
		return userclass{}, fmt.Errorf("invalid input ending program")
	}
	profs := chooseProfs(strings.TrimSpace(class))
	usersClassInfo := userclass{class, profs}
	return usersClassInfo, nil
}

func checkClassinput(choice string) (string, error) {
	classMap := map[string]string{
		"1":  "barbarian",
		"2":  "bard",
		"3":  "cleric",
		"4":  "druid",
		"5":  "fighter",
		"6":  "monk",
		"7":  "paladin",
		"8":  "ranger",
		"9":  "rogue",
		"10": "sorcerer",
		"11": "warlock",
		"12": "wizard",
	}
	choice, exists := classMap[choice]
	if !exists {
		return choice, fmt.Errorf("invalid input")
	}
	return choice, nil
}

func chooseclass() string {
	list, err := getclasseslist()
	if err != nil {
		println("error", err)
	}

	fmt.Println("Please choose a class:")
	index := 1
	for _, value := range list {
		fmt.Printf("%d) %s\n", index, value)
		index++
	}

	var class string
	fmt.Scanln(&class)

	userclass, err := checkClassinput(class)
	if err != nil {
		return ""
	}
	return userclass
}

func checkProfsInput(input string, allprofs []string) (string, string, error) {
	parts := strings.Split(input, ",")
	num1, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return "", "", err
	}
	num2, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return "", "", err
	}
	choice1 := allprofs[num1-1]
	choice2 := allprofs[num2-1]
	return choice1, choice2, nil
}

func chooseProfs(class string) []string {
	info, err := getclassinfo(class)
	if err != nil {
		println("Error", err)
		return nil
	}

	fmt.Println("Please choose 2 proficiencies in this format {n,n}")
	proficencyChoicies := info.ProficiencyChoices[0]
	options := proficencyChoicies.From.Options
	allprofs := []string{}
	for i, value := range options {
		allprofs = append(allprofs, strings.TrimPrefix(value.Item.Name, "Skill: "))
		fmt.Printf("%d) %s\n", i+1, strings.TrimPrefix(value.Item.Name, "Skill: "))
	}

	var userchoices string
	fmt.Scanln(&userchoices)

	strProfCho := []string{}
	choice1, choice2, err := checkProfsInput(userchoices, allprofs)
	if err != nil {
		return strProfCho
	}
	strProfCho = append(strProfCho, choice1, choice2)
	return strProfCho
}
