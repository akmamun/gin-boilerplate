package repository

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
//p, _ := (&helpers.Param{
//	Page:  page,
//	Limit: limit,
//	Path:  ctx.FullPath(),
//	Sort:  "id desc",
//}).Paginate(base.DB, &example)
//ctx.JSON(http.StatusOK, p)

//}
