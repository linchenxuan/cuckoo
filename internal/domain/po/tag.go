package po

type Tag struct {
	Id   int64  `gorm:"primaryKey;autoIncrement;column:id"`
	Name string `gorm:"type:varchar(20);not null"`
}

func (*Tag) TableName() string {
	return "tag"
}
