package route

import (
	"net/http"

	"github.com/CorrectRoadH/Likit/codegen"
	"github.com/CorrectRoadH/Likit/utils"
	"github.com/labstack/echo/v4"
)

type api struct{}

func NewAPIService() codegen.ServerInterface {
	return &api{}
}

func (a *api) Login(ctx echo.Context) error {
	var loginRequest codegen.LoginRequest

	if err := ctx.Bind(&loginRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, &codegen.BaseResponse{
		Status: utils.Ptr("success"),
		Msg:    utils.Ptr(loginRequest.UserName + ":" + loginRequest.Password),
	})
}
