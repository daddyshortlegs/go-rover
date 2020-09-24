package rover

import (
	"testing"
)

func TestNotMoving(t *testing.T) {
	assertEquals(t, "0:0:N", Execute(""))
}

func TestMoving(t *testing.T) {
	assertEquals(t, "0:1:N", Execute("M"))
	assertEquals(t, "0:2:N", Execute("MM"))
	assertEquals(t, "0:5:N", Execute("MMMMM"))
}

func TestWrapsIfGoingNorth(t *testing.T) {
	assertEquals(t, "0:0:N", Execute("MMMMMMMMMM"))
}

func TestRotatingRight(t *testing.T) {
	assertEquals(t, "0:0:E", Execute("R"))
	assertEquals(t, "0:0:S", Execute("RR"))
	assertEquals(t, "0:0:W", Execute("RRR"))
	assertEquals(t, "0:0:N", Execute("RRRR"))
	assertEquals(t, "0:0:E", Execute("RRRRR"))
}

func TestRotatingLeft(t *testing.T) {
	assertEquals(t, "0:0:W", Execute("L"))
	assertEquals(t, "0:0:S", Execute("LL"))
	assertEquals(t, "0:0:E", Execute("LLL"))
	assertEquals(t, "0:0:N", Execute("LLLL"))
	assertEquals(t, "0:0:W", Execute("LLLLL"))
}

func TestMovingEast(t *testing.T) {
	assertEquals(t, "1:0:E", Execute("RM"))
	assertEquals(t, "3:0:E", Execute("RMMM"))
}

func TestWrapsIfGoingEast(t *testing.T) {
	assertEquals(t, "0:0:E", Execute("RMMMMMMMMMM"))
	assertEquals(t, "2:0:E", Execute("RMMMMMMMMMMMM"))
}

func TestMovingWest(t *testing.T) {
	assertEquals(t, "9:0:W", Execute("LM"))
	assertEquals(t, "5:0:W", Execute("LMMMMM"))
	assertEquals(t, "0:0:W", Execute("LMMMMMMMMMM"))
}

func TestMovingSouth(t *testing.T) {
	assertEquals(t, "0:9:S", Execute("LLM"))
	assertEquals(t, "0:5:S", Execute("LLMMMMM"))
	assertEquals(t, "0:0:S", Execute("LLMMMMMMMMMM"))
}

func assertEquals(t *testing.T, expected string, result string) {
	if result != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}
