package tests

import (
	assembler "hack_assembler/src/main"
	"os"
	"testing"
)

func TestHasMoreLines(t *testing.T) {
	t.Run("HasMoreLines on empty file should return false", func(t *testing.T) {
		file, err := os.Open("./files/empty.asm")
		assertNotError(t, err)
		parser := assembler.NewParser(file)
		got := parser.HasMoreLines()
		expect := false

		assertEqualBoolean(t, got, expect)
	})

	t.Run("HasMoreLines on only comments file should return false", func(t *testing.T) {
		file, err := os.Open("./files/only-comments.asm")
		assertNotError(t, err)
		parser := assembler.NewParser(file)
		got := parser.HasMoreLines()
		expect := false

		assertEqualBoolean(t, got, expect)
	})

	t.Run("HasMoreLines on file with lines should return true", func(t *testing.T) {
		file, err := os.Open("./files/Add.asm")
		assertNotError(t, err)
		parser := assembler.NewParser(file)
		got := parser.HasMoreLines()
		expect := true

		assertEqualBoolean(t, got, expect)
	})
}
