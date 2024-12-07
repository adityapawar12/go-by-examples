package main

import "fmt"

type person struct {
    name string
    age int
}

func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    return &p
}

func main() {
    fmt.Println(person{"Bob", 20})
    fmt.Println(person{name: "Alice", age: 40})
    fmt.Println(person{name: "Fred"})
    fmt.Println(&person{name: "Ann", age: 40})
    fmt.Println(newPerson("Bob"))
    
    s := person{name: "Sean", age: 30}
    fmt.Println(s.name)
    fmt.Println(s.age)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 50
    fmt.Println(sp.age)
    fmt.Println(s.age)

    dog := struct {
        name   string
        isGood bool
    }{
        "Rex",
        true,
    }
    fmt.Println(dog)
}
