package parser



type SymbolTable map[string]*Def

func (s *SymbolTable) Populate() {
	s["defun"] = 1
	s["defvar"] = 2
	s["defoutput"] = 3
}
