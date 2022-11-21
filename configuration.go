package etrisfpocdatamodel

import "gorm.io/gorm"

type Controller struct {
	gorm.Model
	ID          string `gorm:"uniqueIndex;column:id" json:"id"`         // Device ID
	Name        string `gorm:"column:name" json:"name"`                 // Device Name
	Type        string `gorm:"column:type" json:"type"`                 // Device Type
	AgentID     string `gorm:"column:agent_id" json:"agent_id"`         // Controller ID
	ServiceName string `gorm:"column:service_name" json:"service_name"` // Service Name
	// Opts []string
}

type Agent struct {
	gorm.Model
	ID   string `gorm:"uniqueIndex;column:id" json:"id"` // Controller ID
	Name string `gorm:"column:ename" json:"name"`        // Device ID
}

type Service struct {
	gorm.Model
	Name string `gorm:"column:name" json:"name"`
	ID   string `gorm:"column:id" json:"id"`
	Addr string `gorm:"column:addr"`
}
