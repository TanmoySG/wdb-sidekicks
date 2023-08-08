package system

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TanmoySG/wunderDB/pkg/utils/system"
		"github.com/TanmoySG/wunderDB/pkg/fs"

)

const (
	WDB_ROOT_PATH_FORMAT = "%s/wdb"

	WDB_PERSISTANT_FILE_PATH = "%s/databases/databases_persisted.json"

	WDB_CONFIG_FILE_PATH_FORMAT = "%s/configs/conf.json"
)

func GetWdbRoot() (*string, error) {
	hostOS, err := system.GetHostOS()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	homeDir := system.GetUserHome(hostOS)
	wdbRootDirectory := fmt.Sprintf(WDB_ROOT_PATH_FORMAT, homeDir)

	return &wdbRootDirectory, nil
}

func CheckIfRootExists(wdbRootpath string) bool{
	return fs.CheckFileExists(wdbRootpath)
}

func GetDatabasesPersistedPath(wdbRootPath string) (*string, error) {
	var configMap map[string]string

	configpath := fmt.Sprintf(WDB_CONFIG_FILE_PATH_FORMAT, wdbRootPath)

	persitedDatabasesBytes, err := os.ReadFile(configpath)
	if err != nil {
		return nil, fmt.Errorf("error reading database file: %s", err)
	}

	err = json.Unmarshal(persitedDatabasesBytes, &configMap)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling file: %s", err)
	}

	wfsPath, ok := configMap["PERSISTANT_STORAGE_PATH"]
	if !ok {
		return nil, fmt.Errorf("persistant file path [%s] not found", configpath)
	}

	databasePersistedPath := fmt.Sprintf(WDB_PERSISTANT_FILE_PATH, wfsPath)
	return &databasePersistedPath, nil
}
