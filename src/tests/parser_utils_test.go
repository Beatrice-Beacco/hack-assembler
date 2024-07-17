package tests

import (
	assembler "hack_assembler/src/main"
	"testing"
)

func TestGetInstructionType(t *testing.T) {
	t.Run("Returns numeric A_INSTRUCTION correctly", func(t *testing.T) {
		command := "@16"
		want := assembler.A_INSTRUCTION
		got, err := assembler.GetInstructionType(command)
		assertNotError(t, err)
		assertEqualInstruction(t, got, want)
	})

	t.Run("Returns alfanumeric A_INSTRUCTION correctly", func(t *testing.T) {
		command := "@variable"
		want := assembler.A_INSTRUCTION
		got, err := assembler.GetInstructionType(command)
		assertNotError(t, err)
		assertEqualInstruction(t, got, want)
	})

	t.Run("Returns valid C_INSTRUCTION correctly", func(t *testing.T) {
		command := "0;JMP"
		want := assembler.C_INSTRUCTION
		got, err := assembler.GetInstructionType(command)
		assertNotError(t, err)
		assertEqualInstruction(t, got, want)
	})

	t.Run("Returns error when jump instruction is invalid", func(t *testing.T) {
		command := "0;JJJ"
		_, err := assembler.GetInstructionType(command)
		assertError(t, err)
	})

	t.Run("Returns valid L_INSTRUCTION correctly", func(t *testing.T) {
		command := "(LABEL)"
		want := assembler.L_INSTRUCTION
		got, err := assembler.GetInstructionType(command)
		assertNotError(t, err)
		assertEqualInstruction(t, got, want)
	})

	t.Run("returns erro when invalid instruction", func(t *testing.T) {
		command := "asfghkh"
		_, err := assembler.GetInstructionType(command)
		assertError(t, err)
	})
}

func assertEqualInstruction(t testing.TB, got, want assembler.InstructionType) {
	t.Helper()

	if got != want {
		t.Fatalf("Expected %v, but got %v instead", want, got)
	}
}
