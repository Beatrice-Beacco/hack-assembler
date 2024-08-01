package main

import (
	"bufio"
	"fmt"
	"hack_assembler/src/main/assembler"
	"os"
	"path"
	"regexp"
	"strings"
)

var nameRegex, _ = regexp.Compile(`(?P<name>[A-Za-z_-]+)\.[A-Za-z]{3}$`)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the file path of the .asm file to compile")
	filePath, _ := reader.ReadString('\n')
	filePath = strings.Trim(filePath, " \n\r")
	if filePath == "" {
		fmt.Println("Please provide a file path for a .asm file")
		return
	}

	fmt.Println("outputPath", "", "Enter the file path to save the .hack file")
	outputPath, _ := reader.ReadString('\n')
	outputPath = strings.Trim(outputPath, " \n\r")
	if outputPath == "" {
		fmt.Println("Please provide an output path for the compiled .hack file")
		return
	}

	fileExtension := path.Ext(filePath)
	fmt.Printf("File extension %s", fileExtension)

	if strings.Trim(fileExtension, " \n\r") != ".asm" {
		fmt.Println("Invalid file extension, please provide a .asm")
		return
	}

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("An error occured while opening the file for path %s: %v", filePath, err)
		return
	}

	match := nameRegex.FindStringSubmatch(filePath)
	fileName := match[nameRegex.SubexpIndex("name")]
	outputFileName := outputPath + "/" + fileName + ".hack"

	assembler := assembler.NewAssembler(file)
	err = assembler.AssembleToFile(outputFileName)

	if err != nil {
		fmt.Printf("An error occurred while assembling the file: %v", err)
		return
	}

	fmt.Printf("File assembled to: %s", outputFileName)

}
