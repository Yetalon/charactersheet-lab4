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

func getEquipment(class string) []string {
	weapons := chooseStartingEquipment(class)
	return weapons
}

func chooseMiscWeapon(endpoint string, weapontype string) (string, error) {
	info, err := getGenericResponse(endpoint)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	miscWeapons := info.Equipment
	allMiscWeapons := []string{}
	fmt.Printf("Please choose a %s\n", weapontype)
	for _, item := range miscWeapons {
		allMiscWeapons = append(allMiscWeapons, item.Name)
		fmt.Printf("%d) %s\n", len(allMiscWeapons), item.Name)
	}
	var miscWeapon string
	fmt.Scanln(&miscWeapon)
	chosen, err := checkChosenEquipment(miscWeapon, allMiscWeapons)
	if err != nil {
		return "", err
	}
	return chosen, nil
}

func checkChosenEquipment(input string, weapons []string) (string, error) {
	num, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return "", err
	}
	choice := weapons[num-1]
	return choice, nil
}

func getEquipmentChoice(list []EquipmentOptions) []string {
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
			var groupedItems []string
			for _, item := range value.Items {
				if item.Of.Name != "" {
					groupedItems = append(groupedItems, item.Of.Name)
				}
			}
			if len(groupedItems) > 0 {
				groupedChoice := strings.Join(groupedItems, " and ")
				listToReturn = append(listToReturn, groupedChoice)
				fmt.Printf("%d) %s\n", len(listToReturn), groupedChoice)
			}
		}
	}
	return listToReturn
}

func chooseStartingEquipment(class string) []string {
	info, err := getclassinfo(class)
	if err != nil {
		fmt.Println("Error", err)
		return nil
	}

	EquipmentChoices := info.StartingEquipmentChoices
	equipment := []string{}

	for i := range EquipmentChoices {
		fmt.Printf("Please choose your %d equipment\n", i+1)
		list := EquipmentChoices[i].From.Options
		choices := getEquipmentChoice(list)

		var userChoice string
		fmt.Scanln(&userChoice)

		weapon, err := checkChosenEquipment(userChoice, choices)
		if err != nil {
			fmt.Println("Error:", err)
			return []string{}
		}
		switch weapon {
		case "Martial Melee Weapons":
			weapon, err = chooseMiscWeapon("equipment-categories/martial-melee-weapons", weapon)
			if err != nil {
				weapon = "Martial Melee Weapons"
			}
		case "Martial Weapons":
			weapon, err = chooseMiscWeapon("equipment-categories/martial-weapons", weapon)
			if err != nil {
				weapon = "Martial Weapons"
			}
		case "Simple Weapons":
			weapon, err = chooseMiscWeapon("equipment-categories/simple-weapons", weapon)
			if err != nil {
				weapon = "Simple Weapons"
			}
		}
		equipment = append(equipment, weapon)
	}

	return equipment
}
