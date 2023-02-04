package main

import "gorm.io/gorm"

type BaseModel struct {
	Uuid     *string
	Hidden   *bool
	Ordering *int32
}

func (b *BaseModel) Encode() {

}

type User struct {
	gorm.Model
	BaseModel
	Username *string
	Email    *string
	Password *string
}

func (u *User) Encode() {

}
