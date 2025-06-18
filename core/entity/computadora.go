package entity

type Computadora struct {
	Id             uint   `gorm:"primaryKey"`
	NumPatrimonial string `gorm:"size:10;unique"`
	Ip             string `gorm:"size:14;unique"`

	UbicacionId uint
	Ubicacion   Ubicacion

	EstadoId uint
	Estado   Estado
}

func (Computadora) TableName() string {
	return "computadora"
}

type Ubicacion struct {
	Id          uint   `gorm:"primaryKey"`
	Nombre      string `gorm:"size:20"`
	Descripcion string `gorm:"size:40"`
	Capacidad   uint8
}

func (Ubicacion) TableName() string {
	return "ubicacion"
}

type UbicacionRest struct {
	Id     uint
	Nombre string
}

type Estado struct {
	Id             uint   `gorm:"primaryKey"`
	Nombre         string `gorm:"size:15"`
	Disponibilidad uint8
}

func (Estado) TableName() string {
	return "estado"
}

type ComputadoraDto struct {
	NumPatrimonial string `json:"num_patrimonial"`
	Ubicacion      int32  `json:"ubicacion"`
	Ip             string `json:"ip"`
}
