package helpers

import "testing"

func TestInitSK(t *testing.T) {
	var _, err = InitSK(nil)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
