package main

import "testing"

// TestReadJSON testing read from JSON file
func TestReadJSON(t *testing.T) {

	timers := getTimers()
	timer := timers[1]
	expectedName := "fingerboard"
	if expectedName != timer.Name {
		t.Errorf("timer.name == %q, want %q",
			timer.Name, expectedName)
	}

}
