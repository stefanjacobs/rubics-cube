package main

// Color represents the colors of Rubiks Cube
type Color int

const (
    red Color = iota // 0 
    blue
    white
	orange
	yellow
	green
)

func (c Color) String() string {
    return [...]string{"R", "B", "W", "O", "Y", "G"}[c]
}

func (c Color) Int() int {
    return [...]int{0, 1, 2, 3, 4, 5}[c]
}