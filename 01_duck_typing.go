package main

import (
	"fmt"
)

// ダックタイピング = とりあえずガァガァ泣くならそれはアヒルでしょ。
// 下記の例は、とりあえず音が鳴るならそれは楽器でしょ、というかんじ。

// ポイントは、
// GuitarやDrumにいちいちInterfaceを明示的に紐付けなくても勝手に紐付けられる
// というところかと思います。

// 多分、これは「そういうもんです」という話で、それ以上でも以下でもない気がする。
// GoのInterfaceは自ずからダックタイピングということになるんだと思う。

type IInstrument interface {
	sound() string
}

type Guitar struct{}

func (g Guitar) sound() string {
	return "ギターは鳴るよ JUGJUG"
}

type Drums struct{}

func (d Drums) sound() string {
	return "ドラムは鳴るよ DAMDAM"
}

func main() {

	instruments := []IInstrument{
		Guitar{},
		Drums{},
	}

	for _, v := range instruments {
		fmt.Println("------------------------------------------")
		fmt.Println(v.sound())
		fmt.Println("------------------------------------------")
	}
}
