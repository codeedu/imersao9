package domain

import(
	"time"
	uuid "github.com/sartori/go.uuid"
)

type CreditCard struct {
	ID string
	Name string
	Number string
	ExpirationMount int32
	ExpirationYear int32
	CVV int32
	Balance float64
	Limit float64
	CreatedAt time.Time
}

func NewCreditCard() *CreditCard {
	c := &CreditCard{}
	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.now()
	return c
}
