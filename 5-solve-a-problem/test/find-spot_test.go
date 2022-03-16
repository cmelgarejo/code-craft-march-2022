package test

import (
	"testing"

	"github.com/cmelgarejo/code-craft-march-2022/5-solve-a-problem/finder"
	"github.com/cmelgarejo/code-craft-march-2022/5-solve-a-problem/mocks"
	"github.com/cmelgarejo/code-craft-march-2022/5-solve-a-problem/spot"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

/*
[Parque 14 y Medio, shopping mall, -25.3390268,-57.4966736, 5, 12]
[Shopping San Lorenzo, shopping mall,-25.3369987,-57.4987932, 1, 12]
[Desert parking, desert hall parking lot,-27.3369987,-57.4987932, 1, 12]
*/
var testSpots = []spot.Spot{
	{
		Name:        "Parque 14 y Medio",
		Description: "Shopping mall",
		Lat:         -25.3390268,
		Long:        -57.4966736,
		Available:   true,
		Price:       5,
		Size:        12,
		ImageURLs: []string{
			"image1.jpg",
			"image2.jpg",
			"image3.jpg",
		},
	},
	{
		Name:        "Shopping San Lorenzo",
		Description: "Shopping mall",
		Lat:         -25.3369987,
		Long:        -57.4987932,
		Available:   false,
		Price:       1,
		Size:        12,
		ImageURLs: []string{
			"image1.jpg",
			"image2.jpg",
			"image3.jpg",
		},
	},
	{
		Name:        "Desert parking",
		Description: "Desert hall parking lot",
		Lat:         -27.3369987,
		Long:        -57.4987932,
		Available:   true,
		Price:       1,
		Size:        12,
		ImageURLs: []string{
			"image1.jpg",
			"image2.jpg",
			"image3.jpg",
		},
	},
}

func TestFindSpot_Available1KmRadius(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLocationSvc := mocks.NewMockDistanceCalculator(ctrl)
	finderSvc := finder.NewFinder(testSpots, mockLocationSvc)
	lat, long := -25.3390268, -57.4976736 // user location
	mockLocationSvc.EXPECT().GetDistance(lat, long, gomock.Any(),gomock.Any()).Return(0.0)
	mockLocationSvc.EXPECT().GetDistance(lat, long, gomock.Any(),gomock.Any()).Return(3.0)
	expected := []spot.Spot{testSpots[0]}
	freeSpots, err := finderSvc.FindFreeSpots(lat, long)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	assert.GreaterOrEqual(t, len(freeSpots), 1)
	assert.ElementsMatch(t, expected, freeSpots)
}

func TestFindSpot_NotAvailable1KmRadius(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLocationSvc := mocks.NewMockDistanceCalculator(ctrl)
	finderSvc := finder.NewFinder(testSpots, mockLocationSvc)
	lat, long := -25.3390268, -57.5966736 // user location

	mockLocationSvc.EXPECT().GetDistance(lat, long, gomock.Any(),gomock.Any()).Return(10.0).AnyTimes()
	freeSpots, err := finderSvc.FindFreeSpots(lat, long)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	assert.Equal(t, 0, len(freeSpots))
}
