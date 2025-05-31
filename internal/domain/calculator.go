package domain

type ScoreCalculation struct{}

func (sc *ScoreCalculation) CalculateScoreSKD(input Score) float32 {
	return input.twk + input.tiu + input.tkp/550*0.4
}

func (sc *ScoreCalculation) CalculateScoreSKB(input Score) float32 {
	return input.skb / 500 * 0.6
}

func (sc *ScoreCalculation) CalculateFinalScore(skd, skb float32) float32 {
	return skd + skb
}
