package route

import (
	"net/http"

	"github.com/CorrectRoadH/Likit/codegen"
	"github.com/CorrectRoadH/Likit/internal/application/domain"
	"github.com/CorrectRoadH/Likit/internal/port/in"
	"github.com/CorrectRoadH/Likit/utils"
	"github.com/labstack/echo/v4"

	"github.com/samber/lo"
)

type adminApiService struct {
	adminUseCase    in.AdminUseCase
	databaseUseCase in.DatabaseUseCase
}

func convertBusiness(business codegen.Business) domain.Business {
	return domain.Business{
		Id:    *business.Id,
		Title: *business.Title,
		Type:  *business.Type,
	}
}

func convertDatabaseConnectConfig(databaseConnectConfig codegen.DatabaseConnectConfig) domain.DatabaseConnectConfig {
	return domain.DatabaseConnectConfig{
		DatabaseType: domain.DatabaseType(*databaseConnectConfig.DatabaseType),
		Host:         *databaseConnectConfig.Host,
		Port:         *databaseConnectConfig.Port,
		Username:     *databaseConnectConfig.Username,
		Password:     *databaseConnectConfig.Password,
		Database:     *databaseConnectConfig.Database,
	}
}

func NewAdminApiService(
	adminUseCase in.AdminUseCase,
	databaseUseCase in.DatabaseUseCase,
) codegen.ServerInterface {
	return &adminApiService{
		adminUseCase:    adminUseCase,
		databaseUseCase: databaseUseCase,
	}
}

func (a *adminApiService) Login(ctx echo.Context) error {
	var loginRequest codegen.LoginRequest

	if err := ctx.Bind(&loginRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, &codegen.BaseResponse{
		Status: utils.Ptr("ok"),
		Msg:    utils.Ptr(loginRequest.UserName + ":" + loginRequest.Password),
	})
}

func (a *adminApiService) UserInfo(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &codegen.ResponseUserInfo{
		Name:             utils.Ptr("test"),
		Avatar:           utils.Ptr("https://avatars.githubusercontent.com/u/11234?v=4"),
		Email:            utils.Ptr("a778917369@gmail.com"),
		AccountId:        utils.Ptr("11234"),
		RegistrationTime: utils.Ptr("2021-01-01 00:00:00"),
		Permissions:      utils.Ptr("admin"),
	})
}

func (u *adminApiService) GetBusinesses(ctx echo.Context) error {
	businesses, err := u.adminUseCase.Businesses(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{
			Status: utils.Ptr("error"),
			Msg:    utils.Ptr(err.Error()),
		})
	}

	return ctx.JSON(http.StatusOK, lo.Map(businesses, func(business domain.Business, _ int) codegen.Business {
		return codegen.Business{
			Id:    utils.Ptr(business.Id),
			Title: utils.Ptr(business.Title),
			Type:  utils.Ptr(business.Type),
		}
	}))
}

func (a *adminApiService) CreateBusiness(ctx echo.Context) error {
	var createBusinessRequest codegen.Business
	if err := ctx.Bind(&createBusinessRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	err := a.adminUseCase.CreateBusiness(ctx.Request().Context(), convertBusiness(createBusinessRequest))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{
			Status: utils.Ptr("error"),
			Msg:    utils.Ptr(err.Error()),
		})
	}

	return ctx.JSON(http.StatusOK, &codegen.BaseResponse{
		Status: utils.Ptr("ok"),
	})
}

// func (a *adminApiService) DeleteBusiness(ctx echo.Context) error {
// 	var deleteBusinessRequest codegen.DeleteBusinessRequest
// 	if err := ctx.Bind(&deleteBusinessRequest); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, err)
// 	}

// 	err := a.adminServer.DeleteBusiness(ctx.Request().Context(), *deleteBusinessRequest.BusinessId)
// 	if err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, err)
// 	}

// 	return ctx.JSON(http.StatusOK, &codegen.BaseResponse{
// 		Status: utils.Ptr("ok"),
// 	})
// }

func (a *adminApiService) UpdateBusiness(ctx echo.Context) error {
	var createBusinessRequest codegen.Business
	if err := ctx.Bind(&createBusinessRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	err := a.adminUseCase.UpdateBusiness(ctx.Request().Context(), convertBusiness(createBusinessRequest))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{
			Status: utils.Ptr("error"),
			Msg:    utils.Ptr(err.Error()),
		})
	}

	return ctx.JSON(http.StatusOK, &codegen.BaseResponse{
		Status: utils.Ptr("ok"),
	})
}

func (a *adminApiService) CreatePostgresDatabase(ctx echo.Context) error {
	panic("TODO: Implement")
}

func (a *adminApiService) CreateRedisDatabase(ctx echo.Context) error {
	panic("TODO: Implement")
}

func (a *adminApiService) GetPostgresDatabaseList(ctx echo.Context) error {
	panic("TODO: Implement")
}

func (a *adminApiService) GetRedisDatabaseList(ctx echo.Context) error {
	panic("TODO: Implement")
}

func (a *adminApiService) UpdatePostgresDatabase(ctx echo.Context) error {
	panic("TODO: Implement")
}

func (a *adminApiService) UpdateRedisDatabase(ctx echo.Context) error {
	panic("TODO: Implement")
}

func (a *adminApiService) TestDatabaseConnection(ctx echo.Context) error {
	var databaseConnectConfig codegen.DatabaseConnectConfig
	if err := ctx.Bind(&databaseConnectConfig); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	err := a.databaseUseCase.TestConnection(convertDatabaseConnectConfig(databaseConnectConfig))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{
			Status: utils.Ptr("error"),
			Msg:    utils.Ptr(err.Error()),
		})
	}

	return ctx.JSON(http.StatusOK, &codegen.BaseResponse{
		Status: utils.Ptr("ok"),
	})
}
