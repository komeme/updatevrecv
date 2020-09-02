package a

type A struct {
	a1 string
}

func (a A) SetA1(content string) {
	a.a1 = content // want "NG"
}

func (a *A) SetA1Valid(content string) {
	a.a1 = content
}

func (a A) GetA1() string{
	return a.a1
}

func f() {
	a := A{}
	a.SetA1("hoge")
	a.GetA1()
}

