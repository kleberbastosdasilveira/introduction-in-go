package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kleberbastosdasilveira/introduction-in-go/src/controller/model/request"
	"github.com/kleberbastosdasilveira/introduction-in-go/src/controller/model/view/test/configuration/rest_err"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("The are some incorrect filds,error=%s", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
