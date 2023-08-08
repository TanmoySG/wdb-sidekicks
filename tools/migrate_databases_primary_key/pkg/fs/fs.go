package fs

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TanmoySG/wunderDB/model"
	"github.com/TanmoySG/wunderDB/pkg/fs"
)

func LoadDatabases(databasesFsPath string) (map[model.Identifier]*model.Database, error) {

	var databases map[model.Identifier]*model.Database

	if fs.CheckFileExists(databasesFsPath) {
		persitedDatabasesBytes, err := os.ReadFile(databasesFsPath)
		if err != nil {
			return nil, fmt.Errorf("error reading database file: %s", err)
		}

		err = json.Unmarshal(persitedDatabasesBytes, &databases)
		if err != nil {
			return nil, fmt.Errorf("error marshaling database file: %s", err)
		}
	}

	return databases, nil
}

func UnloadDatabases(databasesFsPath string, databases map[model.Identifier]*model.Database) error {
	databasesJson, err := json.Marshal(databases)
	if err != nil {
		return err
	}

	if !fs.CheckFileExists(databasesFsPath) {
		os.Create(databasesFsPath)
	}

	err = os.WriteFile(databasesFsPath, databasesJson, 0740)
	if err != nil {
		return err
	}
	return nil
}
