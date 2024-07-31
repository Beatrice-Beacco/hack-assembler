package tests

import (
	assembler "hack_assembler/src/main/assembler"
	"testing"
)

func TestConverter(t *testing.T) {

	converter := assembler.CodeConverter{}

	t.Run("Returns conveted dest correctly", func(t *testing.T) {
		input := "D"
		want := "010"
		got, err := converter.DestToBinary(input)
		assertNotError(t, err)
		assertEqualString(t, got, want)
	})
	t.Run("Returns conveted EMPTY dest correctly", func(t *testing.T) {
		input := ""
		want := "000"
		got, err := converter.DestToBinary(input)
		assertNotError(t, err)
		assertEqualString(t, got, want)
	})
	t.Run("Incorrect dest returns error", func(t *testing.T) {
		input := "DD"
		_, err := converter.DestToBinary(input)
		assertError(t, err)
	})
	t.Run("Returns conveted comp correctly", func(t *testing.T) {
		input := "M"
		want := "110000"
		got, err := converter.CompToBinary(input)
		assertNotError(t, err)
		assertEqualString(t, got, want)
	})
	t.Run("Returns conveted EMPTY comp correctly", func(t *testing.T) {
		input := ""
		want := ""
		got, err := converter.CompToBinary(input)
		assertNotError(t, err)
		assertEqualString(t, got, want)
	})
	t.Run("Incorrect comp returns error", func(t *testing.T) {
		input := "-A-"
		_, err := converter.CompToBinary(input)
		assertError(t, err)
	})
	t.Run("Returns conveted jump correctly", func(t *testing.T) {
		input := "JEQ"
		want := "010"
		got, err := converter.JumpToBinary(input)
		assertNotError(t, err)
		assertEqualString(t, got, want)
	})
	t.Run("Returns conveted EMPTY jump correctly", func(t *testing.T) {
		input := ""
		want := "000"
		got, err := converter.JumpToBinary(input)
		assertNotError(t, err)
		assertEqualString(t, got, want)
	})
	t.Run("Incorrect jump returns error", func(t *testing.T) {
		input := "JJJ"
		_, err := converter.JumpToBinary(input)
		assertError(t, err)
	})

}

// TODO: test 16 bit conversion
func TestCInstructionComposition(t *testing.T) {}
