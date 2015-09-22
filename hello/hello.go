package main

import (
	"fmt"
	"hill/gello/stringutil"
    "math/cmplx"
)

func add(x, y int) int {
    return x + y
}

func swap(x, y string) (string, string) {
    return y, x
}

// go does not support function overload
// func swap(x, y int) (int, int) {
//     return y, x
// }

// named return values, also called a naked return
// only use in short functions
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}


// if x has no defined type, x will get type from its context
func needInt(x int) int  { return x*10 + 1}
func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    fmt.Printf(stringutil.Reverse("\n!oG ,olleG"))
    fmt.Println(add(42, 13))

    a, b := swap("hello", "world")
    fmt.Println(a, b)
    // go won't function overload, on install you get a mult. defintion of swap
    // and incompatible types on the first def. of swap w. strings
    // a, b := swap(1, 3)
    // fmt.Println(a, b)

    fmt.Println(split(18))

    // declare vars in list with type last
    // default integer value is 0
    var i int
    // default bool value is false
    var c, python, java bool
    fmt.Println(i, c, python, java)
    // default string value is ""

    // or initialize vars, and you can omit the type
    var j = 3
    var c2, python2, java2 = true, false, "no!"
    fmt.Println(j, c2, python2, java2)

    // simpler still, a := syntax can simplify the var with typecasting
    k := 4
    c3, python3, java3 := true, false, "no!"
    fmt.Println(k, c3, python3, java3)

    // variables can be declared as in an import block
    var (
        ToBe    bool        = false
        MaxInt  uint64      = 1<<64 - 1
        z       complex128  = cmplx.Sqrt(-5 + 12i)
    )

    const f = "%T(%v)\n"
    fmt.Printf(f, ToBe, ToBe)
    fmt.Printf(f, MaxInt, MaxInt)
    fmt.Printf(f, z, z)

    // constants are named with the = syntax
    const name = "Hill"
    // and cannot be changed thereafter
    // name = "Bill"

    // numeric constants are high percision values
    // but untyped...
    const (
        Big     = 1 << 100
        Small   = Big >> 99
    )

    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
    // whoops, this results in an integer overflow
    // fmt.Println(needInt(Big))
    fmt.Println(needFloat(Big))

}
