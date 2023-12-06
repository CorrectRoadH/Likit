package system_test

import (
	"context"
	"math/rand"
	"strconv"
	"testing"

	"github.com/CorrectRoadH/Likit/config"
	"github.com/CorrectRoadH/Likit/internal/adapter/out/admin"
	"github.com/CorrectRoadH/Likit/internal/application/server"
	"github.com/CorrectRoadH/Likit/internal/port/in"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInitSystem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vote Suite")
}

var _ = Describe("System Init Vote Suite", func() {
	var voteServer in.VoteUseCase

	BeforeEach(func() {
		voteServer, _ = server.NewVoteServer(
			server.NewAdminServer(
				admin.NewBusinessAdapter(config.ProductEnvConfigDatabaseConfig()),
				config.TestEnvRedisConfig(),
				admin.NewUserAdapter(config.ProductEnvConfigDatabaseConfig()),
			),
		)
	})

	Describe("Vote and Count", func() {
		Context("Count and Vote and Count", func() {
			It("the count should be 0", func() {
				messageId := "test_vote_case1_" + strconv.Itoa(rand.Intn(100000))
				count, err := voteServer.Count(context.Background(), "COMMENT_LIKE", messageId)
				Expect(err).To(BeNil())
				Expect(count).To(Equal(0))

				voteServer.Vote(context.Background(), "COMMENT_LIKE", messageId, "134")

				count, err = voteServer.Count(context.Background(), "COMMENT_LIKE", messageId)
				Expect(err).To(BeNil())
				Expect(count).To(Equal(1))
			})
		})
	})

})
