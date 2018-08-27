package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
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

	// 識別子が定義または利用されてる部分を記録する
	defsOrUses := map[*ast.Ident]types.Object{}
	info := &types.Info{
		Defs: defsOrUses,
		Uses: defsOrUses,
	}

	// 型チェックを行うための設定
	config := &types.Config{
		Importer: importer.Default(),
	}

	// 型チェックを行う
	_, err = config.Check("main", fset, []*ast.File{f}, info)
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

		// 識別子が定義または利用されている部分の情報を取得
		obj := defsOrUses[ident]
		if obj == nil {
			return true
		}

		// 型情報を取得し名前付き型でなければ無視
		typ := obj.Type()
		if _, ok := typ.(*types.Named); !ok {
			return true
		}

		fmt.Println(fset.Position(ident.Pos()))

		return true
	})
}
