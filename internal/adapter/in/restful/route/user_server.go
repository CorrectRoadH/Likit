package route

import (
	"net/http"

	"github.com/CorrectRoadH/Likit/codegen"
	"github.com/CorrectRoadH/Likit/utils"
	"github.com/labstack/echo/v4"
)

type userServer struct{}

func NewUserService() codegen.ServerInterface {
	return &userServer{}
}

func (a *userServer) Login(ctx echo.Context) error {
	var loginRequest codegen.LoginRequest

	if err := ctx.Bind(&loginRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, &codegen.BaseResponse{
		Status: utils.Ptr("ok"),
		Msg:    utils.Ptr(loginRequest.UserName + ":" + loginRequest.Password),
	})
}

func (a *userServer) UserInfo(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &codegen.ResponseUserInfo{
		Name:             utils.Ptr("test"),
		Avatar:           utils.Ptr("https://avatars.githubusercontent.com/u/11234?v=4"),
		Email:            utils.Ptr("a778917369@gmail.com"),
		AccountId:        utils.Ptr("11234"),
		RegistrationTime: utils.Ptr("2021-01-01 00:00:00"),
		Permissions:      utils.Ptr("admin"),
	})
}
