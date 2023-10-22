package store

type Storer interface {
	Get(id any) any
}
