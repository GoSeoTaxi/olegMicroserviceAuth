package model

import (
	"time"
)

type User struct {
	ID        int64
	UserT     userInfo
	CreatedAt time.Time
	UpdatedAt time.Time
}

type userInfo struct {
	Name         string
	Email        string
	HashPassword string
	Role         int64
}

type Log struct {
	Type string
	time time.Time
	Data string
}
