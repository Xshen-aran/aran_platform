package databases

func Migrate(m ...interface{}) {
	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(m...)
}

// type Model interface {
// 	TableName() string
// }
