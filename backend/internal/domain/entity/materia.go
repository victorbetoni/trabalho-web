package entity

type Materia struct {
	ID           string `json:"id"`
	Nome         string `json:"nome"`
	CargaHoraria int    `json:"carga_horaria"`
	Professor    string `json:"professor"`
}
