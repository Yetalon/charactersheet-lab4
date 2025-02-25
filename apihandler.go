package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ClassInfo struct {
	ProficiencyChoices       []ProficiencyChoice       `json:"proficiency_choices"`
	StartingEquipmentChoices []StartingEquimentOptions `json:"starting_equipment_options"`
}

type GenericApiResults struct {
	GenericBody []GenericName `json:"results"`
	Equipment   []GenericName `json:"equipment"`
}

type GenericName struct {
	Name string `json:"name"`
}

func getBody(endpoint string) (*http.Response, error) {
	const url = "https://www.dnd5eapi.co/api/2014"
	urlToCall := fmt.Sprintf("%s/%s", url, endpoint)
	resp, err := http.Get(urlToCall)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return resp, nil
}

func getclasseslist() ([]string, error) {
	resp, err := getBody("classes")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	var response map[string]interface{}
	json.Unmarshal(body, &response)

	classList := []string{}
	for _, value := range response["results"].([]interface{}) {
		value := value.(map[string]interface{})
		class_name := value["index"]
		classList = append(classList, class_name.(string))
	}
	return classList, nil
}

func getclassinfo(chosenclass string) (ClassInfo, error) {
	endpoint := fmt.Sprintf("classes/%s", chosenclass)
	resp, err := getBody(endpoint)
	if err != nil {
		fmt.Println("Error:", err)
		return ClassInfo{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return ClassInfo{}, err
	}
	var response ClassInfo
	json.Unmarshal(body, &response)
	return response, nil
}

func getGenericResponse(endpoint string) (GenericApiResults, error) {
	resp, err := getBody(endpoint)
	if err != nil {
		fmt.Println("Error:", err)
		return GenericApiResults{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return GenericApiResults{}, err
	}
	var response GenericApiResults
	json.Unmarshal(body, &response)
	return response, nil
}
