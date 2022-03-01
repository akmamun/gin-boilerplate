package applibs

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
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

type Param struct {
	Page  int
	Sort  string
	Limit int
	Path  string
}

func (p *Param) Paginate(db *gorm.DB, any interface{}) (Result, *gorm.DB) {
	var r Result
	var count int64

	offset := (p.Page - 1) * p.Limit
	//lastIndex := offset * param.Page
	data := db.Offset(offset).Limit(p.Limit)

	if p.Sort != "" {
		data.Order(p.Sort)
	}

	data.Find(any)
	db.Model(any).Count(&count)

	r.Page = p.Page

	r.NextPage = p.GetPageURL(p.Page + 1)
	r.PrevPage = p.PreviousPage()

	r.Total = int(count)
	r.Data = any

	return r, data
}

func (p *Param) GetPageURL(page int) string {
	return fmt.Sprintf("%s%s?page=%d&limit=%d", GetAppURL(), p.Path, page, p.Limit)
}

func (p *Param) PreviousPage() string {
	pageNumber := 1

	if p.Page > 1 {
		pageNumber = p.Page - 1
	}

	return p.GetPageURL(pageNumber)
}
func GetAppURL() string {
	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	return host + ":" + port
}
