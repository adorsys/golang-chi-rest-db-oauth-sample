package migration

import (
	"github.com/mattes/migrate/migrate"
	"log"
)

func Do(dbUrl string) (uint64, []error) {
	if versionBefore, errors := getVersion(dbUrl); errors != nil {
		return 0, errors
	} else {
		if allErrors, ok := migrate.UpSync(dbUrl, "./data/migration"); !ok {
			log.Printf("Could not apply migration. Reason: %s", allErrors)
			return 0, allErrors
		} else {
			if versionNow, errors := getVersion(dbUrl); errors != nil {
				return 0, errors
			} else {
				if versionBefore == versionNow {
					log.Printf("DB is already migrated (version=%d). Nothing to do.", versionNow)
				} else {
					log.Printf("Succesfully migrated DB from version=%d to version=%d", versionBefore, versionNow)
				}
				return versionNow, nil
			}
		}
	}
}

func getVersion(dbUrl string) (uint64, []error) {
	if versionBefore, err := migrate.Version(dbUrl, "./data/migration"); err != nil {
		log.Print("Can't query db for migration information", err)
		return 0, []error{err}
	} else {
		return versionBefore, nil
	}
}
