package main

// Color represents the colors of rubicscube
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
//    return [...]string{"Red___", "Blue__", "White_", "Orange", "Yellow", "Green_"}[c]
    return [...]string{"R", "B", "W", "O", "Y", "G"}[c]

}