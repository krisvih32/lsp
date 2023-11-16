package main

func compile(code string) (bin string) {
	codelexed := lexer(code)
	codesyntaxanalyzed := syntaxanalyzer(codelexed)
	codesemanticanalyzed := semanticanalyzer(codesyntaxanalyzed)
	codeimcodegen := imcodegen(codesemanticanalyzed)
	bin = codegen(codeimcodegen)
	return
}
