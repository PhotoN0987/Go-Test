package main

// インポート文の宣言
import (
	"fmt"
)

// 定数定義
const Pi = 3.14

// メイン関数の宣言
func test() {
	fmt.Println("------メイン------")
	fmt.Println("hello world")
	//関数呼び出し
	variableDeclaration()
	arrayDeclaration()
	mapDeclaration()
	fmt.Println(calc(3))
}

// 変数宣言パターン
func variableDeclaration() {

	fmt.Println("------変数------")

	// 静的型変数宣言
	var num int = 1
	// 動的型推論
	var str = "型は静的に推論"
	// 短縮変数宣言 func内でしか使用できない
	real := 1.23

	fmt.Println(Pi)
	fmt.Println(num)
	fmt.Println(str)
	fmt.Println(real)
}

//　配列,スライス
func arrayDeclaration() {

	fmt.Println("------配列、スライス------")
	// 配列
	var arrays []int = []int{1, 2}
	fmt.Println(arrays)

	// スライス
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice[2])
	fmt.Println(slice[2:4])
	fmt.Println(slice[:2])
	fmt.Println(slice[2:])

	// スライス作成(長さとキャパシティを設定できる)
	sliceM := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(sliceM), cap(sliceM), sliceM)
}

func mapDeclaration() {

	fmt.Println("------マップ------")

	// マップ宣言
	m := map[string]int{"apple": 100, "banana": 500}
	fmt.Println(m)

	// 中身変更
	m["apple"] = 300
	fmt.Println(m)

	// 中身確認
	c, ok := m["apple"]
	fmt.Println(c, ok)

	// 中身追加
	m["orange"] = 400
	fmt.Println(m)
}

// 関数
func calc(inp int) (out int) {

	fmt.Println("------関数------")
	// 戻り値宣言時に変数名を宣言しているため、新規に宣言しなくてよい
	out = inp * inp
	// リターン変数を宣言しているため、戻り変数を指定しなくてよい
	return
}
