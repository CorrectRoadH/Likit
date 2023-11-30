package vote_test

import (
	"context"
	"math/rand"
	"strconv"

	"github.com/CorrectRoadH/Likit/config"
	v1 "github.com/CorrectRoadH/Likit/internal/adapter/out/v1"
	"github.com/CorrectRoadH/Likit/internal/application/server"
	"github.com/CorrectRoadH/Likit/internal/port/in"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sourcegraph/conc"
)

var _ = Describe("Vote", func() {
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

			ctx := context.Background()

			var wg conc.WaitGroup
			for i := 0; i < 10; i++ {
				wg.Go(func() {
					randomUser := "user" + strconv.Itoa(rand.Intn(100))
					vote_server.Vote(ctx, "businessId_1", "messageId_1", randomUser)
				})
			}
			wg.Wait()

			It("should be a novel", func() {
				count, err := vote_server.Count(ctx, "businessId_1", "messageId_1")
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
