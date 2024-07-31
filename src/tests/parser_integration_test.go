package tests

import (
	assembler "hack_assembler/src/main/assembler"
	"hack_assembler/src/main/utils"
	"os"
	"testing"
)

func TestAdvance(t *testing.T) {
	t.Run("Reads all instruction types correctly", func(t *testing.T) {
		file, err := os.Open("./files/parser-integration.asm")
		assertNotError(t, err)
		parser := assembler.NewParser(file)

		//Skip over comments and read L instruction
		parser.HasMoreLines()
		err = parser.Advance()
		assertNotError(t, err)
		expectedLInst := "START"
		assertEqualInstruction(t, parser.PreviousInstructionType, "")
		assertEqualInstruction(t, parser.InstructionType, utils.L_INSTRUCTION)
		assertEqualString(t, parser.Symbol, expectedLInst)

		//Go to A instruction
		parser.HasMoreLines()
		err = parser.Advance()
		assertNotError(t, err)
		expectedAInst := "16"
		assertEqualInstruction(t, parser.PreviousInstructionType, utils.L_INSTRUCTION)
		assertEqualInstruction(t, parser.InstructionType, utils.A_INSTRUCTION)
		assertEqualString(t, parser.Symbol, expectedAInst)

		// Read C inst
		parser.HasMoreLines()
		err = parser.Advance()
		assertNotError(t, err)
		expectedDest := "D"
		expectedComp := "A+1"
		assertEqualInstruction(t, parser.PreviousInstructionType, utils.A_INSTRUCTION)
		assertEqualInstruction(t, parser.InstructionType, utils.C_INSTRUCTION)
		assertEqualString(t, parser.Dest, expectedDest)
		assertEqualString(t, parser.Comp, expectedComp)

		//Read A instruction
		parser.HasMoreLines()
		err = parser.Advance()
		assertNotError(t, err)
		expectedAInstJump := "START"
		assertEqualInstruction(t, parser.PreviousInstructionType, utils.C_INSTRUCTION)
		assertEqualInstruction(t, parser.InstructionType, utils.A_INSTRUCTION)
		assertEqualString(t, parser.Symbol, expectedAInstJump)

		// Read C inst JUMP
		parser.HasMoreLines()
		err = parser.Advance()
		assertNotError(t, err)
		expectedDestJump := "D"
		expectedJump := "JGT"
		assertEqualInstruction(t, parser.PreviousInstructionType, utils.A_INSTRUCTION)
		assertEqualInstruction(t, parser.InstructionType, utils.C_INSTRUCTION)
		assertEqualString(t, parser.Dest, expectedDestJump)
		assertEqualString(t, parser.Jump, expectedJump)
	})
}
