package restful

import (
	"fmt"
	"net/http"
	"os"

	"github.com/CorrectRoadH/Likit/internal/port/in"
	"github.com/labstack/echo/v4"
)

type VoteServer struct {
	voteUseCase in.VoteUseCase
	e           *echo.Echo
}

func NewVoteServer(voteUseCase in.VoteUseCase) *VoteServer {
	e := echo.New()

	s := &VoteServer{
		voteUseCase: voteUseCase,
		e:           e,
	}

	s.register(e.Group("/api/v1"))

	return s
}

type VoteUsersEvent struct {
	List []string `json:"userList"`
}

type VoteEvent struct {
	BusinessId string `json:"businessId"`
	MessageId  string `json:"messageId"`
	UserId     string `json:"userId"`
}

func (v *VoteServer) Vote(c echo.Context) error {
	var event VoteEvent
	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := v.voteUseCase.Vote(c.Request().Context(), event.BusinessId, event.MessageId, event.UserId)
	if err != nil {
		// TODO to differentiate error type, return 400 or 500 error
		return c.JSON(http.StatusInternalServerError, err)
	}

	// need to define a struct for response
	return c.JSON(http.StatusOK, "success")
}

func (v *VoteServer) UnVote(c echo.Context) error {
	var event VoteEvent
	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := v.voteUseCase.UnVote(c.Request().Context(), event.BusinessId, event.MessageId, event.UserId)

	if err != nil {
		// TODO to differentiate error type, return 400 or 500 error
		return c.JSON(http.StatusInternalServerError, err)
	}

	// need to define a struct for response
	return c.JSON(http.StatusOK, "success")

}

func (v *VoteServer) Count(c echo.Context) error {
	businessId := c.Param("businessId")
	messageId := c.Param("messageId")

	count, err := v.voteUseCase.Count(c.Request().Context(), businessId, messageId)
	if err != nil {
		// TODO to differentiate error type, return 400 or 500 error
		return c.JSON(http.StatusInternalServerError, err)
	}

	// need to define a struct for response
	return c.JSON(http.StatusOK, count)
}

func (v *VoteServer) ListUser(c echo.Context) error {
	businessId := c.Param("businessId")
	messageId := c.Param("messageId")

	voteUsers, err := v.voteUseCase.VotedUsers(c.Request().Context(), businessId, messageId)
	if err != nil {
		// TODO to differentiate error type, return 400 or 500 error
		return c.JSON(http.StatusInternalServerError, err)
	}

	result := VoteUsersEvent{
		List: voteUsers,
	}

	return c.JSON(http.StatusOK, result)
}

func (v *VoteServer) register(g *echo.Group) error {
	g.GET("/vote/:businessId/:messageId", v.Count)
	g.GET("/list/:businessId/:messageId", v.ListUser)

	g.POST("/vote", v.Vote)
	g.POST("/unvote", v.UnVote)

	return nil
}

func (v *VoteServer) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	v.e.Start(fmt.Sprintf(":%s", port))
	return nil
}
