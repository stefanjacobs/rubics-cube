package main

// Color represents the colors of Rubiks Cube
type Color int

const (
        red Color = iota
        blue
        white
		orange
		yellow
		green
)

func (c Color) String() string {
    return [...]string{"R", "B", "W", "O", "Y", "G"}[c]
}