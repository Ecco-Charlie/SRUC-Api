package entity

type Usuario struct {
	NumCuenta    uint    `gorm:"primaryKey"`
	Nombre       string  `gorm:"size:20"`
	ApellPaterno string  `gorm:"size:20"`
	ApellMaterno *string `gorm:"size:20"`
	Rol          string  `gorm:"type:enum('administrativo','alumno','invitado')"`
}

type Alumno struct {
	UsuarioNumCuenta uint
	Usuario          Usuario `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Licenciatura     string  `gorm:"size:30"`
}

type Administrativo struct {
	UsuarioNumCuenta uint
	Usuario          Usuario `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Area             string  `gorm:"size:30"`
}

type Acceso struct {
	UsuarioNumCuenta uint    `gorm:"primaryKey"`
	Usuario          Usuario `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	Password         string
}

type LoginDto struct {
	NumCuenta string
	Password  string
}
