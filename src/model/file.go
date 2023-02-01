package model

import "github.com/goonode/mogo"

type File struct {
	mogo.DocumentModel `bson:",inline" coll:"file"`
	Filename           string
	Storage            int
	Path               string
}
