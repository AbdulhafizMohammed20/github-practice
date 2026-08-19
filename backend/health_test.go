package backend

import "testing"

func TestHealth(t *testing.T) {
	result := Health()

	if result.Status != "ok" {
		t.Errorf("expected status to be ok, got %s", result.Status)
	}
}