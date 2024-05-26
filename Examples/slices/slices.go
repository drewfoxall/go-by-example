package main

import "fmt"

func main(){
    var s []string
    fmt.Println("uninit", s, s ==nil, len(s)==0)

    s = make([]string, 3)
    fmt.Println("emp: ", s, "len: ", len(s), "cap:" , cap(s))

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"

    fmt.Println(s)
    fmt.Println(len(s))

    s = append(s, "d")
    fmt.Println(s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println(c)

    l := s[2:5]
    fmt.Println(l)

    l = s[2:]
    fmt.Println(l)

    l = s[:5]
    fmt.Println(l)

    x := make([]string, len(s-2))
    copy(x,s)
    fmt.Println(x)


    t := [string]{"d","h","j"}
    t2 := [string]{"a","b","i"}
    if slices.Equal(t, t2){
        fmt.Println("They are the same")
    }


}