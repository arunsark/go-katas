package rover

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestShouldCreateNewGrid(t *testing.T) {
    grid := NewGrid(4,5)
    assert.Equal(t, len(grid.Field), 4)
    assert.Equal(t, [][]int{
        []int{0, 0, 0, 0, 0},
        []int{0, 0, 0, 0, 0},
        []int{0, 0, 0, 0, 0},
        []int{0, 0, 0, 0, 0}}, grid.Field)
}

func TestShouldSnapRoverToGrid(t *testing.T) {
    grid := NewGrid(4, 5)
    rover := New(Coordinates{1, 2}, NORTH)

    grid.Snap(rover)
    assert.Equal(t, grid.At(rover.Coords), ROVER)
}

func TestShouldOverflowWidth(t *testing.T) {
    grid := NewGrid(4, 4)
    newPosition, _ := grid.OverflowPosition(Coordinates{3, 3}, 1, 0)
    assert.Equal(t, newPosition, Coordinates{0, 3})
}

func TestShouldOverflowHeight(t *testing.T) {
    grid := NewGrid(4,4)
    newPosition, _ := grid.OverflowPosition(Coordinates{3,3}, 0, 1)
    assert.Equal(t, newPosition, Coordinates{3,0})
}

func TestShouldReturnTrueWhenNoCollisionImminent(t *testing.T) {
    grid := NewGrid(4,4)
    grid.Insert(Coordinates{2,1}, NOTHING)
    newPosition, ok := grid.OverflowPosition(Coordinates{1,1}, 1, 0)
    assert.Equal(t, newPosition, Coordinates{2,1})
    assert.Equal(t, true, ok)
}

func TestShouldReturnFalseWhenCollisionImminent(t *testing.T) {
    grid := NewGrid(4,4)
    grid.Insert(Coordinates{2,1}, OBSTACLE)
    _, ok := grid.OverflowPosition(Coordinates{1,1}, 1, 0)
    assert.Equal(t, false, ok)
}

func TestShouldNotMoveWhenCollisionImminent(t *testing.T) {
    grid := NewGrid(4, 4)
    rover := New(Coordinates{0, 0}, NORTH)
    rover.Grid = grid

    grid.Snap(rover)
    rover.Grid.Insert(Coordinates{0, 1}, OBSTACLE)

    rover.Advance()
    assert.Equal(t, Coordinates{0,0}, rover.Coords)

}
