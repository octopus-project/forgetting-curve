package cards 

import (
	"time"

	"github.com/satori/go.uuid"
)

type Card struct {
	Id 				uuid.UUID 	`json:"id"`
	Name			string		`json:"name"`
	Description		string		`json:"description"`
	Category		string		`json:"category"`
	Owner			uuid.UUID	`json:"owner"`
	Step			int			`json:"step"`
	NextReview		time.Time	`json:"next_review"`
}