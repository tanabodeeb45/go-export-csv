// package db

// import (
// 	"microservices/models"

// 	"gorm.io/gorm"
// )

// func FetchAnalyticsJobFormData(db *gorm.DB) ([]models.AnalyticsJobFormRecord, error) {
// 	var records []models.AnalyticsJobFormRecord
// 	if err := db.Find(&records).Error; err != nil {
// 		return nil, err
// 	}
// 	return records, nil
// }
