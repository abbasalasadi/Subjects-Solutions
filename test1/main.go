package main

import (
    "fmt"
    "piscine"  // Use your actual module path
)

func main() {
    link := &piscine.List{}
    
    for i := 0; i < 10; i++ {
        piscine.Enqueue(link, i)
    }
    
    // Print all items
    current := link.Head
    for current != nil {
        fmt.Println(current.Item)
        current = current.Next
    }
}