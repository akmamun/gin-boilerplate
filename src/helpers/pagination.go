package helpers

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"strconv"
)

type Result struct {
	Total    int         `json:"total"`
	Data     interface{} `json:"data"`
	Offset   int         `json:"offset"`
	Limit    int         `json:"limit"`
	Page     int         `json:"page"`
	PrevPage string      `json:"prev_page"`
	NextPage string      `json:"next_page"`
}

func Offset(offset string) int {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = 0
	}
	return offsetInt
}

type Param struct {
	Page  int
	Sort  string
	Limit int
	Path  string
}

func (param *Param) Paginate(db *gorm.DB, any interface{}) (Result, *gorm.DB) {
	var result Result
	var count int64

	offset := (param.Page - 1) * param.Limit
	data := db.Offset(offset).Limit(param.Limit)

	if param.Sort != "" {
		data.Order(param.Sort)
	}

	data.Find(any)
	db.Model(any).Count(&count)

	result.Page = param.Page
	result.NextPage = param.GetPageURL(param.Page + 1)
	result.PrevPage = param.PreviousPage()

	result.Total = int(count)
	result.Data = any

	return result, data
}

func (param *Param) GetPageURL(page int) string {
	return fmt.Sprintf("%s%s?page=%d&limit=%d", GetAppURL(), param.Path, page, param.Limit)
}

func (param *Param) PreviousPage() string {
	pageNumber := 1

	if param.Page > 1 {
		pageNumber = param.Page - 1
	}

	return param.GetPageURL(pageNumber)
}
func GetAppURL() string {
	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	return host + ":" + port
}
