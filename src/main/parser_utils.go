package assembler

import (
	"fmt"
)

type InstructionType string

const (
	A_INSTRUCTION InstructionType = "A"
	C_INSTRUCTION InstructionType = "C"
	L_INSTRUCTION InstructionType = "L"
)

func GetInstructionType(instruction string) (InstructionType, error) {
	switch {
	case IsAInstruction.MatchString(instruction):
		return A_INSTRUCTION, nil
	case IsCInstruction.MatchString(instruction):
		return C_INSTRUCTION, nil
	case IsLInstruction.MatchString(instruction):
		return L_INSTRUCTION, nil
	default:
		return "", fmt.Errorf("invalid insruction %s", instruction)
	}
}
