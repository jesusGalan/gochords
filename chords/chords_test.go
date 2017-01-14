package gochords_test

import "testing"
import "github.com/jesusGalan/gochords"
import "github.com/jesusGalan/goutils"

func TestGetNotesFromToneSucceed(t *testing.T) {
	tone := "D"
	chromaticScale := gochords.TakeNotesFromTone(tone)
	expectedScale := []string{"D", "D#.Eb", "E", "F", "F#.Gb", "G", "G#.Ab", "A", "A#.Bb", "B", "C", "C#.Db"}
	result := goutils.CompareTwoArraysOfStrings(chromaticScale, expectedScale)
	if !result {
		t.Error("Expected scale and returned scale are not equal")
	}

}

func TestGetNotesFromToneError(t *testing.T) {
	tone := "D"
	chromaticScale := gochords.TakeNotesFromTone(tone)
	expectedScale := []string{"C#.Db", "D", "D#.Eb", "E", "F", "F#.Gb", "G", "G#.Ab", "A", "A#.Bb", "B", "C"}

	result := goutils.CompareTwoArraysOfStrings(chromaticScale, expectedScale)

	if result {
		t.Error("Expected scale and returned scale are equal")
	}

}
