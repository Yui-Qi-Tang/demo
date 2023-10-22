package product

import "time"

type Product struct {
	ID        int
	Name      string
	Price     int
	CreatedAt time.Time
}
