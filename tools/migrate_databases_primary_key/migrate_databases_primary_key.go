package main

import (
	"migrate-db-primary-key/pkg/migrate"
	"migrate-db-primary-key/pkg/system"

	log "github.com/sirupsen/logrus"
)

func main() {
	wdb_root, err := system.GetWdbRoot()
	if err != nil {
		log.Error(err)
	}

	exists := system.CheckIfRootExists(*wdb_root)
	if !exists {
		log.Fatal("Wdb Root doesn't exist")
	}

	wdb_databases_path, err := system.GetDatabasesPersistedPath(*wdb_root)
	if err != nil {
		log.Fatal(err)
	}

	err = migrate.Migrate(*wdb_databases_path)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("data updated")
}
