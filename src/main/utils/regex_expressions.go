package utils

import "regexp"

var IsEmptyLinePattern, _ = regexp.Compile(`^\s*\n$`)
var IsCommentLinePattern, _ = regexp.Compile(`^\s*\/\/\s*.*$`)
var IsAInstruction, _ = regexp.Compile(`^\s*@\s*(?P<symbol>[\wA-Z_$\.]+)\s*(?:\/\/.*)?$`)
var IsCInstruction, _ = regexp.Compile(`^\s*(?P<dest>(D|M|A|DM|AM|AD|ADM))(?:\s*=\s*(?P<comp>[!-]?([DAM]|1|0)\s*(\+|-|\||&)?\s*([DAM]|1|0)?))?(?:\s*;\s*(?P<jump>JGT|JEQ|JGE|JLT|JNE|JLE|JMP))?\s*(?:\/\/.*)?$`)
var IsLInstruction, _ = regexp.Compile(`^\s*\(\s*(?P<symbol>[\wA-Z_$\.]+)\s*\)\s*(?:\/\/.*)?$`)
