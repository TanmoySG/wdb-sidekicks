package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "usage" {
			fmt.Println("./roles_hidden_field_update {data-in-json} 'value-to-set'")
		}
		os.Exit(0)
	} else if len(os.Args) < 3 {
		log.Fatal("not enough inputs")
		os.Exit(1)
	}

	var filepath, rolesJsonString, rolesHiddenFieldSetValue string

	flag := os.Args[1]
	if flag == "-f" {
		filepath = os.Args[2]
		fileContentBytes, err := os.ReadFile(filepath)
		if err != nil {
			log.Fatalf("failed to parse bool: %s", err)
		}
		rolesJsonString = string(fileContentBytes)
		rolesHiddenFieldSetValue = os.Args[3]

	} else {
		rolesJsonString = os.Args[1]
		rolesHiddenFieldSetValue = os.Args[2]
	}

	rolesHiddenFieldSetValueBool, err := strconv.ParseBool(rolesHiddenFieldSetValue)
	if err != nil {
		log.Fatalf("failed to parse bool: %s", err)
	}

	parsedJsonMap := parseJson(rolesJsonString)
	mutatedJsonMap := mutate(parsedJsonMap, rolesHiddenFieldSetValueBool)

	if flag == "-f" {
		saveToFile(mutatedJsonMap, filepath)
	} else {
		mutatedJsonString, err := json.Marshal(mutatedJsonMap)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(mutatedJsonString))
	}

}

func mutate(roles map[string]interface{}, setArg bool) map[string]interface{} {
	for roleId, role := range roles {
		if _, ok := role.(map[string]interface{})["hidden"]; !ok {
			roles[roleId].(map[string]interface{})["hidden"] = setArg
		}
	}

	return roles
}

func parseJson(jsonString string) map[string]interface{} {
	var parsedJsonMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &parsedJsonMap)
	if err != nil {
		log.Fatal(err)
	}

	return parsedJsonMap
}

func saveToFile(jsonMap map[string]interface{}, jsonFilepath string) {
	jsonBytes, err := json.Marshal(jsonMap)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(jsonFilepath, jsonBytes, 0740)
	if err != nil {
		log.Fatal(err)
	}
}
