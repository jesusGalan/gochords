package gochords

import (
	"os"

	"github.com/jesusGalan/goscales"
	"github.com/jesusGalan/goutils"
)

//note_list = take_all_notes_from(tone)
//note_list = note_list + note_list
//scale_list = config_scale(note_list, scale_name, tone)

func main() {

}

func getScale(scale, tone string) []string {
	return []string{"hola", "mundo"}
}

func getScaleName() string {
	return os.Args[1]
}

func getNotesForTwoOctaves(notes []string) []string {
	return []string{notes[0], notes[1], notes[2], notes[3], notes[4], notes[5],
		notes[6], notes[0], notes[1], notes[2], notes[3], notes[4],
		notes[5], notes[6]}
}

func getChromaticNotesForTwoOctaves(notes []string) []string {
	return []string{notes[0], notes[1], notes[2], notes[3], notes[4], notes[5],
		notes[6], notes[7], notes[8], notes[9], notes[10], notes[11],
		notes[0], notes[1], notes[2], notes[3], notes[4], notes[5],
		notes[6], notes[7], notes[8], notes[9], notes[10], notes[11]}
}

/*TakeNotesFromTone must return a chromatic scale and begin with chosen tone.*/
func TakeNotesFromTone(tone string) []string {
	notes := []string{"C", "C#.Db", "D", "D#.Eb", "E", "F", "F#.Gb", "G", "G#.Ab", "A", "A#.Bb", "B"}
	notes = getChromaticNotesForTwoOctaves(notes)

	noteMap := make(map[string]int)

	noteMap["C"] = 0
	noteMap["C#"] = 1
	noteMap["Db"] = 1
	noteMap["D"] = 2
	noteMap["D#"] = 3
	noteMap["Eb"] = 3
	noteMap["E"] = 4
	noteMap["F"] = 5
	noteMap["F#"] = 6
	noteMap["Gb"] = 6
	noteMap["G"] = 7
	noteMap["G#"] = 8
	noteMap["Ab"] = 8
	noteMap["A"] = 9
	noteMap["A#"] = 10
	noteMap["Bb"] = 10
	noteMap["B"] = 11

	posTone := noteMap[tone]
	newNotes := []string{notes[posTone], notes[posTone+1], notes[posTone+2],
		notes[posTone+3], notes[posTone+4], notes[posTone+5], notes[posTone+6],
		notes[posTone+7], notes[posTone+8], notes[posTone+9], notes[posTone+10],
		notes[posTone+11]}

	return newNotes
}

/*ShowBasicScaleNotes return the notes without being processed*/
func ShowBasicScaleNotes(tone, scaleName string, degree int) []string {
	chromaticOnTone := TakeNotesFromTone(tone)

	scaleFormula := goscales.GetAll(goscales.TestRepo())

	naturalScalePositions := goutils.SumArrayOfStringsAsInt(
		scaleFormula[scaleName][goutils.ParseIntToString(degree)]["formula"])

	x := 0
	y := 0
	scaleNotes := make([]string, 0, len(naturalScalePositions))

	for x < len(chromaticOnTone) {
		if x == 0 {
			scaleNotes = append(scaleNotes, tone)
		} else {
			for y < len(naturalScalePositions)-1 {

				if x == goutils.ParseStringToInt(naturalScalePositions[y]) {
					scaleNotes = append(scaleNotes, chromaticOnTone[goutils.ParseStringToInt(naturalScalePositions[y])])
				}
				y = y + 1
			}
		}
		y = 0
		x = x + 1
	}

	return scaleNotes
}

/*Scale is an Object containing properties of a scale*/
type Scale struct {
	degree                    int
	name, degreeName, tone    string
	formula, intervals, notes []string
}

/*DoScale can return a object that contains all about a specific scale*/
func DoScale(tone, scaleName string, degree int) Scale {
	allScales := goscales.GetAll(goscales.TestRepo())
	scaleFormula := allScales[scaleName][goutils.ParseIntToString(degree)]["formula"]
	degreeName := allScales[scaleName][goutils.ParseIntToString(degree)]["name"][0]
	scaleIntervals := allScales[scaleName][goutils.ParseIntToString(degree)]["intervals"]

	var scale Scale
	scale.name = scaleName
	scale.tone = tone
	scale.degreeName = degreeName
	scale.degree = degree
	scale.formula = scaleFormula
	scale.notes = ShowBasicScaleNotes(tone, scaleName, degree)
	scale.intervals = scaleIntervals

	return scale

}

/*WhatScaleFor should return the scales that contains all elements in
a list of intervals.*/
func WhatScaleFor(positions []string) []string {

	scalesJSON := goscales.GetAll(goscales.TestRepo())
	var scaleList []string
	count := 0
	for key := range scalesJSON {

		for x := range scalesJSON[key] {
			count = 0
			for y := range scalesJSON[key][x]["intervals"] {

				for givenInterval := range positions {

					if positions[givenInterval] == scalesJSON[key][x]["intervals"][y] {
						count = count + 1
					} else {
						sameElement := GetSameElement(positions[givenInterval])

						if sameElement == scalesJSON[key][x]["intervals"][y] {
							count = count + 1
						}
					}
				}

			}
			if count == len(positions) {
				scaleList = append(scaleList, scalesJSON[key][x]["name"][0])
			}
		}

	}
	return scaleList
}

/*GetSameElement is a method for traspose musical interval nomenclature
to equivalent one*/
func GetSameElement(interval string) string {
	var sameElem string
	switch {
	case interval == "1":
		sameElem = "1"

	case interval == "b2":
		sameElem = "b9"

	case interval == "2":
		sameElem = "bb3"

	case interval == "b3":
		sameElem = "#2"

	case interval == "3":
		sameElem = "b4"

	case interval == "4":
		sameElem = "bb5"

	case interval == "b5":
		sameElem = "#4"

	case interval == "5":
		sameElem = "bb6"

	case interval == "b6":
		sameElem = "#5"

	case interval == "6":
		sameElem = "bb7"

	case interval == "b7":
		sameElem = "#6"

	case interval == "7":
		sameElem = "7"

	case interval == "b9":
		sameElem = "b2"

	case interval == "bb3":
		sameElem = "2"

	case interval == "#2":
		sameElem = "b3"

	case interval == "b4":
		sameElem = "3"

	case interval == "bb5":
		sameElem = "4"

	case interval == "#4":
		sameElem = "b5"

	case interval == "bb6":
		sameElem = "5"

	case interval == "#5":
		sameElem = "b6"

	case interval == "bb7":
		sameElem = "6"

	case interval == "#6":
		sameElem = "b7"
	}
	if sameElem == "" {
		return "no se que es " + interval
	}
	return sameElem

}
