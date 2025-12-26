package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person
    EmployeeID int
}

func (e Employee) PrintInfo() {
    fmt.Println(e.Name)
    fmt.Println(e.Age)
    fmt.Println(e.EmployeeID)
}

func main() {
    var name string
    var age int
    var id int

    fmt.Scan(&name)
    fmt.Scan(&age)
    fmt.Scan(&id)

    emp := Employee{
        Person: Person{
            Name: name,
            Age:  age,
        },
        EmployeeID: id,
    }

    emp.PrintInfo()
}
