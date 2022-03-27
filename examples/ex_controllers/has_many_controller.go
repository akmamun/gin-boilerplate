package ex_controllers

import (
	examples "gin-boilerplate/examples/ex_models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CreditCardData struct {
	Number string `json:"number"`
}

//GetHasManyRelationUserData fetch user data with preload
func (base *Controller) GetHasManyRelationUserData(ctx *gin.Context) {
	var user []examples.User
	base.DB.Preload("CreditCards").Find(&user)
	ctx.JSON(http.StatusOK, &user)

}

//GetHasManyRelationCreditCardData fetch credit-card data with preload
func (base *Controller) GetHasManyRelationCreditCardData(ctx *gin.Context) {
	var creditCards []examples.CreditCard
	base.DB.Find(&creditCards)
	ctx.JSON(http.StatusOK, &creditCards)

}

// GetUserDetails based on user_id
func (base *Controller) GetUserDetails(ctx *gin.Context) {
	var user []examples.User
	userId, _ := strconv.Atoi(ctx.DefaultQuery("user_id", ""))

	err := base.DB.Preload("CreditCards").First(&user, userId).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"user_id": "Enter valid user"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
