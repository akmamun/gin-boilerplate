package migrations

import (
	"gin-boilerplate/models"
	"gin-boilerplate/pkg/database"
)

// Migrate Add list of model add for migrations
// TODO later separate migration each models
func Migrate() {
	var migrationModels = []interface{}{&models.Example{}}
	err := database.GetDB().AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
