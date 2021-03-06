# updatevrecv

Go Analyzer for detectiong updates in value receiver method

## What's this?
value receiver のメソッドで構造体のフィールドを更新しようとしているコードを検出する

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) SetName(name string) {
	p.Name = name // want "field update in value receiver method"
}
func (p *Person) SetAge(age int) {
	p.Age = age // OK
}


func main() {
	p := Person{
		Name: "foo",
		Age:  24,
	}
	p.SetName("bar")
	fmt.Println(p.Name) // "foo" not modified
    
    p.SetAge(100)
    fmt.Println(p.Age) // 100
 
}

```


## Install

```
$ go get github.com/komeme/updatevrecv/cmd/updatevrecv
```

## Usage

```
$ go vet -vettool=`which updatevrecv` pkgname
```
