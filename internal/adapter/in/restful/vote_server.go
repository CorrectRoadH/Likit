package restful

import (
	"fmt"
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

func (v *VoteServer) Vote(c echo.Context) error {
	// TODO get value from body

	businessId := c.Param("businessId")
	messageId := c.Param("messageId")
	userId := c.QueryParam("userId")

	return v.voteUseCase.Vote(c.Request().Context(), businessId, messageId, userId)
}

func (v *VoteServer) UnVote(c echo.Context) error {
	return nil
}

func (v *VoteServer) Count(c echo.Context) error {
	businessId := c.Param("businessId")
	messageId := c.Param("messageId")

	count, err := v.voteUseCase.Count(c.Request().Context(), businessId, messageId)
	if err != nil {
		// return 400 or 500 error
		return err
	}

	// need to define a struct for response
	return c.JSON(200, count)
}

func (v *VoteServer) ListUser(c echo.Context) error {
	return nil
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
