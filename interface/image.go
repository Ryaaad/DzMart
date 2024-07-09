package interfaces

type BaseImage struct {
	PublicID string `gorm:"type:varchar(255);primaryKey;column:public_id"`
}
