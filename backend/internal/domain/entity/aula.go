package entity

type Aula struct {
	ID        string     `json:"id"`
	Materia   *Materia   `json:"materia"`
	Professor *Professor `json:"professor"`
}
