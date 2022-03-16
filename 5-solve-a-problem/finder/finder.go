package finder

import (
	"github.com/cmelgarejo/code-craft-march-2022/5-solve-a-problem/location"
	"github.com/cmelgarejo/code-craft-march-2022/5-solve-a-problem/spot"
)

type FinderService interface {
	FindFreeSpots(lat, long float64) ([]spot.Spot, error)
}

type FinderServiceImpl struct {
	spots       []spot.Spot
	locationSvc location.DistanceCalculator
}

func (fs *FinderServiceImpl) FindFreeSpots(lat, long float64) ([]spot.Spot, error) {
	freeSpots := []spot.Spot{}
	for _, spot := range fs.spots {
		if spot.Available && fs.locationSvc.GetDistance(lat, long, spot.Lat, spot.Long) <= 1 {
			freeSpots = append(freeSpots, spot)
		}
	}
	return freeSpots, nil
}

func NewFinder(spots []spot.Spot, locationSvc location.DistanceCalculator) FinderService {
	return &FinderServiceImpl{spots: spots, locationSvc: locationSvc}
}
