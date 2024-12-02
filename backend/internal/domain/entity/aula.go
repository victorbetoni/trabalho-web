package entity

import "time"

type Aula struct {
	ID           string     `json:"id"`
	Titulo       string     `json:"titulo"`
	Data         time.Time  `json:"data"`
	CargaHoraria int        `json:"carga_horaria"`
	Professor    *Professor `json:"professor"`
}
