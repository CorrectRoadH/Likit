package vote_test

import (
	"context"
	"math/rand"
	"strconv"
	"testing"

	"github.com/CorrectRoadH/Likit/config"
	"github.com/CorrectRoadH/Likit/internal/application/server"
	"github.com/CorrectRoadH/Likit/internal/port/in"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sourcegraph/conc"
	"github.com/sourcegraph/conc/iter"
)

func TestSimpleVote(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vote Suite")
}

var _ = Describe("Simple Vote Suite", func() {
	var vote_server in.VoteUseCase

	BeforeEach(func() {
		vote_server = server.NewSimpleVoteServer(
			config.TestEnvRedisConfig(),
		)
	})

	Describe("Vote a random value", func() {
		Context("Vote 100 count with 100 users", func() {
			It("the count should be 100", func() {
				ctx := context.Background()

				businessId := "test_vote_case1_" + strconv.Itoa(rand.Intn(100000))
				messageId := "messageId_" + strconv.Itoa(rand.Intn(100000))

				var wg conc.WaitGroup
				for i := 0; i < 100; i++ {
					wg.Go(func() {
						randomUser := "user" + strconv.Itoa(rand.Intn(100000))
						_, err := vote_server.Vote(ctx, businessId, messageId, randomUser)
						Expect(err).To(BeNil())
					})
				}
				wg.Wait()

				count, err := vote_server.Count(ctx, businessId, messageId)
				Expect(err).To(BeNil())
				Expect(count).To(Equal(100))
			})
		})

		Context("Vote with different users", func() {
			var voted_users []string
			var no_voted_users []string
			var ctx context.Context
			var businessId string
			var messageId string

			BeforeEach(func() {
				voted_users = []string{
					"user1",
					"user2",
					"user3",
					"user4",
				}

				no_voted_users = []string{
					"user5",
					"user6",
					"user7",
					"user8",
				}
				ctx = context.Background()

				businessId = "test_vote_case2_" + strconv.Itoa(rand.Intn(100000))
				messageId = "messageId_" + strconv.Itoa(rand.Intn(100000))

				iter.ForEach[string](voted_users, func(user *string) {
					_, err := vote_server.Vote(ctx, businessId, messageId, *user)
					Expect(err).To(BeNil())
				})

			})

			It("should be voted that voted user", func() {
				for _, user := range voted_users {
					voted, err := vote_server.IsVoted(ctx, businessId, messageId, user)
					Expect(err).To(BeNil())
					Expect(voted).To(Equal(true))
				}

				for _, user := range no_voted_users {
					voted, err := vote_server.IsVoted(ctx, businessId, messageId, user)
					Expect(err).To(BeNil())
					Expect(voted).To(Equal(false))
				}
			})

			It("should only contain the voted user", func() {
				for _, user := range voted_users {
					users, err := vote_server.VotedUsers(ctx, businessId, messageId)
					Expect(err).To(BeNil())
					Expect(users).To(ContainElement(user))
				}

				for _, user := range no_voted_users {
					users, err := vote_server.VotedUsers(ctx, businessId, messageId)
					Expect(err).To(BeNil())
					Expect(users).To(Not(ContainElement(user)))
				}
			})
		})

		Context("Vote with same users", func() {
			var voted_users []string
			var no_voted_users []string
			var ctx context.Context
			var businessId string
			var messageId string

			BeforeEach(func() {
				voted_users = []string{
					"user1",
					"user1",
					"user2",
					"user2",
					"user3",
					"user3",
					"user3",
					"user4",
				}

				no_voted_users = []string{
					"user5",
					"user6",
					"user7",
					"user8",
				}
				ctx = context.Background()

				businessId = "test_vote_case3_" + strconv.Itoa(rand.Intn(100000))
				messageId = "messageId_" + strconv.Itoa(rand.Intn(100000))

				iter.ForEach[string](voted_users, func(user *string) {
					_, err := vote_server.Vote(ctx, businessId, messageId, *user)
					Expect(err).To(BeNil())
				})

			})

			It("should be voted that voted user", func() {
				for _, user := range voted_users {
					voted, err := vote_server.IsVoted(ctx, businessId, messageId, user)
					Expect(err).To(BeNil())
					Expect(voted).To(Equal(true))
				}

				for _, user := range no_voted_users {
					voted, err := vote_server.IsVoted(ctx, businessId, messageId, user)
					Expect(err).To(BeNil())
					Expect(voted).To(Equal(false))
				}
			})

			It("should only contain the voted user", func() {
				for _, user := range voted_users {
					users, err := vote_server.VotedUsers(ctx, businessId, messageId)
					Expect(err).To(BeNil())
					Expect(users).To(ContainElement(user))
				}

				for _, user := range no_voted_users {
					users, err := vote_server.VotedUsers(ctx, businessId, messageId)
					Expect(err).To(BeNil())
					Expect(users).To(Not(ContainElement(user)))
				}
			})

			It("should only contain 4 voted user", func() {
				count, err := vote_server.Count(ctx, businessId, messageId)
				Expect(err).To(BeNil())
				Expect(count).To(Equal(4))
			})
		})

		Context("Vote with same users", func() {
			var voted_users []string
			var unvoted_users []string
			var ctx context.Context
			var businessId string
			var messageId string

			BeforeEach(func() {
				voted_users = []string{
					"user1",
					"user1",
					"user2",
					"user2",
					"user3",
					"user3",
					"user3",
					"user3",
					"user4",
				}

				unvoted_users = []string{
					"user5",
					"user3",
					"user2",
				}
				ctx = context.Background()

				businessId = "test_vote_case4_" + strconv.Itoa(rand.Intn(100000))
				messageId = "messageId_" + strconv.Itoa(rand.Intn(100000))

				iter.ForEach[string](voted_users, func(user *string) {
					_, err := vote_server.Vote(ctx, businessId, messageId, *user)
					Expect(err).To(BeNil())
				})

				iter.ForEach[string](unvoted_users, func(user *string) {
					_, err := vote_server.UnVote(ctx, businessId, messageId, *user)
					Expect(err).To(BeNil())
				})
			})

			It("should only contain 2 voted user", func() {
				count, err := vote_server.Count(ctx, businessId, messageId)
				Expect(err).To(BeNil())
				Expect(count).To(Equal(2))
			})
		})

	})

})
