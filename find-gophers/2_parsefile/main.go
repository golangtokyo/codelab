package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

const src = `package main
func main() {
	println("Hello, 世界")
}`

func main() {
	// ファイルごとのトークンの位置を記録するFileSetを作成する
	fset := token.NewFileSet()

	// ファイル単位でパースするソースコードはsrcで指定する
	f, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// 抽象構文木(AST)を出力する
	ast.Print(fset, f)
}
