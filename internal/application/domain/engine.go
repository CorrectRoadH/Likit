package domain

const (
	SimpleVote string = "SIMPLE_VOTE"

	MiddleVote string = "MIDDLE_VOTE"
)

type Feature struct {
	Features []string
}

type Require struct {
	Database map[string]int
}
