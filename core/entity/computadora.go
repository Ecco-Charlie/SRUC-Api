package entity

type Computadora struct {
	Id             uint   `gorm:"primaryKey"`
	NumPatrimonial string `gorm:"size:10;unique"`
	Ip             string `gorm:"size:14"`

	UbicacionId uint
	Ubicacion   Ubicacion

	EstadoId uint
	Estado   Estado
}

type Ubicacion struct {
	Id          uint   `gorm:"primaryKey"`
	Nombre      string `gorm:"size:20"`
	Descripcion string `gorm:"size:40"`
	Capacidad   uint8
}

type Estado struct {
	Id             uint   `gorm:"primaryKey"`
	Nombre         string `gorm:"size:15"`
	Disponibilidad uint8
}

type Clase struct {
	Id uint `gorm:"primaryKey"`

	UbicacionId uint
	Ubicacion   Ubicacion

	Dia    string `gorm:"size:9"`
	Inicio CTime
	Fin    CTime
	Grupo  string `gorm:"size:8"`
}
