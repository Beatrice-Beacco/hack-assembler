package assembler

import (
	"bufio"
	"os"
)

type Parser struct {
	fileLines []string
	count     int
	Symbol    string
	Dest      string
	Comp      string
	Jump      string
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
		nextLine := p.count + 1
		line := p.fileLines[p.count]
		isEmpty := IsEmptyLinePattern.MatchString(line)
		isComment := IsCommentLinePattern.MatchString(line)
		if !isEmpty && !isComment {
			return true
		}
		p.count = nextLine
	}
	return false
}

func (p *Parser) Advance() {
	hasMoreLines := p.HasMoreLines()

	//If there are no more lines don't advance
	if !hasMoreLines {
		return
	}

	// TODO:
}
