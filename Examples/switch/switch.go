package main

import "fmt"
import "time"

func main() {
    i := 2

    fmt.Print("Write ", i , " as")
    switch i {
        case 1:
            fmt.Println("one")
        case 2:
            fmt.Println("two")
        case 3:
            fmt.Println("three")
    }

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("its the weekend")
    default: 
        fmt.Println("it is a weekday")
    }

    t := time.Now()
    switch{
    case t.Hour() < 12:
        fmt.Println("It is before noon")
    default:
        fmt.Println("It is after noon")
    }

    whatAmI := func(i interface{}) {
        switch t := i.(type)
    case  bool:
        fmt.Println("I am a bool")
    case int:
        fmt.Println("I am an int")
    default:
        fmt.Printf("I don't know the type %s", t)
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("Sko")
    
}