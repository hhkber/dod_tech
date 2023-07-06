package short_mode

import "testing"

func TestLongRunning(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	t.Log("long running test")
}

func TestShortRunning(t *testing.T) {
	t.Log("short running test")
}
