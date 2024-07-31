package tests

import (
	assembler "hack_assembler/src/main/assembler"
	"testing"
)

func TestRegisterBinary(t *testing.T) {

	t.Run("Built in R0 is returned correctly", func(t *testing.T) {
		symbolTable := assembler.NewSymbolTable()
		want := "0000000000000000"
		got := symbolTable.GetBinaryAddress("R0")
		assertEqualString(t, got, want)
	})

	t.Run("Built in R15 is returned correctly", func(t *testing.T) {
		symbolTable := assembler.NewSymbolTable()
		want := "0000000000001111"
		got := symbolTable.GetBinaryAddress("R15")
		assertEqualString(t, got, want)
	})

	t.Run("Built in SCREEN is returned correctly", func(t *testing.T) {
		symbolTable := assembler.NewSymbolTable()
		want := "0100000000000000"
		got := symbolTable.GetBinaryAddress("SCREEN")
		assertEqualString(t, got, want)
	})

	t.Run("Built in KBD is returned correctly", func(t *testing.T) {
		symbolTable := assembler.NewSymbolTable()
		want := "0110000000000000"
		got := symbolTable.GetBinaryAddress("KBD")
		assertEqualString(t, got, want)
	})

	t.Run("Numeric address is returned correctly", func(t *testing.T) {
		symbolTable := assembler.NewSymbolTable()
		want := "0000000001111000"
		got := symbolTable.GetBinaryAddress("120")
		contains := symbolTable.Contains("120")
		assertEqualBoolean(t, contains, true)
		assertEqualString(t, got, want)
	})
	t.Run("Numeric address is returned correctly 2 times", func(t *testing.T) {
		symbolTable := assembler.NewSymbolTable()
		want := "0000000001111000"
		got1 := symbolTable.GetBinaryAddress("120")
		got2 := symbolTable.GetBinaryAddress("120")
		contains := symbolTable.Contains("120")
		assertEqualBoolean(t, contains, true)
		assertEqualString(t, got1, want)
		assertEqualString(t, got2, want)
	})

	t.Run("First non-numeric address is automatically set as 16", func(t *testing.T) {
		symbolTable := assembler.NewSymbolTable()
		want := "0000000000010000"
		got := symbolTable.GetBinaryAddress("NON-NUMERIC")
		contains := symbolTable.Contains("NON-NUMERIC")
		assertEqualBoolean(t, contains, true)
		assertEqualString(t, got, want)
	})
	t.Run("First non-numeric address is returned correctly 2 times", func(t *testing.T) {
		symbolTable := assembler.NewSymbolTable()
		want := "0000000000010000"
		got1 := symbolTable.GetBinaryAddress("NON-NUMERIC")
		got2 := symbolTable.GetBinaryAddress("NON-NUMERIC")
		contains := symbolTable.Contains("NON-NUMERIC")
		assertEqualBoolean(t, contains, true)
		assertEqualString(t, got1, want)
		assertEqualString(t, got2, want)
	})
	t.Run("First non-numeric address is automatically set as 16 and second is set as 17", func(t *testing.T) {
		symbolTable := assembler.NewSymbolTable()
		want1 := "0000000000010000"
		want2 := "0000000000010001"
		got1 := symbolTable.GetBinaryAddress("NON-NUMERIC1")
		got2 := symbolTable.GetBinaryAddress("NON-NUMERIC2")
		contains1 := symbolTable.Contains("NON-NUMERIC1")
		assertEqualBoolean(t, contains1, true)
		contains2 := symbolTable.Contains("NON-NUMERIC2")
		assertEqualBoolean(t, contains2, true)
		assertEqualString(t, got1, want1)
		assertEqualString(t, got2, want2)
	})

	t.Run("If first non-numeric is 16 and a numeric address is set as 17, the second non numeric address must be 18", func(t *testing.T) {
		symbolTable := assembler.NewSymbolTable()
		wantNonNumeric1 := "0000000000010000"
		wantNumeric := "0000000000010001"
		wantNonNumeric2 := "0000000000010010"
		gotNonNumeric1 := symbolTable.GetBinaryAddress("NON-NUMERIC1")
		gotNumeric := symbolTable.GetBinaryAddress("17")
		gotNonNumeric2 := symbolTable.GetBinaryAddress("NON-NUMERIC2")
		containsNonNumeric1 := symbolTable.Contains("NON-NUMERIC1")
		assertEqualBoolean(t, containsNonNumeric1, true)
		containsNumeric := symbolTable.Contains("17")
		assertEqualBoolean(t, containsNumeric, true)
		containsNonNumeric2 := symbolTable.Contains("NON-NUMERIC2")
		assertEqualBoolean(t, containsNonNumeric2, true)
		assertEqualString(t, gotNonNumeric1, wantNonNumeric1)
		assertEqualString(t, gotNumeric, wantNumeric)
		assertEqualString(t, gotNonNumeric2, wantNonNumeric2)
	})
}
