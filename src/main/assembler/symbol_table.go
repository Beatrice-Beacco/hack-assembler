package assembler

import (
	"fmt"
	"strconv"
)

//TODO: traduco KBD, SCREEN R1, R2 etc...

var builtInSymbols = map[string]int{
	"R0":     0,
	"R1":     1,
	"R2":     2,
	"R3":     3,
	"R4":     4,
	"R5":     5,
	"R6":     6,
	"R7":     7,
	"R8":     8,
	"R9":     9,
	"R10":    10,
	"R11":    11,
	"R12":    12,
	"R13":    13,
	"R14":    14,
	"R15":    15,
	"SP":     0,
	"LCL":    1,
	"ARG":    2,
	"THIS":   3,
	"THAT":   4,
	"SCREEN": 16384,
	"KBD":    24576,
}

type SymbolTable struct {
	nextAvailableAddress int
	symbols              map[string]int
}

func NewSymbolTable() SymbolTable {
	return SymbolTable{nextAvailableAddress: 16, symbols: builtInSymbols}
}

func (st *SymbolTable) AddEntry(symbol string) {
	//If the symbol is a built-in symbol skip add
	_, ok := builtInSymbols[symbol]

	if ok {
		return
	}

	//If the symbol is numeric add entry "number": number (register number)
	intSynbol, err := strconv.Atoi(symbol)
	if err == nil {
		st.symbols[symbol] = intSynbol
		return
	}

	//If it's not numeric assign a register
	//Increment the next available address until a free register is found
	for _, ok := st.symbols[strconv.Itoa(st.nextAvailableAddress)]; ok; {
		st.nextAvailableAddress++
	}

	st.symbols[symbol] = st.nextAvailableAddress //Set the symbol as the previously found next free register
	st.nextAvailableAddress++                    //Increment the next free register address by 1 for the next entry
}

func (st *SymbolTable) Contains(symbol string) bool {
	_, ok := st.symbols[symbol]
	return ok
}

func (st *SymbolTable) GetBinaryAddress(symbol string) string {
	if !st.Contains(symbol) {
		st.AddEntry(symbol)
	}

	register := st.symbols[symbol]
	binary := strconv.FormatInt(int64(register), 2)
	return fmt.Sprintf("%016s", binary) //Pad with leading 0s until the length is 16 bits
}
