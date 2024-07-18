package assembler

import (
	"fmt"
	"strings"
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

func GetAInstructionSymbol(instruction string) (symbol string, err error) {
	if !IsAInstruction.MatchString(instruction) {
		return symbol, fmt.Errorf("invalid A insruction %s", instruction)
	}

	match := IsAInstruction.FindStringSubmatch(instruction)
	return match[IsAInstruction.SubexpIndex("symbol")], nil
}

func GetLInstructionSymbol(instruction string) (synbol string, err error) {
	if !IsLInstruction.MatchString(instruction) {
		return "", fmt.Errorf("invalid L insruction %s", instruction)
	}

	match := IsLInstruction.FindStringSubmatch(instruction)
	return match[IsLInstruction.SubexpIndex("symbol")], nil
}

func GetCInstructionSymbols(instruction string) (dest, comp, jump string, err error) {
	if !IsCInstruction.MatchString(instruction) {
		return "", "", "", fmt.Errorf("invalid C insruction %s", instruction)
	}

	match := IsCInstruction.FindStringSubmatch(instruction)
	dest = match[IsCInstruction.SubexpIndex("dest")]
	comp = strings.ReplaceAll(match[IsCInstruction.SubexpIndex("comp")], " ", "") //Remove all whitespaces from comp match
	jump = match[IsCInstruction.SubexpIndex("jump")]
	return dest, comp, jump, nil
}
