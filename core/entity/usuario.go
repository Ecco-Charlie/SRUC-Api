package entity

import "github.com/golang-jwt/jwt/v5"

type Usuario struct {
	NumCuenta    uint    `gorm:"primaryKey"`
	Nombre       string  `gorm:"size:20"`
	ApellPaterno string  `gorm:"size:20"`
	ApellMaterno *string `gorm:"size:20"`
	Rol          string  `gorm:"type:enum('administrativo','alumno','invitado')"`

	Administrativo *Administrativo `gorm:"foreignKey:UsuarioNumCuenta;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Alumno         *Alumno         `gorm:"foreignKey:UsuarioNumCuenta;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Alumno struct {
	UsuarioNumCuenta uint `gorm:"primaryKey"`
	LicenciaturaId   uint
	Licenciatura     Licenciatura
}

type Licenciatura struct {
	Id     uint   `gorm:"primaryKey"`
	Nombre string `gorm:"size:30"`
}

type Administrativo struct {
	UsuarioNumCuenta uint `gorm:"primaryKey"`
	AreaId           uint
	Area             Area
	Acceso           *Acceso `gorm:"foreignKey:UsuarioNumCuenta;references:UsuarioNumCuenta;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Area struct {
	Id     uint   `gorm:"primaryKey"`
	Nombre string `gorm:"size:30"`
}

type Acceso struct {
	UsuarioNumCuenta uint `gorm:"primaryKey"`
	Password         string
}

type LoginDto struct {
	NumCuenta string
	Password  string
}

type UserData struct {
	NumCuenta uint   `json:"numcuenta"`
	Nombre    string `json:"nombre"`
	jwt.RegisteredClaims
}

type UsuarioExtraData struct {
	Administrativo *Administrativo
	Alumno         *Alumno
}
