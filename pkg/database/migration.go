package database

import (
	examples "gin-boilerplate/examples/ex_models"
	"gin-boilerplate/models"
)

//Add list of model add for migrations
//var migrationModels = []interface{}{&ex_models.Example{}, &model.Example{}, &model.Address{})}
var migrationModels = []interface{}{&models.Example{}, &examples.User{}, &examples.CreditCard{}}
