package student

type Student struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
