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

	rolesJsonString := os.Args[1]
	rolesHiddenFieldSetValue := os.Args[2]
	rolesHiddenFieldSetValueBool, err := strconv.ParseBool(rolesHiddenFieldSetValue)
	if err != nil {
		log.Fatalf("failed to parse bool: %s", err)
	}

	var parsedJsonMap map[string]interface{}
	err = json.Unmarshal([]byte(rolesJsonString), &parsedJsonMap)
	if err != nil {
		log.Fatal(err)
	}

	mutatedJsonMap := mutate(parsedJsonMap, rolesHiddenFieldSetValueBool)
	mutatedJsonString, err := json.Marshal(mutatedJsonMap)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(mutatedJsonString))
}

func mutate(roles map[string]interface{}, setArg bool) map[string]interface{} {
	for roleId, role := range roles {
		if _, ok := role.(map[string]interface{})["hidden"]; !ok {
			roles[roleId].(map[string]interface{})["hidden"] = setArg
		}
	}

	return roles
}
