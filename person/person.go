package person

type Person struct { // 대문자는 public 소문자는 private
	name string
	age  int
}

func (p Person) SetDetails(name string, age int) {
	p.name = name
	p.age = age
}
