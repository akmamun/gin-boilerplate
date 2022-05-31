package pagination

import (
	"gin-boilerplate/pkg/database"
	"gorm.io/gorm"
	"math"
)

type Param struct {
	Page    int64
	Limit   int64
	OrderBy string
	Search  string
}

type Result struct {
	TotalRecord int64       `json:"total_record"`
	TotalPage   int64       `json:"total_page"`
	Offset      int64       `json:"offset"`
	Limit       int64       `json:"limit"`
	Page        int64       `json:"page"`
	PrevPage    int64       `json:"prev_page"`
	NextPage    int64       `json:"next_page"`
	Data        interface{} `json:"data"`
}

func Paginate(param *Param, resultData interface{}) *Result {
	db := database.GetDB()

	if param.Page < 1 {
		param.Page = 1
	}
	if param.Limit == 0 {
		param.Limit = 10
	}

	done := make(chan bool, 1)
	var result Result
	var count, offset int64

	go countResults(db, resultData, done, &count)

	if param.Page == 1 {
		offset = 0
	} else {
		offset = (param.Page - 1) * param.Limit
	}
	db.Offset(int(offset)).
		Limit(int(param.Limit)).
		Order(param.OrderBy).
		Find(resultData)

	<-done

	result.TotalRecord = count
	result.Data = resultData
	result.Page = param.Page

	result.Offset = offset
	result.Limit = param.Limit
	result.TotalPage = int64(math.Ceil(float64(count) / float64(param.Limit)))

	if param.Page > 1 {
		result.PrevPage = param.Page - 1
	} else {
		result.PrevPage = param.Page
	}

	if param.Page == result.TotalPage {
		result.NextPage = param.Page
	} else {
		result.NextPage = param.Page + 1
	}
	return &result
}

// count through separate channel
func countResults(db *gorm.DB, anyType interface{}, done chan bool, count *int64) {
	db.Model(anyType).Count(count)
	done <- true
}
