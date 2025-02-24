package main

import (
	"fmt"
	"strconv"
	"strings"
)

type EquipmentBase struct {
	Index string `json:"index"`
	Name  string `json:"name"`
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
	Of     *EquipmentBase `json:"of,omitempty"`
	Choice *EquipmentMisc `json:"choice,omitempty"`
	Items  []struct {
		Count int           `json:"count"`
		Of    EquipmentBase `json:"of"`
	} `json:"items,omitempty"`
}

type EquipmentFrom struct {
	Options []EquipmentOptions `json:"options"`
}

type StartingEquimentOptions struct {
	Desc string        `json:"desc"`
	From EquipmentFrom `json:"from"`
}

func getWeapons(class string) []string {
	weapons := chooseStartingWeapons(class)
	return weapons
}

func checkChosenWeapon(input string, weapons []string) (string, error) {
	num, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return "", err
	}
	choice := weapons[num-1]
	return choice, nil
}

func getweaponchoice(list []EquipmentOptions) []string {
	listToReturn := []string{}
	for _, value := range list {
		if value.Of != nil && value.Of.Name != "" {
			listToReturn = append(listToReturn, value.Of.Name)
			fmt.Printf("%d) %s\n", len(listToReturn), value.Of.Name)
		}
		if value.Choice != nil && value.Choice.From.EquipmentCat.Name != "" {
			listToReturn = append(listToReturn, value.Choice.From.EquipmentCat.Name)
			fmt.Printf("%d) %s\n", len(listToReturn), value.Choice.From.EquipmentCat.Name)
		}
		if len(value.Items) > 0 {
			for _, item := range value.Items {
				if item.Of.Name != "" {
					listToReturn = append(listToReturn, item.Of.Name)
					fmt.Printf("%d) %s\n", len(listToReturn), item.Of.Name)
				}
			}
		}
	}
	return listToReturn
}

func chooseStartingWeapons(class string) []string {
	info, err := getclassinfo(class)
	if err != nil {
		fmt.Println("Error", err)
		return nil
	}

	EquipmentChoices := info.StartingEquipmentChoices
	firstWeaponChoice := EquipmentChoices[0].From.Options
	secondWeaponChoice := EquipmentChoices[1].From.Options

	fmt.Println("Please choose a primary equipment")
	firstChoice := getweaponchoice(firstWeaponChoice)

	var userfirstchoice string
	fmt.Scanln(&userfirstchoice)

	primaryWeapon, err := checkChosenWeapon(userfirstchoice, firstChoice)
	if err != nil {
		fmt.Println("Error:", err)
		return []string{}
	}

	fmt.Println("Please choose a seconday equipment")
	secondChoice := getweaponchoice(secondWeaponChoice)

	var usersecondchoice string
	fmt.Scanln(&usersecondchoice)

	secondaryWeapon, err := checkChosenWeapon(usersecondchoice, secondChoice)
	if err != nil {
		fmt.Println("Error: ", err)
		return []string{}
	}

	return []string{primaryWeapon, secondaryWeapon}
}
