package backend

import (
	"github.com/gin-gonic/gin"
	"mqenergy-go/app/controller/base"
	"mqenergy-go/app/service/backend"
	"mqenergy-go/entities/user"
	"mqenergy-go/pkg/response"
	"net/http"
)

type UserController struct {
	base.Controller
}

var User = UserController{}

// Index 获取列表
func (c UserController) Index(ctx *gin.Context) {
	var requestParams user.IndexRequest
	if err := ctx.Bind(&requestParams); err != nil {
		response.BadRequestException(ctx, "")
		return
	}
	list, err := backend.User.GetList(requestParams)
	if err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	response.ResponseJson(ctx, http.StatusOK, response.Success, "", list)
}
