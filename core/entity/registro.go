package entity

type Registro struct {
	Id     uint
	Inicio CTime
	Fin    CTime

	ProgramaId string
	Programa   Programa

	ServicioId string
	Servicio   Servicio

	ComputadoraId uint
	Computadora   Computadora

	UsuarioId uint
	Usuario   Usuario
}

type Programa struct {
	Id     string `gorm:"size:20"`
	Nombre string `gorm:"size:25"`
}

type Servicio struct {
	Id     string `gorm:"size:20"`
	Nombre string `gorm:"size:25"`
}
