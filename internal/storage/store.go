package storage

import (
	"github.com/360EntSecGroup-Skylar/excelize"
)

type Store interface {
	Write(f *excelize.File, err error) error
	Find(c chan string, userText string)
	Delete()
}
