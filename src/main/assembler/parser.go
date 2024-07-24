package assembler

import (
	"bufio"
	"fmt"
	utils "hack_assembler/src/main/utils"
	"os"
)

type Parser struct {
	fileLines       []string
	count           int
	InstructionType utils.InstructionType
	Symbol          string
	Dest            string
	Comp            string
	Jump            string
}

func NewParser(file *os.File) Parser {
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return Parser{fileLines: fileLines}
}

func (p *Parser) HasMoreLines() bool {
	// if count >= fileLines.size return false
	if p.count >= len(p.fileLines)-1 {
		return false
	}
	//While count < fileLines.size - 1
	//Set count + 1 as next line index
	//If next line is NOT empty or comment, return true and exit loop
	//If next line is empty or comment, increment count by 1 (set count = nextLine)
	for p.count < len(p.fileLines)-1 {
		nextLineIndex := p.count + 1
		nextLineInstruction := p.fileLines[nextLineIndex]
		isEmpty := utils.IsEmptyLinePattern.MatchString(nextLineInstruction)
		isComment := utils.IsCommentLinePattern.MatchString(nextLineInstruction)
		if !isEmpty && !isComment {
			return true
		}
		p.count = nextLineIndex
	}
	return false
}

func (p *Parser) Advance() error {
	//Fetch next line instruction
	nextLine := p.count + 1
	nextInstruction := p.fileLines[nextLine]

	//Find instruction type
	instructionType, err := utils.GetInstructionType(nextInstruction)

	if err != nil {
		fmt.Printf("Error occurred while decoding instruction: %v", err)
		return err
	}

	p.InstructionType = instructionType

	if p.InstructionType == utils.A_INSTRUCTION {
		symbol, err := utils.GetAInstructionSymbol(nextInstruction)
		if err != nil {
			fmt.Printf("Error occurred while decoding instruction: %v", err)
			return err
		}
		p.Symbol = symbol
		p.resetCInstructionFields() //TODO: valutare se necessario
	}

	if p.InstructionType == utils.L_INSTRUCTION {
		symbol, err := utils.GetLInstructionSymbol(nextInstruction)
		if err != nil {
			fmt.Printf("Error occurred while decoding instruction: %v", err)
			return err
		}
		p.Symbol = symbol
		p.resetCInstructionFields() //TODO: valutare se necessario

	}

	if p.InstructionType == utils.C_INSTRUCTION {
		dest, comp, jump, err := utils.GetCInstructionSymbols(nextInstruction)
		if err != nil {
			fmt.Printf("Error occurred while decoding instruction: %v", err)
			return err
		}
		p.Dest = dest
		p.Comp = comp
		p.Jump = jump
	}

	//Set current line as next line
	p.count = nextLine
	return nil
}

func (p *Parser) resetCInstructionFields() {
	p.Dest = ""
	p.Comp = ""
	p.Jump = ""
}
