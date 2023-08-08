package migrate

import (
	"migrate-db-primary-key/pkg/fs"

	"github.com/TanmoySG/wunderDB/model"
	log "github.com/sirupsen/logrus"
)

var (
	defaultPKey model.Identifier = "recordId"
)

func Migrate(databaseFilepath string) error {
	loadedDbs, err := fs.LoadDatabases(databaseFilepath)
	if err != nil {
		return err
	}

	updateDatabases(loadedDbs)

	err = fs.UnloadDatabases(databaseFilepath, loadedDbs)
	if err != nil {
		return err
	}

	return nil
}

func updateDatabases(databases map[model.Identifier]*model.Database) {
	for dbId, db := range databases {
		log.Infof("Running migration for Database: [%s]", dbId)
		updateCollectionsPrimaryKey(db)
	}
}

func updateCollectionsPrimaryKey(database *model.Database) {
	for collectionId, collection := range database.Collections {
		log.Infof("Running migration for Collection: [%s]", collectionId)
		if collection.PrimaryKey == nil {
			collection.PrimaryKey = &defaultPKey
			updateRecordsPrimaryKeyValue(collection)
		}
	}
}

func updateRecordsPrimaryKeyValue(collection *model.Collection) {
	for _, record := range collection.Data {
		if record.RecordId == "" {
			record.RecordId = record.Identifier
		}
	}
}
