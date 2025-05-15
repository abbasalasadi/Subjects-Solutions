package piscine

type Node struct {
    Item int
    Next *Node
}

type List struct {
    Head *Node
    Tail *Node
}