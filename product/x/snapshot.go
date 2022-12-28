package x

type ProductWithSnapshotID struct {
	ProductID      string `gorm:"type:char(36);index:,"`
	ProductVersion int64  `gorm:"type:char(36);index:,"`
}
