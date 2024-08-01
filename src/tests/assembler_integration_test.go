package tests

import (
	assembler "hack_assembler/src/main/assembler"
	"os"
	"testing"
)

func TestAssembler(t *testing.T) {

	t.Run("Add.asm converted correctly", func(t *testing.T) {
		file, err := os.Open("./programs/Add.asm")
		assertNotError(t, err)
		testAssembler := assembler.NewAssembler(file)
		err = testAssembler.AssembleToFile("./test_outputs/Add.hack")
		assertNotError(t, err)
		expectedFile, err := os.ReadFile("./expected_outputs/Add.hack")
		assertNotError(t, err)
		outputFile, err := os.ReadFile("./test_outputs/Add.hack")
		assertNotError(t, err)
		assertEqualByteArray(t, outputFile, expectedFile)
	})

	t.Run("Max.asm converted correctly", func(t *testing.T) {
		file, err := os.Open("./programs/Max.asm")
		assertNotError(t, err)
		testAssembler := assembler.NewAssembler(file)
		err = testAssembler.AssembleToFile("./test_outputs/Max.hack")
		assertNotError(t, err)
		expectedFile, err := os.ReadFile("./expected_outputs/Max.hack")
		assertNotError(t, err)
		outputFile, err := os.ReadFile("./test_outputs/Max.hack")
		assertNotError(t, err)
		assertEqualByteArray(t, outputFile, expectedFile)
	})
	t.Run("Pong.asm converted correctly", func(t *testing.T) {
		file, err := os.Open("./programs/Pong.asm")
		assertNotError(t, err)
		testAssembler := assembler.NewAssembler(file)
		err = testAssembler.AssembleToFile("./test_outputs/Pong.hack")
		assertNotError(t, err)
		expectedFile, err := os.ReadFile("./expected_outputs/Pong.hack")
		assertNotError(t, err)
		outputFile, err := os.ReadFile("./test_outputs/Pong.hack")
		assertNotError(t, err)
		assertEqualByteArray(t, outputFile, expectedFile)
	})
	t.Run("Rect.asm converted correctly", func(t *testing.T) {
		file, err := os.Open("./programs/Rect.asm")
		assertNotError(t, err)
		testAssembler := assembler.NewAssembler(file)
		err = testAssembler.AssembleToFile("./test_outputs/Rect.hack")
		assertNotError(t, err)
		expectedFile, err := os.ReadFile("./expected_outputs/Rect.hack")
		assertNotError(t, err)
		outputFile, err := os.ReadFile("./test_outputs/Rect.hack")
		assertNotError(t, err)
		assertEqualByteArray(t, outputFile, expectedFile)
	})
}
