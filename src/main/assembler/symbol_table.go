package assembler

import "strconv"

//TODO: traduco KBD, SCREEN R1, R2 etc...

var builtInSymbols = map[string]int{}

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

func (st *SymbolTable) GetAddress(symbol string) int {
	if !st.Contains(symbol) {
		st.AddEntry(symbol)
	}

	register := st.symbols[symbol]
	return register
}
