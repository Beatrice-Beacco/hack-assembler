package assembler

import (
	"fmt"
	"strconv"
)

var builtInSymbols = map[string]string{
	"R0":     "0000000000000000",
	"R1":     "0000000000000001",
	"R2":     "0000000000000010",
	"R3":     "0000000000000011",
	"R4":     "0000000000000100",
	"R5":     "0000000000000101",
	"R6":     "0000000000000110",
	"R7":     "0000000000000111",
	"R8":     "0000000000001000",
	"R9":     "0000000000001001",
	"R10":    "0000000000001010",
	"R11":    "0000000000001011",
	"R12":    "0000000000001100",
	"R13":    "0000000000001101",
	"R14":    "0000000000001110",
	"R15":    "0000000000001111",
	"SP":     "0000000000000000",
	"LCL":    "0000000000000001",
	"ARG":    "0000000000000010",
	"THIS":   "0000000000000011",
	"THAT":   "0000000000000100",
	"SCREEN": "0100000000000000",
	"KBD":    "0110000000000000",
}

type SymbolTable struct {
	nextAvailableAddress int
	symbols              map[string]string
}

func NewSymbolTable() *SymbolTable {
	symbolsMap := make(map[string]string)
	for k, v := range builtInSymbols {
		symbolsMap[k] = v
	}
	return &SymbolTable{nextAvailableAddress: 16, symbols: symbolsMap}
}

func (st *SymbolTable) AddSynbol(symbol string) {
	//If the symbol is a built-in symbol skip add
	_, ok := builtInSymbols[symbol]

	if ok {
		return
	}

	//If the symbol is numeric add entry "number": number (register number)
	intSynbol, err := strconv.Atoi(symbol)
	if err == nil {
		st.symbols[symbol] = intTo16BitBinary(intSynbol)
		return
	}

	//If it's not numeric assign a register
	//Increment the next available address until a free register is found
	for {
		_, nextAddressAlreadyUsed := st.symbols[strconv.Itoa(st.nextAvailableAddress)]
		if !nextAddressAlreadyUsed {
			break
		}
		st.nextAvailableAddress++
	}

	st.symbols[symbol] = intTo16BitBinary(st.nextAvailableAddress) //Set the symbol as the previously found next free register
	st.nextAvailableAddress++                                      //Increment the next free register address by 1 for the next entry
}
func (st *SymbolTable) AddEntry(symbol string, registerIndex int) {
	//If the symbol is a built-in symbol skip add
	_, ok := builtInSymbols[symbol]

	if ok {
		return
	}

	st.symbols[symbol] = intTo16BitBinary(registerIndex)
}

func (st *SymbolTable) Contains(symbol string) bool {
	_, ok := st.symbols[symbol]
	return ok
}

func (st *SymbolTable) GetBinaryAddress(symbol string) string {
	if !st.Contains(symbol) {
		st.AddSynbol(symbol)
	}

	register := st.symbols[symbol]
	return register
}

func intTo16BitBinary(intValue int) string {
	binary := strconv.FormatInt(int64(intValue), 2)
	return fmt.Sprintf("%016s", binary) //Pad with leading 0s until the length is 16 bits
}
