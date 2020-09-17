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
}

func assertEquals(t *testing.T, expected string, result string) {
	if result != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}
