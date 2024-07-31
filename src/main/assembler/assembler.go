package assembler

import (
	"fmt"
	"hack_assembler/src/main/utils"
	"os"
	"strings"
)

type Assembler struct {
	parser        *Parser
	codeConverter CodeConverter
	symbolTable   *SymbolTable
}

func NewAssembler(file *os.File) *Assembler {
	parser := NewParser(file)
	symbolTable := NewSymbolTable()
	codeConverter := CodeConverter{}
	return &Assembler{parser: parser, codeConverter: codeConverter, symbolTable: symbolTable}
}

func (a *Assembler) AssembleToFile(filePath string) {
	instructions, err := a.instructionsToBinary()
	defer a.parser.ResetParser()

	if err != nil {
		fmt.Println(err)
		return
	}

	fileContent := strings.Join([]string(instructions), "\n")
	byteArray := []byte(fileContent)

	err = os.WriteFile(filePath, byteArray, 0666)

	if err != nil {
		fmt.Println(err)
	}
}

func (a *Assembler) instructionsToBinary() ([]string, error) {
	instructions := []string{}
	for a.parser.HasMoreLines() {
		if a.parser.InstructionType == utils.C_INSTRUCTION {
			loadFromAInstruction := a.parser.PreviousInstructionType == utils.A_INSTRUCTION
			binaryCInstruction, err := a.codeConverter.CInstructionToBinary(a.parser.Comp, a.parser.Dest, a.parser.Jump, loadFromAInstruction)
			if err != nil {
				return []string{}, err
			}
			instructions = append(instructions, binaryCInstruction)
		} else {
			binaryRegister := a.symbolTable.GetBinaryAddress(a.parser.Symbol)
			instructions = append(instructions, binaryRegister)
		}
	}
	return instructions, nil
}
