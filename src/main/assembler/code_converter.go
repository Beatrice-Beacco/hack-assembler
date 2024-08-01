package assembler

import (
	"fmt"
	"strings"
)

var compTable = map[string]string{
	"0":   "101010",
	"1":   "111111",
	"-1":  "111010",
	"D":   "001100",
	"A":   "110000",
	"M":   "110000",
	"!D":  "001101",
	"!A":  "110001",
	"!M":  "110001",
	"-D":  "001111",
	"-A":  "110011",
	"-M":  "110011",
	"D+1": "011111",
	"A+1": "110111",
	"M+1": "110111",
	"D-1": "001110",
	"A-1": "110010",
	"M-1": "110010",
	"D+A": "000010",
	"D+M": "000010",
	"D-A": "010011",
	"D-M": "010011",
	"A-D": "000111",
	"M-D": "000111",
	"D&A": "000000",
	"D&M": "000000",
	"D|A": "010101",
	"D|M": "010101",
}
var destTable = map[string]string{
	"M":   "001",
	"D":   "010",
	"DM":  "011",
	"MD":  "011",
	"A":   "100",
	"AM":  "101",
	"MA":  "101",
	"AD":  "110",
	"DA":  "110",
	"ADM": "111",
	"AMD": "111",
	"MAD": "111",
	"MDA": "111",
	"DAM": "111",
	"DMA": "111",
}
var jumpTable = map[string]string{
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

type CodeConverter struct{}

func (c CodeConverter) DestToBinary(dest string) (string, error) {

	if len(dest) == 0 {
		return "000", nil
	}

	binary, ok := destTable[dest]

	if !ok {
		return "", fmt.Errorf("invalid dest %s", dest)
	}

	return binary, nil
}
func (c CodeConverter) CompToBinary(comp string) (string, error) {

	if len(comp) == 0 {
		return "", nil
	}

	binary, ok := compTable[comp]

	if !ok {
		return "", fmt.Errorf("invalid comp %s", comp)
	}

	return binary, nil
}
func (c CodeConverter) JumpToBinary(jump string) (string, error) {

	if len(jump) == 0 {
		return "000", nil
	}

	binary, ok := jumpTable[jump]

	if !ok {
		return "", fmt.Errorf("invalid jump %s", jump)
	}

	return binary, nil
}

func (c CodeConverter) CInstructionToBinary(comp, dest, jump string) (string, error) {

	//TODO: trovo un modo migliore
	if len(comp) == 0 {
		comp, dest = dest, comp
	}

	compBinary, errComp := c.CompToBinary(comp)

	if errComp != nil {
		return "", fmt.Errorf("an error occurred decoding the comp instruction %s: %v", comp, errComp)
	}

	destBinary, errDest := c.DestToBinary(dest)

	if errDest != nil {
		return "", fmt.Errorf("an error occurred decoding the dest instruction %s: %v", dest, errDest)
	}

	jumpBinary, errJump := c.JumpToBinary(jump)

	if errJump != nil {
		return "", fmt.Errorf("an error occurred decoding the jump instruction %s: %v", jump, errJump)
	}

	loadFromABit := canLoadFromABit(comp)
	return fmt.Sprintf("111%s%s%s%s", loadFromABit, compBinary, destBinary, jumpBinary), nil //concatenate to form the 16 bit c instruction
}

func canLoadFromABit(comp string) string {
	if strings.Contains(comp, "M") {
		return "1"
	}

	return "0"
}
