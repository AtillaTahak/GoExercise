package main

import (
	"reflect"
	"testing"
)

func TestPlayer(t *testing.T) {
	player := Player{"John", 100}
	player2 := Player{"John2", 200}
	if !reflect.DeepEqual(player, player2) {
		t.Error("Expected true, got ", false)
	}
}
// we can out cached go test -v ./ -v -run TestCalculateValues -count=1
func TestCalculateValues(t *testing.T) {
	result := calculateValues(1, 2)
	if result != 4 {
		t.Error("Expected 3, got ", result)
	}
}