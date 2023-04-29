package test

import (
	"citest/tools"
	"fmt"
	"testing"
)

func TestCI(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			tools.PrintArgs(i)
		})
	}
	t.Log("Success.")
}
