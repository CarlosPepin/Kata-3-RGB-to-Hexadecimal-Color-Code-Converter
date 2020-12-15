package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	x := 1
	y := 2
	n := 3

	total := Sum(x, y)
	if total != n {
		t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", x, y, total, n)
	}
}

func TestEntryType(t *testing.T) {
	entrySample := RGB{11, 33, 243}
	entrySampleType := reflect.TypeOf(entrySample).Name()
	entrySampleExpect := reflect.TypeOf(RGB{}).Name()

	if entrySampleType != entrySampleExpect {
		t.Errorf("The data type of entry of %v was incorrect, got: %v, want: %v.", entrySample, entrySampleType, entrySampleExpect)
	}
}

func TestOutputType(t *testing.T) {
	entrySample := RGB{11, 33, 243}
	outputTypeExpect := reflect.TypeOf(HEX{}).Name()

	outputTypeRecieved := reflect.TypeOf(entrySample.Rgb2Hex()).Name()
	if outputTypeRecieved != outputTypeExpect {
		t.Errorf("The data type of the output of %v was incorrect, got: %v, want: %v.", entrySample, outputTypeRecieved, outputTypeExpect)
	}

}

func TestConversion(t *testing.T) {
	entrySample := RGB{11, 33, 243}
	outputExpect := HEX{"0b21f3"}.str

	outputRecieved := entrySample.Rgb2Hex().str
	if outputRecieved != outputExpect {
		t.Errorf("The data type of the output of %v was incorrect, got: %v, want: %v.", entrySample, outputRecieved, outputExpect)
	}
}
