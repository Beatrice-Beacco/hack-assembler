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
	t.Run("Returns numeric A_INSTRUCTION with spaces and comment correctly", func(t *testing.T) {
		command := "  @16  //AAAAA"
		want := assembler.A_INSTRUCTION
		got, err := assembler.GetInstructionType(command)
		assertNotError(t, err)
		assertEqualInstruction(t, got, want)
	})
	t.Run("Returns numeric A_INSTRUCTION with spaces between correctly", func(t *testing.T) {
		command := "@  16"
		want := assembler.A_INSTRUCTION
		got, err := assembler.GetInstructionType(command)
		assertNotError(t, err)
		assertEqualInstruction(t, got, want)
	})

	t.Run("Returns valid L_INSTRUCTION correctly", func(t *testing.T) {
		command := "(LABEL)"
		want := assembler.L_INSTRUCTION
		got, err := assembler.GetInstructionType(command)
		assertNotError(t, err)
		assertEqualInstruction(t, got, want)
	})
	t.Run("Returns valid L_INSTRUCTION with spaces and comment correctly", func(t *testing.T) {
		command := "  (LABEL)   //aaaa"
		want := assembler.L_INSTRUCTION
		got, err := assembler.GetInstructionType(command)
		assertNotError(t, err)
		assertEqualInstruction(t, got, want)
	})
	t.Run("Returns valid L_INSTRUCTION with spaces between correctly", func(t *testing.T) {
		command := "(   LABEL  )"
		want := assembler.L_INSTRUCTION
		got, err := assembler.GetInstructionType(command)
		assertNotError(t, err)
		assertEqualInstruction(t, got, want)
	})

	t.Run("Returns valid C_INSTRUCTION correctly", func(t *testing.T) {
		command := "D=D+1;JLE"
		want := assembler.C_INSTRUCTION
		got, err := assembler.GetInstructionType(command)
		assertNotError(t, err)
		assertEqualInstruction(t, got, want)
	})

	t.Run("Returns valid C_INSTRUCTION no comp correctly", func(t *testing.T) {
		command := "D;JLE"
		want := assembler.C_INSTRUCTION
		got, err := assembler.GetInstructionType(command)
		assertNotError(t, err)
		assertEqualInstruction(t, got, want)
	})

	t.Run("Returns valid C_INSTRUCTION no jump correctly", func(t *testing.T) {
		command := "D=D+1"
		want := assembler.C_INSTRUCTION
		got, err := assembler.GetInstructionType(command)
		assertNotError(t, err)
		assertEqualInstruction(t, got, want)
	})

	t.Run("Returns valid C_INSTRUCTION with spaces and comment correctly", func(t *testing.T) {
		command := "  D=D+1;JLE   ///aaaaaaa"
		want := assembler.C_INSTRUCTION
		got, err := assembler.GetInstructionType(command)
		assertNotError(t, err)
		assertEqualInstruction(t, got, want)
	})
	t.Run("Returns valid C_INSTRUCTION with spaces between correctly", func(t *testing.T) {
		command := "D  =  D  +  1  ;  JLE"
		want := assembler.C_INSTRUCTION
		got, err := assembler.GetInstructionType(command)
		assertNotError(t, err)
		assertEqualInstruction(t, got, want)
	})

	t.Run("Returns error when invalid instruction", func(t *testing.T) {
		command := "asfghkh"
		_, err := assembler.GetInstructionType(command)
		assertError(t, err)
	})
}
func TestGetSymbols(t *testing.T) {
	t.Run("Returns A_INSTRUCTION symbol correctly", func(t *testing.T) {
		command := "@16"
		want := "16"
		got, err := assembler.GetAInstructionSymbol(command)
		assertNotError(t, err)
		assertEqualString(t, got, want)
	})
	t.Run("Returns A_INSTRUCTION symbol with spaces and comment correctly", func(t *testing.T) {
		command := " @16  //comment!!!"
		want := "16"
		got, err := assembler.GetAInstructionSymbol(command)
		assertNotError(t, err)
		assertEqualString(t, got, want)
	})
	t.Run("Returns A_INSTRUCTION symbol with spaces between correctly", func(t *testing.T) {
		command := " @ 16 "
		want := "16"
		got, err := assembler.GetAInstructionSymbol(command)
		assertNotError(t, err)
		assertEqualString(t, got, want)
	})
	t.Run("Returns error in invalid A_INSTRUCTION", func(t *testing.T) {
		command := "@@16!!!!"
		_, err := assembler.GetAInstructionSymbol(command)
		assertError(t, err)
	})
	t.Run("Returns L_INSTRUCTION symbol correctly", func(t *testing.T) {
		command := "(TEST)"
		want := "TEST"
		got, err := assembler.GetLInstructionSymbol(command)
		assertNotError(t, err)
		assertEqualString(t, got, want)
	})
	t.Run("Returns L_INSTRUCTION symbol with spaces and comment correctly", func(t *testing.T) {
		command := "  (TEST)  //Comment!!!"
		want := "TEST"
		got, err := assembler.GetLInstructionSymbol(command)
		assertNotError(t, err)
		assertEqualString(t, got, want)
	})
	t.Run("Returns L_INSTRUCTION symbol with spaces between correctly", func(t *testing.T) {
		command := "(  TEST  )"
		want := "TEST"
		got, err := assembler.GetLInstructionSymbol(command)
		assertNotError(t, err)
		assertEqualString(t, got, want)
	})
	t.Run("Returns error in invalid L_INSTRUCTION", func(t *testing.T) {
		command := "(!!''XXX)"
		_, err := assembler.GetLInstructionSymbol(command)
		assertError(t, err)
	})
	t.Run("Returns C_INSTRUCTION symbols correctly", func(t *testing.T) {
		command := "D=D+1;JLE"
		wantDest := "D"
		wantComp := "D+1"
		wantJump := "JLE"
		gotDest, gotComp, gotJump, err := assembler.GetCInstructionSymbols(command)
		assertNotError(t, err)
		assertEqualString(t, gotDest, wantDest)
		assertEqualString(t, gotComp, wantComp)
		assertEqualString(t, gotJump, wantJump)
	})
	t.Run("Returns C_INSTRUCTION symbols no comp correctly", func(t *testing.T) {
		command := "D;JLE"
		wantDest := "D"
		wantComp := ""
		wantJump := "JLE"
		gotDest, gotComp, gotJump, err := assembler.GetCInstructionSymbols(command)
		assertNotError(t, err)
		assertEqualString(t, gotDest, wantDest)
		assertEqualString(t, gotComp, wantComp)
		assertEqualString(t, gotJump, wantJump)
	})
	t.Run("Returns C_INSTRUCTION symbols no jump correctly", func(t *testing.T) {
		command := "D=D+1"
		wantDest := "D"
		wantComp := "D+1"
		wantJump := ""
		gotDest, gotComp, gotJump, err := assembler.GetCInstructionSymbols(command)
		assertNotError(t, err)
		assertEqualString(t, gotDest, wantDest)
		assertEqualString(t, gotComp, wantComp)
		assertEqualString(t, gotJump, wantJump)
	})
	t.Run("Returns C_INSTRUCTION with spaces and symbols correctly", func(t *testing.T) {
		command := "  D=D+1;JLE  //EEEEEE!!!"
		wantDest := "D"
		wantComp := "D+1"
		wantJump := "JLE"
		gotDest, gotComp, gotJump, err := assembler.GetCInstructionSymbols(command)
		assertNotError(t, err)
		assertEqualString(t, gotDest, wantDest)
		assertEqualString(t, gotComp, wantComp)
		assertEqualString(t, gotJump, wantJump)
	})
	t.Run("Returns C_INSTRUCTION with spaces beetween correctly", func(t *testing.T) {
		command := "D = D + 1 ; JLE"
		wantDest := "D"
		wantComp := "D+1"
		wantJump := "JLE"
		gotDest, gotComp, gotJump, err := assembler.GetCInstructionSymbols(command)
		assertNotError(t, err)
		assertEqualString(t, gotDest, wantDest)
		assertEqualString(t, gotComp, wantComp)
		assertEqualString(t, gotJump, wantJump)
	})
	t.Run("Returns error in invalid C_INSTRUCTION", func(t *testing.T) {
		command := "as;e=11+R;JJJ"
		_, _, _, err := assembler.GetCInstructionSymbols(command)
		assertError(t, err)
	})
}

func assertEqualInstruction(t testing.TB, got, want assembler.InstructionType) {
	t.Helper()

	if got != want {
		t.Fatalf("Expected %v, but got %v instead", want, got)
	}
}
