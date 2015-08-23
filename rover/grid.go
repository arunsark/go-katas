package rover

const (
    NOTHING = iota
    OBSTACLE = iota
    ROVER = iota
)

type Grid struct {
    Width int
    Height int
    Field [][]int
}

func NewGrid(width, height int) *Grid {
    grid := new(Grid)
    grid.Width = width
    grid.Height = height
    grid.Field = make([][]int, width)
    for index := range grid.Field {
        grid.Field[index] = make([]int, height)
    }
    return grid
}

func (grid *Grid)Snap(rover *Rover) *Grid {
    grid.Insert(rover.Coords, ROVER)
    return grid
}

func (grid *Grid)Insert(coords Coordinates, value int) {
    grid.Field[coords.x][coords.y] = value
}

func (grid *Grid)At(coords Coordinates) int{
    return grid.Field[coords.x][coords.y]
}

func (grid *Grid)OverflowPosition(coords Coordinates, dX, dY int) (Coordinates, bool) {
    newCoords := coords.Move(dX, dY)
    if newCoords.x >= grid.Width {
        newCoords.x -= grid.Width
    }
    if newCoords.y >= grid.Height {
        newCoords.y -= grid.Height
    }

    if grid.Field[newCoords.x][newCoords.y] != NOTHING {
        return coords, false
    }
    return newCoords, true
}
