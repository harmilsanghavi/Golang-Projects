package services

import (
	"hospital-mgt-sys/initializers"
	"hospital-mgt-sys/model"
	"time"
)

func FirstInsertData() {
	var count int64
	initializers.DB.AutoMigrate(&model.Temp{})
	initializers.DB.Model(&model.Temp{}).Count(&count)
	if count > 0 {
		// Data already exists, don't add anything
		return
	}

	// Create some initial records
	users := []model.Temp{
		{Time: parseTime("2023-07-05 10:00:00")},
		{Time: parseTime("2023-07-05 11:00:00")},
		{Time: parseTime("2023-07-05 12:00:00")},
		{Time: parseTime("2023-07-05 13:00:00")},
		{Time: parseTime("2023-07-05 15:00:00")},
		{Time: parseTime("2023-07-05 16:00:00")},
		{Time: parseTime("2023-07-05 17:00:00")},
		{Time: parseTime("2023-07-05 18:30:00")},
		{Time: parseTime("2023-07-05 19:00:00")},
		{Time: parseTime("2023-07-05 19:30:00")},
	}
	initializers.DB.Create(&users)
}
func parseTime(timeStr string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		panic(err)
	}
	return t
}
