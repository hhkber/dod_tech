package execution_mode

import "testing"

func TestA(t *testing.T) {
	t.Parallel()
	t.Log("A")
}

func TestB(t *testing.T) {
	t.Parallel()
	t.Log("B")
}

func TestC(t *testing.T) {
	t.Log("C")
}
