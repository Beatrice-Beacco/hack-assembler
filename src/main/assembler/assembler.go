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

func (a *Assembler) AssembleToFile(filePath string) error {
	instructions, err := a.instructionsToBinary()
	defer a.parser.ResetParser()

	if err != nil {
		fmt.Println(err)
		return err
	}

	fileContent := strings.Join([]string(instructions), "\r\n")
	fileContent = fileContent + "\r\n"
	byteArray := []byte(fileContent)

	err = os.WriteFile(filePath, byteArray, 0666)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (a *Assembler) instructionsToBinary() ([]string, error) {
	instructions := []string{}

	//First pass: save L instructions symbols in map
	for a.parser.HasMoreLines() {
		err := a.parser.Advance()

		if err != nil {
			return []string{}, err
		}
		if a.parser.InstructionType == utils.L_INSTRUCTION && !a.symbolTable.Contains(a.parser.Symbol) {
			a.symbolTable.AddEntry(a.parser.Symbol, a.parser.instructionCount)
		}
	}
	a.parser.ResetParser() //Reset parser

	//Second pass: Convert instructions and A instruction symbols
	for a.parser.HasMoreLines() {
		err := a.parser.Advance()

		if err != nil {
			return []string{}, err
		}

		if a.parser.InstructionType == utils.C_INSTRUCTION {
			binaryCInstruction, err := a.codeConverter.CInstructionToBinary(a.parser.Comp, a.parser.Dest, a.parser.Jump)
			if err != nil {
				return []string{}, err
			}
			instructions = append(instructions, binaryCInstruction)
		} else if a.parser.InstructionType == utils.A_INSTRUCTION {
			binaryRegister := a.symbolTable.GetBinaryAddress(a.parser.Symbol)
			instructions = append(instructions, binaryRegister)
		}
	}
	return instructions, nil
}
