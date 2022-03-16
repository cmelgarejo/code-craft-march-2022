package location

//go:generate mockgen -destination=../mocks/mock_location.go -package=mocks github.com/cmelgarejo/code-craft-march-2022/5-solve-a-problem/location DistanceCalculator

import "math"

type DistanceCalculator interface {
	GetDistance(lat1, long1, lat2, long2 float64) float64
}

type DistanceCalculatorImpl struct {
}

func (d *DistanceCalculatorImpl) GetDistance(lat1, long1, lat2, long2 float64) float64 {
	return math.Sqrt(math.Pow(lat1-lat2, 2) + math.Pow(long1-long2, 2))
}
