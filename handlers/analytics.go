// package handlers

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"microservices/models"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"time"

// 	"gorm.io/gorm"
// )

// // ExportAnalyticsJobFormHandler handles the request to export analytics job form data as CSV
// func ExportAnalyticsJobFormHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	records, err := db.FetchAnalyticsJobFormData(db)
// 	if err != nil {
// 		http.Error(w, "Error fetching data", http.StatusInternalServerError)
// 		return
// 	}

// 	// Write analytics job form data to CSV
// 	filePath, err := writeCSV(records)
// 	if err != nil {
// 		http.Error(w, "Error writing CSV", http.StatusInternalServerError)
// 		return
// 	}

// 	// Respond with the file path where the CSV has been saved
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(fmt.Sprintf("CSV file saved at: %s", filePath)))
// }

// // writeCSV writes the records to a CSV file and saves it in the current directory
// func writeCSV(records []models.AnalyticsJobFormRecord) (string, error) {
// 	fileName := "analytics_job_form_data.csv"
// 	absPath, err := filepath.Abs(fileName)
// 	if err != nil {
// 		return "", err
// 	}

// 	file, err := os.Create(absPath)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer file.Close()

// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	// Write CSV header
// 	if err := writer.Write([]string{"ID", "JobName", "CreatedAt", "FormData"}); err != nil {
// 		return "", err
// 	}

// 	// Write analytics job form data to CSV
// 	for _, record := range records {
// 		row := []string{
// 			fmt.Sprintf("%d", record.ID),
// 			record.JobName,
// 			record.CreatedAt.Format(time.RFC3339),
// 			record.FormData,
// 		}
// 		if err := writer.Write(row); err != nil {
// 			return "", err
// 		}
// 	}

// 	return absPath, nil
// }
