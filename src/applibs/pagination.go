package applibs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Pagination struct {
	Limit   int    `json:"limit"`
	Page    int    `json:"page"`
	OrderBy string `json:"order_by"`
}

func LimitOffsetOrPagination(ctx *gin.Context) Pagination {
	limit := 10
	page := 1
	orderBy := `id DESC`
	query := ctx.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "orderBy":
			orderBy = queryValue
			break
		}
	}
	fmt.Println(query, limit, orderBy)

	return Pagination{
		Limit:   limit,
		Page:    page,
		OrderBy: orderBy,
	}
}
