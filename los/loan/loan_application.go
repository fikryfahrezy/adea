package loan

import "io"

type FileSaveFunc func(filename string, r io.Reader) (string, error)

type LoanApp struct {
	saveFile   FileSaveFunc
	repository *Repository
}

func NewApp(fileSaveFunc FileSaveFunc, repository *Repository) *LoanApp {
	return &LoanApp{
		saveFile:   fileSaveFunc,
		repository: repository,
	}
}
