package utils

// Paginator struct
type Paginator struct {
	Next     bool
	Previous bool
	Offset   int
}

// Paginate paginates the records
func Paginate(page, limit, count int) Paginator {

	previous := true
	if page == 0 || page == 1 {
		previous = false
	}
	p := Paginator{
		Next:     (page * limit) <= count,
		Previous: previous,
		Offset:   (page - 1) * limit,
	}
	return p
}
