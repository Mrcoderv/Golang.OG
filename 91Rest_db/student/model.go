package student

type Student struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:100;not null" json:"name"`
}
