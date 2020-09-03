package a

import "b"

type A struct {
	a1 string
	a2 *string
}

func (a A) SetA1(content string) {
	a.a1 = content // want "field update in value receiver method"
}

func (a A) SetA2(content string) {
	a.a2 = &content // want "field update in value receiver method"
}

func (a *A) SetA1Ptr(content string) {
	a.a1 = content // OK
}

func (a *A) SetA2Ptr(content string) {
	a.a2 = &content // OK
}

func (a A) GetA1() string {
	return a.a1
}

type Hoge struct{}

func (h Hoge) hoge() {
	b.Var = 200 // OK
}

func f() {
	a := A{}
	a.SetA1("hoge")
	a.GetA1()
}

func main() {
	h := Hoge{}
	h.hoge()
}
