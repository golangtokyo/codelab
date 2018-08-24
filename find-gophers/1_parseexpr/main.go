package main

import (
	"go/ast"
	"go/parser"
	"log"
)

func main() {
	// 式をパースする
	expr, err := parser.ParseExpr("v + 1")
	if err != nil {
		log.Fatal("Error:", err)
	}
	// v + 1の抽象構文木(AST)を表示する
	ast.Print(nil, expr)
}
