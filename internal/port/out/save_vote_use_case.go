package out

type SaveVoteUseCase interface {
	Vote(businessId string, messageId string, userId string) error

	Count(businessId, messageId, userId string) (int, error)
}
