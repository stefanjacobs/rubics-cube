package main

// Color represents the colors of Rubiks Cube
type Color int

const (
    _ Color = iota
    red
    blue
    white
	orange
	yellow
	green
)

func (c Color) String() string {
    return [...]string{"_", "R", "B", "W", "O", "Y", "G"}[c]
}

func (c Color) Int() int {
    return [...]int{0, 1, 2, 3, 4, 5, 6}[c]
}

func (c Color) Byte() byte {
    return [...]byte{0, 1, 2, 3, 4, 5, 6}[c]
}