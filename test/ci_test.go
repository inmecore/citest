package test

import "testing"

func TestCI(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Logf("%d", i)
	}
}
