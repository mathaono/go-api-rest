package models

type Personality struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	History string `json:"history"`
}

// Variável mockada - Não está sendo criada na base de dados
var Personalities []Personality
