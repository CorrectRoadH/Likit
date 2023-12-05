package in

import "github.com/CorrectRoadH/Likit/internal/application/domain"

type BasicVoteUseCase interface {
	CheckRequire(config domain.Config) error
}
