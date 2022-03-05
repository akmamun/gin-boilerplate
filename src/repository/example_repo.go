package repository

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"pkg/src/helpers"
	"pkg/src/logger"
	"pkg/src/models"
)


func Search(search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			db = db.Where("data LIKE ?", "%"+search+"%")
			//db = db.Or("description LIKE ?", "%"+search+"%")
		}
		return db
	}
}
func GetExamples(ctx *gin.Context, db *gorm.DB, args *helpers.Args) ([]models.Example, int64, int64, error) {
	var example []models.Example
	var filteredData, totalData int64

	table := "posts"
	query := db.Select(table + ".*")
	query = query.Offset(Offset(args.Offset))
	query = query.Limit(Limit(args.Limit))
	query = query.Order(SortOrder(table, args.Sort, args.Order))
	query = query.Scopes(Search(args.Search))

	if err := query.Preload("Tags").Find(&example).Error; err != nil {
		logger.Errorf("GetPosts error: %v", err)
		return example, filteredData, totalData, err
	}

	// // Count filtered table
	// // We are resetting offset to 0 to return total number.
	// // This is a fix for Gorm offset issue
	query = query.Offset(0)
	query.Table(table).Count(&filteredData)

	// // Count total table
	db.Table(table).Count(&totalData)

	return posts, filteredData, totalData, nil
}
	//var example []models.Example
	//err := db.Find(&example).Error
	//if err != nil {
	//	logger.Errorf("error: %v", err)
	//}
	//
	//p, _ := (&helpers.Args{
	//	Page:  page,
	//	Limit: limit,
	//	Path:  ctx.FullPath(),
	//	Sort:  "id asc",
	//}).Paginate(base.DB, &example)
	//ctx.JSON(http.StatusOK, p)

}

//func GetPosts(c *gin.Context, db *gorm.DB) ([]models.Example, int64, int64, error) {
//func (ctx *gin.Context, db *gorm.DB) GetExamples([] model.Example) {
//var example []models.Example

//err := base.DB.Find(&example).Error
////err := base.rep.Find(&example)
//if err != nil {
//	logger.Errorf("error: %v", err)
//}
//page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
//limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "0"))
//p, _ := (&helpers.Args{
//	Page:  page,
//	Limit: limit,
//	Path:  ctx.FullPath(),
//	Sort:  "id desc",
//}).Paginate(base.DB, &example)
//ctx.JSON(http.StatusOK, p)

//}
