package model

type Error struct {
	Title       string
	Description string
}

func new(t string, d string) Error {
	return Error{Title: t, Description: d}
}
