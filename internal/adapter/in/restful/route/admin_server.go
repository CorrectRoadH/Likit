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
	userUseCase     in.UserUseCase
	voteUseCase     in.VoteAdminUseCase
}

func convertBusiness(business codegen.Business) domain.Business {
	return domain.Business{
		Id:    *business.Id,
		Title: *business.Title,
		Type:  *business.Type,
		Config: domain.Config{
			DataSourceConfig: lo.Map(*business.Config.DataSourceConfig, func(dataSourceConfig codegen.DatabaseConnectConfig, _ int) domain.DatabaseConnectConfig {
				return convertDatabaseConnectConfig(dataSourceConfig)
			}),
		},
	}
}

func convertDatabaseConnectConfig(databaseConnectConfig codegen.DatabaseConnectConfig) domain.DatabaseConnectConfig {
	return domain.DatabaseConnectConfig{
		Id:           *databaseConnectConfig.Id,
		Title:        *databaseConnectConfig.Title,
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
	userUseCase in.UserUseCase,
	voteUseCase in.VoteAdminUseCase,
) codegen.ServerInterface {
	return &adminApiService{
		adminUseCase:    adminUseCase,
		databaseUseCase: databaseUseCase,
		userUseCase:     userUseCase,
		voteUseCase:     voteUseCase,
	}
}

func (a *adminApiService) Login(ctx echo.Context) error {
	var loginRequest codegen.LoginRequest

	if err := ctx.Bind(&loginRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	verified, err := a.userUseCase.Login(ctx.Request().Context(), loginRequest.UserName, loginRequest.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{
			Status: utils.Ptr("error"),
			Msg:    utils.Ptr(err.Error()),
		})
	}

	if !verified {
		return ctx.JSON(http.StatusUnauthorized, codegen.ResponseInternalServerError{
			Status: utils.Ptr("error"),
			Msg:    utils.Ptr("username or password is wrong"),
		})
	}

	return ctx.JSON(http.StatusOK, &codegen.BaseResponse{
		Status: utils.Ptr("ok"),
	})
}

func (a *adminApiService) UserInfo(ctx echo.Context) error {

	// TODO wait to implement
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

	return ctx.JSON(http.StatusOK, codegen.ResponseBusinessList{
		Businesses: utils.Ptr(lo.Map(businesses, func(business domain.Business, _ int) codegen.Business {
			return codegen.Business{
				Id:    utils.Ptr(business.Id),
				Title: utils.Ptr(business.Title),
				Type:  utils.Ptr(business.Type),
			}
		})),
	})
}

func (a *adminApiService) CreateBusiness(ctx echo.Context) error {
	var createBusinessRequest codegen.Business
	if err := ctx.Bind(&createBusinessRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	err := a.adminUseCase.CreateBusiness(ctx.Request().Context(), convertBusiness(createBusinessRequest))
	// TODO handle multi error
	if err == nil {
		err = a.voteUseCase.CreateBusiness(ctx.Request().Context(), convertBusiness(createBusinessRequest))
	}
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{
			Status: utils.Ptr("error"),
			Msg:    utils.Ptr(err.Error()),
		})
	}

	return ctx.JSON(http.StatusOK, &codegen.ResponseOK{
		Status: utils.Ptr("ok"),
	})
}

func (a *adminApiService) DeleteBusiness(ctx echo.Context, params codegen.DeleteBusinessParams) error {
	err := a.adminUseCase.DeleteBusiness(ctx.Request().Context(), params.Id)

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

	return ctx.JSON(http.StatusOK, &codegen.ResponseOK{
		Status: utils.Ptr("ok"),
	})
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

func (a *adminApiService) GetDatabaseConfigureList(ctx echo.Context) error {
	connections, err := a.databaseUseCase.DatabaseConfigureList()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{
			Status: utils.Ptr("error"),
			Msg:    utils.Ptr(err.Error()),
		})
	}

	array := lo.Map(connections, func(connection domain.DatabaseConnectConfig, _ int) codegen.DatabaseConnectConfig {
		return codegen.DatabaseConnectConfig{
			Id:           utils.Ptr(connection.Id),
			Title:        utils.Ptr(connection.Title),
			DatabaseType: utils.Ptr(string(connection.DatabaseType)),
			Host:         utils.Ptr(connection.Host),
			Port:         utils.Ptr(connection.Port),
			Username:     utils.Ptr(connection.Username),
			Password:     utils.Ptr(connection.Password),
			Database:     utils.Ptr(connection.Database),
		}
	})

	return ctx.JSON(http.StatusOK, codegen.ResponseDatabaseList{
		DataSourceConfig: &array,
	})
}

func (a *adminApiService) CreateDatabase(ctx echo.Context) error {
	err := a.databaseUseCase.CreateDatabase(ctx.Request().Context(), domain.DatabaseConnectConfig{})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{
			Status: utils.Ptr("error"),
			Msg:    utils.Ptr(err.Error()),
		})
	}
	return ctx.JSON(http.StatusOK, &codegen.ResponseOK{
		Status: utils.Ptr("ok"),
	})
}

func (a *adminApiService) DeleteDatabase(ctx echo.Context, params codegen.DeleteDatabaseParams) error {
	err := a.databaseUseCase.DeleteDatabase(ctx.Request().Context(), domain.DatabaseConnectConfig{
		Id: params.Id,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, codegen.ResponseInternalServerError{
			Status: utils.Ptr("error"),
			Msg:    utils.Ptr(err.Error()),
		})
	}
	return ctx.JSON(http.StatusOK, &codegen.ResponseOK{
		Status: utils.Ptr("ok"),
	})
}

func (a *adminApiService) UpdateDatabase(ctx echo.Context) error {
	panic("TODO: Implement")
}
