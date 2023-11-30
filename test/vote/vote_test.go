package vote_test

import (
	"context"
	"math/rand"
	"strconv"
	"testing"

	"github.com/CorrectRoadH/Likit/config"
	v1 "github.com/CorrectRoadH/Likit/internal/adapter/out/v1"
	"github.com/CorrectRoadH/Likit/internal/application/server"
	"github.com/CorrectRoadH/Likit/internal/port/in"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sourcegraph/conc"
)

func TestVOte(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vote Suite")
}

var _ = Describe("Vote Suite", func() {
	var vote_server in.VoteUseCase

	BeforeEach(func() {
		vote_server = server.NewVoteServer(
			v1.NewRedisAdapter(
				config.TestEnvRedisConfig(),
			),
		)
	})

	Describe("Vote a random value", func() {
		Context("Vote 100 count with 100 users", func() {
			It("the count should be 100", func() {
				ctx := context.Background()

				businessId := "test_vote_businessId_" + strconv.Itoa(rand.Intn(100000))
				messageId := "messageId_" + strconv.Itoa(rand.Intn(100000))

				var wg conc.WaitGroup
				for i := 0; i < 100; i++ {
					wg.Go(func() {
						randomUser := "user" + strconv.Itoa(rand.Intn(100))
						vote_server.Vote(ctx, businessId, messageId, randomUser)
					})
				}
				wg.Wait()

				count, err := vote_server.Count(ctx, businessId, messageId)
				Expect(err).To(BeNil())
				Expect(count).To(Equal(100))
			})
		})

		Context("with fewer than 300 pages", func() {
			It("should be a short story", func() {
			})
		})
	})

})
