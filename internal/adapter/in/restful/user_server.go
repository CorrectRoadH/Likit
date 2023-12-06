package restful

import (
	"github.com/CorrectRoadH/Likit/internal/port/in"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserServer struct {
	userUseCase in.UserUseCase
}

func NewUserServer(userUseCase in.UserUseCase) *UserServer {
	return &UserServer{
		userUseCase: userUseCase,
	}
}

// translation data
type UserEvent struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// controller implement
func (u *UserServer) Login(c echo.Context) error {
	var event UserEvent

	// bad request return 400
	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// TODO this _ value is judge return success or login error
	_, err := u.userUseCase.Login(c.Request().Context(), event.Username, event.Password)

	// service error return 500
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "login success!")
}

// register http path
func (u *UserServer) register(g *echo.Group) error {

	g.POST("/user/login", u.Login)
	g.POST("/user/add", u.CreateUser)

	return nil
}

func (u *UserServer) CreateUser(c echo.Context) error {
	var event UserEvent
	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := u.userUseCase.CreateUser(c.Request().Context(), event.Username, event.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "create success!")
}
