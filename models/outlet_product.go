package models

import (
	"time"
)

type OutletProductStateType uint8

const (
	OutletProductStateInactive OutletProductStateType = iota + 1
	OutletProductStateActive
)

type OutletProduct struct {
	ID         int64
	OutletID   int
	ProductID  int
	Price      int
	OrderPrice int
	Product    Product
	Outlet     Outlet
	State      OutletProductStateType
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (s OutletProductStateType) ToString() string {
	types := map[OutletProductStateType]string{
		OutletProductStateActive:   "active",
		OutletProductStateInactive: "inactive",
	}
	return types[s]
}

func ToOutletProductStateType(str string) OutletProductStateType {
	var outletProduct OutletProductStateType

	switch str {
	case "active":
		outletProduct = OutletProductStateActive
	case "inactive":
		outletProduct = OutletProductStateInactive
	}
	return outletProduct
}
