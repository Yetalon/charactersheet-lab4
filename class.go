package main

import (
	"fmt"
	"strconv"
	"strings"
)

type userclass struct {
	className         string
	proficencychoices []string
	weaponChoices     []string
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

type EquipmentBase struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type EquipmentMisc struct {
	From struct {
		EquipmentCat struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"equipment_category"`
	} `json:"from"`
}

type EquipmentOptions struct {
	Of     EquipmentBase `json:"of"`
	Choice EquipmentMisc `json:"choice"`
}

type EquipmentFrom struct {
	Options []EquipmentOptions `json:"options"`
}

type StartingEquimentOptions struct {
	Desc string        `json:"desc"`
	From EquipmentFrom `json:"from"`
}

func getclass() userclass {
	class := chooseclass()
	profs := chooseProfs(strings.TrimSpace(class))
	weps := chooseStartingWeapons(class)
	usersClassInfo := userclass{class, profs, weps}
	return usersClassInfo
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

func checkChoiceInput(input string, allprofs []string) (string, string, error) {
	parts := strings.Split(input, ",")
	num1, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return "", "", err
	}
	var num2 int
	if len(parts) != 1 {
		num2, err = strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return "", "", err
		}
	}
	choice1 := allprofs[num1-1]
	var choice2 string
	if len(parts) != 1 {
		choice2 = allprofs[num2-1]
	}
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
	choice1, choice2, err := checkChoiceInput(userchoices, allprofs)
	if err != nil {
		return strProfCho
	}
	strProfCho = append(strProfCho, choice1, choice2)
	return strProfCho
}

func chooseStartingWeapons(class string) []string {
	info, err := getclassinfo(class)
	if err != nil {
		fmt.Println("Error", err)
		return nil
	}
	fmt.Println("Please choose a primary equipment")
	EquipmentChoices := info.StartingEquipmentChoices
	firstWeaponChoice := EquipmentChoices[0].From.Options
	secondWeaponChoice := EquipmentChoices[1].From.Options

	first := []string{}

	for i, value := range firstWeaponChoice {
		if value.Of.Name != "" {
			first = append(first, value.Of.Name)
			fmt.Printf("%d) %s\n", i+1, value.Of.Name)
		}
		if value.Choice.From.EquipmentCat.Name != "" {
			first = append(first, value.Choice.From.EquipmentCat.Name)
			fmt.Printf("%d) %s\n", i+1, value.Choice.From.EquipmentCat.Name)
		}
	}
	var userfirstchoice string
	fmt.Scanln(&userfirstchoice)

	primaryWeapon, _, err := checkChoiceInput(userfirstchoice, first)

	second := []string{}

	fmt.Println("Please choose a seconday equipment")
	for i, value := range secondWeaponChoice {
		if value.Of.Name != "" {
			second = append(second, value.Of.Name)
			fmt.Printf("%d) %s\n", i+1, value.Of.Name)
		}
		if value.Choice.From.EquipmentCat.Name != "" {
			second = append(second, value.Choice.From.EquipmentCat.Name)
			fmt.Printf("%d) %s\n", i+1, value.Choice.From.EquipmentCat.Name)
		}
	}
	var usersecondchoice string
	fmt.Scanln(&usersecondchoice)

	secondaryWeapon, _, err := checkChoiceInput(usersecondchoice, second)

	return []string{primaryWeapon, secondaryWeapon}
}
