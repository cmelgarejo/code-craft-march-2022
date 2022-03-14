package rover

import "fmt"

const (
	X       = 0
	Y       = 1
	RIGHT   = "R"
	LEFT    = "L"
	FORWARD = "F"
	BACK    = "B"
	NORTH   = "N"
	SOUTH   = "S"
	EAST    = "E"
	WEST    = "W"
)

var coordinates = map[int]string{
	0: NORTH,
	1: EAST,
	2: SOUTH,
	3: WEST,
}

type Rover struct {
	position  []int
	direction string
}

func NewRover(position []int, direction string) Rover {
	return Rover{
		position:  position,
		direction: direction,
	}
}

func (r *Rover) Position() []int {
	return r.position
}

func (r *Rover) Direction() string {
	return r.direction
}

func (r *Rover) Move(cmds string) {
	for _, cmd := range cmds{
		cmd := string(cmd)

		// fmt.Printf("\nPRE -> CMD:%v, POS:%v, DIR:%v", cmd, r.position, r.direction)
		switch cmd {
		case LEFT,RIGHT:
			r.rotate(cmd)
		case BACK,FORWARD:
			r.move(cmd)
		default:
			fmt.Printf("invalid command %v \n", cmd)
		}
		// fmt.Printf("\nPOS -> CMD:%v, POS:%v, DIR:%v\n", cmd, r.position, r.direction)

	}
}

func (r *Rover) move(cmd string) {
	movement := 1
	if cmd == BACK {
		movement = -1
	}

	switch r.direction {
	case NORTH:
		r.position[Y] = r.position[Y] - movement
	case SOUTH:
		r.position[Y] = r.position[Y] + movement
	case EAST:
		r.position[X] = r.position[X] + movement
	case WEST:
		r.position[X] = r.position[X] - movement
	}
}


func (r *Rover) rotate(cmd string) {
	current := r.getDir()
	if cmd == LEFT {
		current -= 1
	} else if cmd == RIGHT {
		current += 1
	}

	if current > 3 {
		current = 0
	} else if current < 0 {
		current = 3
	}

	r.direction = coordinates[current]
}

func (r *Rover) getDir() int {
	for k, v := range coordinates {
		if v == r.direction {
			return k
		}
	}
	return 0 //default => NORTH
}
