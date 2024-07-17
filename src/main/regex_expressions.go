package assembler

import "regexp"

var IsEmptyLinePattern, _ = regexp.Compile(`^\s*\n$`)
var IsCommentLinePattern, _ = regexp.Compile(`^\s*\/\/\s*.*$`)
var IsAInstruction, _ = regexp.Compile(`^\s*@[\wA-Z_$]+\s*$`)
var IsCInstruction, _ = regexp.Compile(`^\s*(D|A|0);(JGT|JEQ|JGE|JLT|JNE|JLE|JMP)\s*$`)
var IsLInstruction, _ = regexp.Compile(`^\s*\([\wA-Z_$]+\)\s*$`)
