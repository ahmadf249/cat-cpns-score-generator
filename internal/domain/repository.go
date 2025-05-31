package domain

type ScoreRepository interface {
	GenerateScore(input Score) (*ScoreResult, error)
}