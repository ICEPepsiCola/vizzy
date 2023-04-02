package entity

type Page5 struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Area     string `json:"area"`
	Count    int    `json:"count"`
	Percent  string `json:"percent"`
	Duration string `json:"duration"`
}
