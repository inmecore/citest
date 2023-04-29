package test

import (
	"fmt"
	"testing"
)

func TestCI(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Logf("%d", i)
		})
	}
}
