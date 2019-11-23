package models

import (
	"github.com/victor-nach/nigerian-senators/api/db"
)

type Senators struct {
	FullName		string
	State			string
	District		string
	PhoneNumber		string
	email			string
}

// db := db.Db()


// GetAll - returns an array containing all the senators
func (s *Senators) GetAll() (*[]Senators, error) {
	senators := []Senators{}
	err := db.Db().Debug.Find(&Senators).Error
}