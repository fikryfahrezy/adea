package data

import (
	"encoding/json"
	"errors"
	"io"
	"sync"

	"github.com/fikryfahrezy/adea/los-inmen/model"
)

type JsonFile struct {
	path   string
	DbUser map[string]model.User
	DbLoan map[string]model.LoanApplication
	sync.RWMutex
}

func NewJson(path string) *JsonFile {
	return &JsonFile{
		DbUser: make(map[string]model.User),
		DbLoan: make(map[string]model.LoanApplication),
		path:   path,
	}
}

func (f *JsonFile) ScanToMap(r io.Reader, tableName string) error {
	f.Lock()
	defer f.Unlock()

	switch tableName {
	case "user":
		if err := json.NewDecoder(r).Decode(&f.DbUser); err != nil {
			return err
		}
	case "loan":
		if err := json.NewDecoder(r).Decode(&f.DbLoan); err != nil {
			return err
		}
	default:
		return errors.New("table not exist")
	}

	return nil
}

func (f *JsonFile) Generate(w io.Writer) error {
	f.Lock()
	defer f.Unlock()

	res := map[string]interface{}{
		"user": f.DbUser,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		return err
	}

	return nil
}
