package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Mbot struct {
	gorm.Model
	Name         string
	Sendid       int
	Secretstring string
}
