package main

type Storer interface {
	Push([]byte) (int, error)
	Fetch(int) ([]byte, error)
}
