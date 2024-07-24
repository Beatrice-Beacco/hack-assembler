package utils

import "regexp"

var IsEmptyLinePattern, _ = regexp.Compile(`^\s*\n$`)
var IsCommentLinePattern, _ = regexp.Compile(`^\s*\/\/\s*.*$`)
var IsAInstruction, _ = regexp.Compile(`^\s*@\s*(?P<symbol>[\wA-Z_$\.]+)\s*(?:\/\/.*)?$`)
var IsCInstruction, _ = regexp.Compile(`^\s*(?P<dest>(D|M|A))(?:\s*=\s*(?P<comp>([DAM]|\d+)\s*(\+|-)\s*([DAM]|\d+)))?(?:\s*;\s*(?P<jump>JGT|JEQ|JGE|JLT|JNE|JLE|JMP))?\s*(?:\/\/.*)?$`)
var IsLInstruction, _ = regexp.Compile(`^\s*\(\s*(?P<symbol>[\wA-Z_$\.]+)\s*\)\s*(?:\/\/.*)?$`)
