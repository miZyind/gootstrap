package models

import "github.com/gofrs/uuid"

// ObjectType define base object type
type ObjectType int

// ObjectType enum
const (
	Food ObjectType = iota
	Clothing
	Housing
	Transportation
	Education
	Entertainment
)

func (objectType ObjectType) String() string {
	return [...]string{
		"Food",
		"Clothing",
		"Housing",
		"Transportation",
		"Education",
		"Entertainment",
	}[objectType]
}

// Object define base object properties
type Object struct {
	ID     uuid.UUID `json:"id"`
	TypeID int       `json:"typeID"`
	Name   string    `json:"name"`
	Price  string    `json:"price"`
}
