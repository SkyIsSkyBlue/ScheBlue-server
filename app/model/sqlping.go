package model

type SqlPing struct {
	PingId    string `gorm:"type:char(10);not null;primaryKey"`
	PongValue string `gorm:"type:char(10);not null;primaryKey"`
}

func (*SqlPing) TableName() string {
	return "sql_ping_model"
}
