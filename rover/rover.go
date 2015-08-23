package rover

const (
    NORTH = iota
    EAST = iota
    SOUTH = iota
    WEST = iota
)

type Rover struct {
    Coords Coordinates
    Facing int
    Grid *Grid
}

type Coordinates struct {
    x int
    y int
}

func (initialPosition *Coordinates) Move(dX, dY int) Coordinates {
    newCoords := Coordinates{initialPosition.x + dX, initialPosition.y + dY}
    return newCoords
}

func New(coordinates Coordinates, facing int) *Rover {
    rover := new(Rover)
    rover.Facing = facing
    rover.Coords = coordinates
    rover.Grid = NewGrid(100,100)
    rover.Grid.Snap(rover)
    return rover
}

func (rover *Rover) Advance() {
    rover.Grid.Insert(rover.Coords, NOTHING)
    var dx, dy int

    switch direction := rover.Facing; direction {
    case NORTH:
        dy += 1
    case EAST:
        dx += 1
    case WEST:
        dx -= 1
    case SOUTH:
        dy -= 1
    }

    newCoords, ok := rover.Grid.OverflowPosition(rover.Coords, dx, dy)
    if !ok {

    }
    rover.Coords = newCoords
    rover.Grid.Snap(rover)
}

func (rover *Rover) Retreat() {
    rover.Grid.Insert(rover.Coords, NOTHING)
    var dx, dy int

    switch direction := rover.Facing; direction {
    case NORTH:
        dy -= 1
    case EAST:
        dx -= 1
    case WEST:
        dx += 1
    case SOUTH:
        dy += 1
    }

    newCoords, ok := rover.Grid.OverflowPosition(rover.Coords, dx, dy)
    if !ok {

    }
    rover.Coords = newCoords
    rover.Grid.Snap(rover)
}

func (rover *Rover) TurnRight() {
    switch direction := rover.Facing; direction {
    case NORTH:
        rover.Facing = EAST
    case EAST:
        rover.Facing = SOUTH
    case WEST:
        rover.Facing = NORTH
    case SOUTH:
        rover.Facing = WEST
    }
}

func (rover *Rover) TurnLeft() {
    switch direction := rover.Facing; direction {
    case NORTH:
        rover.Facing = WEST
    case EAST:
        rover.Facing = NORTH
    case WEST:
        rover.Facing = SOUTH
    case SOUTH:
        rover.Facing = EAST
    }
}
