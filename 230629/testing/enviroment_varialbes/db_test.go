//go:build integration
// +build integration

package ennvironment_varialbes

import (
	"os"
	"testing"
)

func TestInsert(t *testing.T) {
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip("skipping integration test")
	}

	t.Log("insert test")
}
