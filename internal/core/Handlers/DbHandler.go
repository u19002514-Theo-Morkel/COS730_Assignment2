package Handlers

import "gorm.io/gorm"

type DbHandler struct {
	DB *gorm.DB
}

func NewDB(db *gorm.DB) DbHandler {
	return DbHandler{db}
}
