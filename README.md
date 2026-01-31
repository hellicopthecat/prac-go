# Learn Go
## Create Go Project
- go mod init github.com/깃헙유저명/저장소명
- terminal -> ls -> go.mod 생성됨 (go.mod === package.json)
- folder root -> main.go 생성
- cmd + shift + p -> go:choose environment로 환경설정

---
### Variables in Go
var name string = "this is name" // 어떤 타입인지 설정해줘야함 <br/>
name = 12 // already type def string <br/>

var name string = "this is name" // 어떤 타입인지 설정해줘야함 <br/>
name := "hello i am hellicat" // 위와 같음 <br/>
name = hellicat

단 := 는 함수 안에서만 작성할 수 있음 <br/>
함수 밖에서는 var 나 const 로 작성해야함

---
### Types of Go
- bool
- string
- int // 8/16/32/64 -> 양수 음수 가능
- uint // 8/16/32/64 -> 양수만 가능
- byte
- rune
- float32
- complex64

---
### functions
#### Make Functions
```golang
func plus(a int, b int) { // X
	return a + b
} 

func plus(a int, b int) int {
// 무엇을 리턴할지 명시해줘야한다.
// func plus (a,b int) int {} 이것도 가능
	return a + b
}

func main() {
	result := plus(44, 22)
	fmt.Println(result) // 66
}
```

#### Multiful Returns
```golang
func plus(a int, b int, name string) (int, string) {
	return a + b , name
}

func main() {
	result, name := plus(44, 22, "Hellicat")
	fmt.Println(result) // 66
	fmt.Println(name) // Hellicat
}
```

#### Mutiful Arguments
```golang
func plus(a ...int) int {
	total := 0
	for _, items := range a { 
    // "_" 로 사용안함을 선언해줄수 있다.
    // 사용안하는 것을 꼭 "_"로 선언해줘야 컴파일에 성공한다.
    // for index, items := range a { 
    // 첫번째는 index, 두번째는 해당 배열의 아이템들
		total += items
	}
	return total
}

func main() {
	result := plus(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(result)
}
```
---
### FMT Package
```golang
package main

import "fmt"

func main() {
	x := 12409124350
	fmt.Printf("%b\n", x) // %b 바이너리
	fmt.Printf("%o\n", x) // %b 8진법
	fmt.Printf("%x\n", x) // %b 16진법
	fmt.Printf("%U\n", x) // %b 유니코드
	// Sprintf는 단순히 콘솔에 print하는 것이 아니라 format된 string을 return 한다.
	xAsBinary := fmt.Sprintf("%b\n", x)
	fmt.Println(x, " :: ", xAsBinary)
	// import fmt를 누르면 해당 설명 페이지로 안내됨
}
```
---
### Slice and Arrays
Go의 Array에는 항상 정해진 갯수가 있다.
```golang
func main() {
	foods := [3]string{"potato", "pizza", "pasta"}
	for _, food := range foods {
		fmt.Println(food)
	}
	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])
	}
}
```
Array를 쓰고 싶지 않고 계속 커질 수 있는 slice를 쓰고 싶다면 초기에 수를 선언을 안해주면 됨
```golang
func main() {
	foods := []string{"potato", "pizza", "pasta"}
	fmt.Printf("%v\n", foods) // [potato pizza pasta]
	foods = append(foods, "bread") // 배열에 추가
	fmt.Printf("%v\n", foods) // [potato pizza pasta bread]
}
```
### Pointers
Go는 컴퓨터 메모리에 접근할 수 있도록 해준다.<br/>
프로그램을 빠르게 만들 수 있고 data를 mutate하는걸 간단하게 해준다.
```golang
func main() {
	a := 2
	b := a // 여기서 a의 값 2를 복사해 b에 담았기 때문에 12로 a값을 바꿔도 이미 복사를 해두었기에 print를 해도 2가 출력됨
	a = 12
	a = 123124
	a = 4235345
	fmt.Println(b) // 2
}
```
b가 a의 값을 쓰기 위한 방법
```golang
func main() {
	a := 2
	b := &a
	a = 12
	a = 123124
	a = 4235345
	fmt.Println(&a) // a의 주소
	fmt.Println(&b) // a를 저장한 b의 주소
	fmt.Println(*b) // 4235345
}
```
---
### Struct
```golang
package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) sayHello() {
	// 내가 만든 타입구조체에 함수를 넣어주기 위해서는 해당 타입을 연결해주면 된다
	fmt.Printf("hello my name is %s and I'm %d \n", p.name, p.age)
}

func main() {
	choi := person{name: "choi", age: 24}
	park := person{"park", 22} // 이런식으로 타입의 순서만 맞다면 label을 작성 안해도 됨.

	choi.sayHello()

	choi.name = "choi choi"
	fmt.Println(choi)

	choi.sayHello()
	park.sayHello()
}
```

### Struct & Pointers
#### Receiver Function
Java의 Getter 와 Setter를 생각하면 편하다. <br/>
먼저 Setter와 비슷한 경우를 보자면 다음과 같다.
```golang
// /person/person.go
package person

type Person struct { // 대문자는 public 소문자는 private
	name string
	age  int
}

func (p Person) SetDetails(name string, age int) {
	p.name = name
	p.age = age
}
// main.go
package main

import (
	"fmt"

	"github.com/hellicopthecat/learngo/person"
)

func main() {
	choi := person.Person{}
	choi.SetDetails("choi", 30)
	fmt.Println(choi) // { 0}
}
```
```golang
// /person/person.go
func (p Person) SetDetails(name string, age int) {
  ...
}
// main.go
func main() {
  ...
	fmt.Println(choi) // { 0}
}
// SetDetails의 내부 함수에서는 Person이 누군지 알지만
// main.go에서는 두 구조는 완전히 달라 Person이 복사본이 되어 서로 다른 값을 가지게 된다.
```

```golang
// /person/person.go
func (p *Person) SetDetails(name string, age int) {
  // 구조에 포인터를 달아주면 해당 값을 바꿔줄 수 있다.
  ...
}
// main.go
func main() {
  ...
	fmt.Println(choi) // {"choi", 30}
}
```

Getter같은 function은 다음과 같이 만들면 된다.
```golang
func (p Person) Name() string {
  return p.name
}
```