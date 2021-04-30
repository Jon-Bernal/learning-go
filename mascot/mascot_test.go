package mascot_test

import (
	"testing"

	"github.com/Jon-Bernal/learning-go/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Tux" {
		t.Fatal("Wrong mascot :(")
	}
}