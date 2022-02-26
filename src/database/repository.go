package database

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	pagination "pkg/src/applibs"
	"pkg/src/logger"
)

func Save(db *gorm.DB, model interface{}) (interface{}, error) {
	if err := db.Create(model).Error; err != nil {
		logger.Errorf("error: %v", err)
		return model, err
	}

	return model, nil
}

//
//func Save(model interface{}) interface{} {
//	query := DB.Create(model)
//	if query.Error != nil {
//		fmt.Println(query.Error)
//	}
//	return query
//}

func _GetOne(model interface{}, field string, value interface{}) interface{} {
	query := DB.Where(fmt.Sprintf("%v = ?", field), value).Find(&model)
	if query.Error != nil {
		fmt.Println(query.Error)
	}
	return query
}

func _GetFirst(model interface{}, field string) interface{} {
	query := DB.First(&model, field)
	if query.Error != nil {
		fmt.Println(query.Error)
	}
	return query
}

func _GetLast(model interface{}, fieldData schema.Field) interface{} {
	query := DB.Last(&model, fieldData)
	if query.Error != nil {
		fmt.Println(query.Error)
	}
	return query
}
func GetAll(model interface{}, pagination *pagination.Pagination) interface{} {
	offset := (pagination.Page - 1) * pagination.Limit
	query := DB.Limit(pagination.Limit).Offset(offset).Order(pagination.OrderBy)
	result := query.Model(&model).Where(model).Find(&model)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return result
}

//
//import (
//	"context"
//)

//
//func Save(ctx context.Context, model interface{}) error {
//	err := _Save(ctx, model)
//	return err
//}
//func Get(ctx context.Context, model interface{}, orderBy string) error {
//	query := _Get(ctx, model, orderBy)
//	return query
//}
//
//func update() {
//
//}
