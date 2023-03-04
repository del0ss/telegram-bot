package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
)

type Excel struct {
	File string
}

func New(file string) *Excel {
	return &Excel{
		File: file,
	}
}

func (e *Excel) OpenFile() (*excelize.File, error) {
	f, err := excelize.OpenFile(e.File)
	if err != nil {
		return nil, err
	}
	return f, nil
}
