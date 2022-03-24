package ex_controllers

import (
	examples "gin-boilerplate/examples/ex_models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetHasManyRelationData with Preload
func (base *Controller) GetHasManyRelationData(ctx *gin.Context) {
	var creditCards []examples.CreditCard
	base.DB.Preload("CreditCards").Preload("User").Find(&creditCards)
	ctx.JSON(http.StatusOK, &creditCards)

}
