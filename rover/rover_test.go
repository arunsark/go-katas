package rover

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewRover(t *testing.T) {
    rover := New(Coordinates{0,0}, SOUTH)
    assert.Equal(t, Coordinates{0,0}, rover.Coords)
}

func TestRoverShouldBeFacingTheInitialDirection(t *testing.T) {
    rover := New(Coordinates{0,0}, SOUTH)
    assert.Equal(t, SOUTH, rover.Facing)
}

func TestShouldAdvanceNorth(t *testing.T) {
    rover := New(Coordinates{0, 0}, NORTH)
    rover.Advance()
    assert.Equal(t, rover.Coords, Coordinates{0, 1})
}

func TestShouldAdvanceEast(t *testing.T) {
    rover := New(Coordinates{0, 0}, EAST)
    rover.Advance()
    assert.Equal(t, rover.Coords, Coordinates{1, 0})
}

func TestShouldAdvanceWest(t *testing.T) {
    rover := New(Coordinates{1, 0}, WEST)
    rover.Advance()
    assert.Equal(t, rover.Coords, Coordinates{0, 0})
}

func TestShouldAdvanceSouth(t *testing.T) {
    rover := New(Coordinates{0, 1}, SOUTH)
    rover.Advance()
    assert.Equal(t, rover.Coords, Coordinates{0, 0})
}

func TestShouldRetreatNorth(t *testing.T) {
    rover := New(Coordinates{0, 1}, NORTH)
    rover.Retreat()
    assert.Equal(t, rover.Coords, Coordinates{0, 0})
}

func TestShouldRetreatEast(t *testing.T) {
    rover := New(Coordinates{1, 0}, EAST)
    rover.Retreat()
    assert.Equal(t, rover.Coords, Coordinates{0, 0})
}

func TestShouldRetreatWest(t *testing.T) {
    rover := New(Coordinates{0, 0}, WEST)
    rover.Retreat()
    assert.Equal(t, rover.Coords, Coordinates{1, 0})
}

func TestShouldRetreatSouth(t *testing.T) {
    rover := New(Coordinates{0, 0}, SOUTH)
    rover.Retreat()
    assert.Equal(t, rover.Coords, Coordinates{0, 1})
}

func TestShouldTurnRight(t *testing.T) {
    rover := New(Coordinates{0, 0}, NORTH)
    rover.TurnRight()
    assert.Equal(t, rover.Facing, EAST)
    rover.TurnRight()
    assert.Equal(t, rover.Facing, SOUTH)
    rover.TurnRight()
    assert.Equal(t, rover.Facing, WEST)
    rover.TurnRight()
    assert.Equal(t, rover.Facing, NORTH)
}

func TestShouldTurnLeft(t *testing.T) {
    rover := New(Coordinates{0, 0}, NORTH)
    rover.TurnLeft()
    assert.Equal(t, rover.Facing, WEST)
    rover.TurnLeft()
    assert.Equal(t, rover.Facing, SOUTH)
    rover.TurnLeft()
    assert.Equal(t, rover.Facing, EAST)
    rover.TurnLeft()
    assert.Equal(t, rover.Facing, NORTH)
}

func TestShouldStaySnappedToGrid(t *testing.T) {
    rover := New(Coordinates{0, 0}, NORTH)
    rover.Advance()
    assert.Equal(t, rover.Grid.At(Coordinates{0, 0}), NOTHING)
    assert.Equal(t, rover.Grid.At(rover.Coords), ROVER)
}

func TestShouldAdvanceWithOverflow(t *testing.T) {
    rover := New(Coordinates{99, 99}, NORTH)
    rover.Advance()
    assert.Equal(t, rover.Coords, Coordinates{99, 0})
}

func TestShouldRetreatWithOverflow(t *testing.T) {
    rover := New(Coordinates{99, 99}, EAST)
    rover.Advance()
    assert.Equal(t, rover.Coords, Coordinates{0, 99})
}

func TestShouldReportObstacle(t *testing.T) {
    rover := New(Coordinates{0, 0}, NORTH)
    rover.Grid.Insert(Coordinates{1, 1}, OBSTACLE)
    rover.Advance()
    rover.TurnRight()
    rover.Advance()
    assert.Equal(t, rover.Coords, Coordinates{0, 1})
}
