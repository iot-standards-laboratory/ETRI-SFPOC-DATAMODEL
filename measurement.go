package etrisfpocdatamodel

import "gorm.io/gorm"

type MeasurementData struct {
	gorm.Model
	ID     string
	Tags   []string
	Status map[string]interface{}
}
