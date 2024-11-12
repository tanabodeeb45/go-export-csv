package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	_ "github.com/lib/pq"
	"github.com/xuri/excelize/v2"
)

const (
	DBName   = "cloudviu_dev"
	DBUser   = "cloudviu_dev"
	DBPass   = "localpass"
	DBHost   = "127.0.0.1"
	DBPort   = "5445"
	DBSchema = "public"
)

type Records struct {
	EntryID         int64          `db:"entry_id"`
	EntryDate       sql.NullString `db:"entry_date"`
	RecordID        int64          `db:"record_id"`
	UserID          int64          `db:"user_id"`
	UserMobile      string         `db:"user_mobile"`
	UserName        string         `db:"user_name"`
	OutletID        int64          `db:"outlet_id"`
	OutletName      string         `db:"outlet_name"`
	StoreCode       string         `db:"store_code"`
	CampaignName    string         `db:"campaign_name"`
	FormName        string         `db:"form_name"`
	ItemLabel       string         `db:"item_label"`
	ItemType        string         `db:"item_type"`
	ItemProperty    sql.NullString `db:"item_property"`
	ItemProductName sql.NullString `db:"item_product_name"`
	Value           sql.NullInt64  `db:"value"`
	Content         string         `db:"content"`
	AnalyticsKey    sql.NullString `db:"analytics_key"`
	CreatedAt       sql.NullString `db:"created_at"`
	UpdatedAt       sql.NullString `db:"updated_at"`
}

func main() {
	connStr := createConnectionString()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v\n", err)
	}

	records := fetchData(db)

	if err = writeExcel("analytics_job_form_records.xlsx", records); err != nil {
		log.Fatalf("Failed to write data to Excel: %v\n", err)
	}

	log.Println("Data written to Excel successfully!")
}

func createConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=10",
		DBHost, DBPort, DBUser, DBPass, DBName)
}

func fetchData(db *sql.DB) [][]string {
	var records [][]string
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		startTime := time.Now()

		query := fmt.Sprintf(`
			SELECT 
				ajfr.entry_id, ajfr.entry_date, ajfr.record_id, ajfr.user_id,
				ajfr.user_mobile, ajfr.user_name, ajfr.outlet_id, ajfr.outlet_name,
				s.store_code, ajfr.campaign_name, ajfr.form_name, ajfr.item_label,
				ajfr.item_type, ajfr.item_property, ajfr.item_product_name, 
				ajfr.value, ajfr.content, ajfr.analytics_key, ajfr.created_at, ajfr.updated_at
			FROM %s.analytics_job_form_record ajfr
			LEFT JOIN %s.outlet s ON ajfr.outlet_id = s.id`, DBSchema, DBSchema)

		rows, err := db.Query(query)
		if err != nil {
			log.Fatalf("Failed to fetch data: %v\n", err)
		}
		defer rows.Close()

		for rows.Next() {
			var record Records
			if err := rows.Scan(
				&record.EntryID,
				&record.EntryDate,
				&record.RecordID,
				&record.UserID,
				&record.UserMobile,
				&record.UserName,
				&record.OutletID,
				&record.OutletName,
				&record.StoreCode,
				&record.CampaignName,
				&record.FormName,
				&record.ItemLabel,
				&record.ItemType,
				&record.ItemProperty,
				&record.ItemProductName,
				&record.Value,
				&record.Content,
				&record.AnalyticsKey,
				&record.CreatedAt,
				&record.UpdatedAt,
			); err != nil {
				log.Fatalf("Failed to scan data: %v\n", err)
			}

			mu.Lock()
			records = append(records, convertRecordToSlice(record))
			mu.Unlock()
		}

		if err = rows.Err(); err != nil {
			log.Fatalf("Error fetching data: %v\n", err)
		}

		duration := time.Since(startTime)
		log.Printf("Fetched %d records in %v\n", len(records), duration)
	}()

	wg.Wait()
	return records
}

func convertRecordToSlice(record Records) []string {
	var entryDate, createdAt, updatedAt string

	if record.EntryDate.Valid {
		entryDate = record.EntryDate.String
	}
	if record.CreatedAt.Valid {
		createdAt = formatDate(record.CreatedAt.String)
	}
	if record.UpdatedAt.Valid {
		updatedAt = formatDate(record.UpdatedAt.String)
	}

	return []string{
		fmt.Sprintf("%d", record.EntryID),
		entryDate,
		fmt.Sprintf("%d", record.RecordID),
		fmt.Sprintf("%d", record.UserID),
		record.UserMobile,
		record.UserName,
		fmt.Sprintf("%d", record.OutletID),
		record.OutletName,
		record.StoreCode,
		record.CampaignName,
		record.FormName,
		record.ItemLabel,
		record.ItemType,
		record.ItemProperty.String,
		record.ItemProductName.String,
		fmt.Sprintf("%d", record.Value.Int64),
		record.Content,
		record.AnalyticsKey.String,
		createdAt,
		updatedAt,
	}
}

func formatDate(dateStr string) string {
	if parsedTime, err := time.Parse(time.RFC3339, dateStr); err == nil {
		return parsedTime.Add(7 * time.Hour).Format("02-01-2006 15:04:05")
	}
	return dateStr
}

func writeExcel(filename string, records [][]string) error {
	f := excelize.NewFile()

	headers := []string{
		"entry_id", "entry_date", "record_id", "user_id", "user_mobile",
		"user_name", "outlet_id", "outlet_name", "store_code", "campaign_name",
		"form_name", "item_label", "item_type", "item_property", "item_product_name",
		"value", "content", "analytics_key", "created_at", "updated_at",
	}

	for j, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(j+1, 1)
		if err := f.SetCellValue("Sheet1", cell, header); err != nil {
			return err
		}
	}

	f.SetColWidth("Sheet1", "S", "S", 30)
	f.SetColWidth("Sheet1", "T", "T", 30)

	for i, record := range records {
		for j, value := range record {
			cell, _ := excelize.CoordinatesToCellName(j+1, i+2)

			if j == 0 || j == 2 || j == 3 || j == 6 || j == 15 {
				if numericValue, err := strconv.ParseInt(value, 10, 64); err == nil {
					if err := f.SetCellValue("Sheet1", cell, numericValue); err != nil {
						return err
					}
				}
			} else if j == 18 || j == 19 {
				if formattedTime := formatDate(value); formattedTime != "" {
					if err := f.SetCellValue("Sheet1", cell, formattedTime); err != nil {
						return err
					}
				} else {
					if err := f.SetCellValue("Sheet1", cell, value); err != nil {
						return err
					}
				}
			} else {
				if err := f.SetCellValue("Sheet1", cell, value); err != nil {
					return err
				}
			}
		}
	}

	if err := f.SaveAs(filename); err != nil {
		return err
	}

	return nil
}
