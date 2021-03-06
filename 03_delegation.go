package main

import (
	"fmt"
)

// 委譲
// https://qiita.com/sonatard/items/2b4b70694fd680f6297c

// ポイントは、GuitarやDrumに共通する処理をInstrumentに持たせられること。

// なお、委譲元から委譲先のプロパティにアクセスしたり、その逆を行ったりすることはできない点に注意。
// アクセス可能にする方法もあるけど、アンチパターンらしい。
// 詳しくはようわからん。

type Instrument struct {
}

func (i Instrument) describeSelf() {
	fmt.Println("私は楽器です！")
}

type Guitar struct {
	Instrument
}

type Drum struct {
	Instrument
}

func main() {
	g := Guitar{}
	g.describeSelf()

	d := Drum{}
	d.describeSelf()
}
