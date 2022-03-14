package test

import (
	"strings"
	"testing"

	"github.com/cmelgarejo/code-craft-march-2022/1-mars-rover/rover"
	"github.com/stretchr/testify/assert"
)

func TestRover_Init(t *testing.T) {

	marsrover := rover.NewRover([]int{1, 1}, rover.NORTH)

	assert.Equal(t, []int{1, 1}, marsrover.Position())
	assert.Equal(t, rover.NORTH, marsrover.Direction())
}

func TestRover_Movement(t *testing.T) {
	tcs := []struct {
		Name        string
		InitPos     []int
		InitDir     string
		Commands    string
		ExpectedPos []int
		ExpectedDir string
	}{
		{
			Name:        "move forward north",
			InitPos:     []int{1, 1},
			InitDir:     rover.NORTH,
			Commands:    "F",
			ExpectedPos: []int{1, 0},
			ExpectedDir: rover.NORTH,
		},
		{
			Name:        "move forward south",
			InitPos:     []int{1, 1},
			InitDir:     rover.SOUTH,
			Commands:    "F",
			ExpectedPos: []int{1, 2},
			ExpectedDir: rover.SOUTH,
		},
		{
			Name:        "move forward east",
			InitPos:     []int{1, 1},
			InitDir:     rover.EAST,
			Commands:    "F",
			ExpectedPos: []int{2, 1},
			ExpectedDir: rover.EAST,
		},
		{
			Name:        "move forward west",
			InitPos:     []int{1, 1},
			InitDir:     rover.WEST,
			Commands:    "F",
			ExpectedPos: []int{0, 1},
			ExpectedDir: rover.WEST,
		},
		//
		{
			Name:        "move back north",
			InitPos:     []int{1, 1},
			InitDir:     rover.NORTH,
			Commands:    rover.BACK,
			ExpectedPos: []int{1, 2},
			ExpectedDir: rover.NORTH,
		},
		{
			Name:        "move back south",
			InitPos:     []int{1, 1},
			InitDir:     rover.SOUTH,
			Commands:    rover.BACK,
			ExpectedPos: []int{1, 0},
			ExpectedDir: rover.SOUTH,
		},
		{
			Name:        "move back east",
			InitPos:     []int{1, 1},
			InitDir:     rover.EAST,
			Commands:    rover.BACK,
			ExpectedPos: []int{0, 1},
			ExpectedDir: rover.EAST,
		},
		{
			Name:        "move back west",
			InitPos:     []int{1, 1},
			InitDir:     rover.WEST,
			Commands:    rover.BACK,
			ExpectedPos: []int{2, 1},
			ExpectedDir: rover.WEST,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			marsrover := rover.NewRover(tc.InitPos, tc.InitDir)
			marsrover.Move(tc.Commands)
			assert.Equal(t, tc.ExpectedPos, marsrover.Position())
			assert.Equal(t, tc.ExpectedDir, marsrover.Direction())
		})
	}

}

func TestRover_Direction(t *testing.T) {
	tcs := []struct {
		Name        string
		InitPos     []int
		InitDir     string
		Commands    string
		ExpectedPos []int
		ExpectedDir string
	}{
		{
			Name:        "turn right from north",
			InitPos:     []int{1, 1},
			InitDir:     rover.NORTH,
			Commands:    "R",
			ExpectedPos: []int{1, 1},
			ExpectedDir: rover.EAST,
		},
		{
			Name:        "turn left from north",
			InitPos:     []int{1, 1},
			InitDir:     rover.NORTH,
			Commands:    "L",
			ExpectedPos: []int{1, 1},
			ExpectedDir: rover.WEST,
		},
		{
			Name:        "turn right from west",
			InitPos:     []int{1, 1},
			InitDir:     rover.WEST,
			Commands:    "R",
			ExpectedPos: []int{1, 1},
			ExpectedDir: rover.NORTH,
		},
		{
			Name:        "turn left from west",
			InitPos:     []int{1, 1},
			InitDir:     rover.WEST,
			Commands:    "L",
			ExpectedPos: []int{1, 1},
			ExpectedDir: rover.SOUTH,
		},
		{
			Name:        "turn right from east",
			InitPos:     []int{1, 1},
			InitDir:     rover.EAST,
			Commands:    "R",
			ExpectedPos: []int{1, 1},
			ExpectedDir: rover.SOUTH,
		},
		{
			Name:        "turn left from east",
			InitPos:     []int{1, 1},
			InitDir:     rover.EAST,
			Commands:    "L",
			ExpectedPos: []int{1, 1},
			ExpectedDir: rover.NORTH,
		},
		{
			Name:        "turn right from south",
			InitPos:     []int{1, 1},
			InitDir:     rover.SOUTH,
			Commands:    "R",
			ExpectedPos: []int{1, 1},
			ExpectedDir: rover.WEST,
		},
		{
			Name:        "turn left from south",
			InitPos:     []int{1, 1},
			InitDir:     rover.SOUTH,
			Commands:    "L",
			ExpectedPos: []int{1, 1},
			ExpectedDir: rover.EAST,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			marsrover := rover.NewRover(tc.InitPos, tc.InitDir)
			marsrover.Move(tc.Commands)
			assert.Equal(t, tc.ExpectedPos, marsrover.Position())
			assert.Equal(t, tc.ExpectedDir, marsrover.Direction())
		})
	}
}

func TestRover_Commands(t *testing.T) {
	tcs := []struct {
		Name        string
		InitPos     []int
		InitDir     string
		Commands    string
		ExpectedPos []int
		ExpectedDir string
	}{
		{
			Name:        "single command",
			InitPos:     []int{1, 1},
			InitDir:     rover.NORTH,
			Commands:    rover.RIGHT,
			ExpectedPos: []int{1, 1},
			ExpectedDir: rover.EAST,
		},
		{
			Name:        "single invalid command",
			InitPos:     []int{1, 1},
			InitDir:     rover.NORTH,
			Commands:    "P",
			ExpectedPos: []int{1, 1},
			ExpectedDir: rover.NORTH,
		},
		{
			Name:        "multiple invalid commands",
			InitPos:     []int{1, 1},
			InitDir:     rover.NORTH,
			Commands:    "PEEPEE",
			ExpectedPos: []int{1, 1},
			ExpectedDir: rover.NORTH,
		},
		{
			Name:        "multiple valid commands",
			InitPos:     []int{1, 1},
			InitDir:     rover.EAST,
			Commands:    strings.Join([]string{rover.FORWARD, rover.LEFT, rover.FORWARD}, ""),
			ExpectedPos: []int{2, 0},
			ExpectedDir: rover.NORTH,
		},
		{
			Name:        "multiple valid commands 2",
			InitPos:     []int{1, 1},
			InitDir:     rover.NORTH,
			Commands:    strings.Join([]string{rover.FORWARD, rover.LEFT, rover.FORWARD}, ""),
			ExpectedPos: []int{0, 0},
			ExpectedDir: rover.WEST,
		},
		{
			Name:        "multiple valid commands 3",
			InitPos:     []int{1, 1},
			InitDir:     rover.NORTH,
			Commands:    strings.Join([]string{rover.RIGHT, rover.FORWARD, rover.LEFT, rover.FORWARD}, ""),
			ExpectedPos: []int{2, 0},
			ExpectedDir: rover.NORTH,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			marsrover := rover.NewRover(tc.InitPos, tc.InitDir)
			marsrover.Move(tc.Commands)
			assert.Equal(t, tc.ExpectedPos, marsrover.Position())
			assert.Equal(t, tc.ExpectedDir, marsrover.Direction())
		})
	}
}
