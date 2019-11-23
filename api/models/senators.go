package models

import (
	"log"
	// "github.com/victor-nach/nigerian-senators/api/db"
)

type Senator struct {
	FullName		string
	State			string
	District		string
	PhoneNumber		string
	email			string
}

// var Senators []Senators
// db := db.Db()


// GetAll - returns an array containing all the senators
func (s Senator) GetAll() ([]Senator, error) {
	Senators := []Senator{}
	err := Db().Model(&Senator{}).Find(&Senators).Error
	if err != nil {
		log.Println(err)
	}
	log.Println(Senators)
	return Senators, nil
}
