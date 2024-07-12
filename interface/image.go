package interfaces

type BaseImage struct {
	PublicID string `gorm:"type:varchar(255);primaryKey"`
	Url      string `gorm:"type:varchar(255);column:Url;"`
}
