package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	// ファイルごとのトークンの位置を記録するFileSetを作成する
	fset := token.NewFileSet()

	// ファイル単位で構文解析を行う
	f, err := parser.ParseFile(fset, "_gopher.go", nil, 0)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// 抽象構文木を深さ優先で探索する
	ast.Inspect(f, func(n ast.Node) bool {

		// 識別子ではない場合は無視
		ident, ok := n.(*ast.Ident)
		if !ok {
			return true
		}

		// 識別子がGopherという名前でなければ無視
		if ident.Name != "Gopher" {
			return true
		}

		fmt.Println(fset.Position(ident.Pos()))

		return true
	})
}
