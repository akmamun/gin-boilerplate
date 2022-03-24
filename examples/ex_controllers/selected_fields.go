package ex_controllers

import (
	"gin-boilerplate/examples/ex_models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SelectedFiledFetch fields fetch from defining new struct
type SelectedFiledFetch struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func (base *Controller) GetSelectedFieldData(ctx *gin.Context) {
	var selectData []SelectedFiledFetch
	base.DB.Model(&ex_models.Article{}).Find(&selectData)
	ctx.JSON(http.StatusOK, selectData)

}
