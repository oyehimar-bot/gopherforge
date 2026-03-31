package main

import (
    "fmt"
    "rsc.io/quote" // This will be your new dependency!
)

func main() {
    fmt.Println("Casting a wisdom spell:")
    fmt.Println(quote.Go()) // Prints a Go-related quote
}