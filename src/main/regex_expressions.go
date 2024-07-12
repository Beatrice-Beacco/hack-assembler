package assembler

import "regexp"

var IsEmptyLinePattern, _ = regexp.Compile(`^\s*\n$`)
var IsCommentLinePattern, _ = regexp.Compile(`^\s*\/\/\s*.*$`)
