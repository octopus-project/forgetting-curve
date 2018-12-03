package cards

import (
	"time"

	"github.com/satori/go.uuid"
)

var Buckets = map[int]int{
	1: 1,
	2: 3,
	3: 5,
	4: 10,
	5: 30,
	6: 60,
}

var CardsList = map[uuid.UUID]Card{}

func CalculateNextReview(Step int) time.Time {
	now := time.Now()

	return now.AddDate(0, 0, Buckets[Step] )
}

func Create(Values Card) Card {

	Id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	Values.Id 			= Id
	Values.Step 		= 1
	Values.NextReview	= CalculateNextReview(Values.Step)

	CardsList[Values.Id] = Values

	return Values
}

func All() map[uuid.UUID]Card {
	return CardsList
}

func Exists(Id uuid.UUID) bool {
	_, ok := CardsList[Id]
	
	return ok
}

func Get(Id uuid.UUID) Card {
	return CardsList[Id]
}

func Step(Element Card, Step int) Card {
	if Element.Step == 6 {
		return Element
	} else {
		Element.Step = Step
		Element.NextReview = CalculateNextReview(Element.Step)
		Element = Update(Element.Id, Element)
		return Element
	}
}

func Delete(Id uuid.UUID) bool {
	if Exists(Id) == true {
		delete(CardsList, Id)
		return true
	} else {
		return false
	}
}

func Update(Id uuid.UUID, Entry Card) Card {
	CardsList[Id] = Entry

	return CardsList[Id]
}