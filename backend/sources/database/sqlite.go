package database

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLite struct {
	DB *gorm.DB
}

func NewSQLite(dbPath string) (*SQLite, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Auto-migrate tables
	if err := db.AutoMigrate(&Activity{}, &VscodeActivity{}, &Summary{}); err != nil {
		return nil, err
	}
	return &SQLite{DB: db}, nil
}

// SaveActivity saves a general activity
func (s *SQLite) SaveActivity(activity Activity) error {
	return s.DB.Create(&activity).Error
}

// SaveVscodeActivity saves a VS Code-specific activity
func (s *SQLite) SaveVscodeActivity(vscodeActivity VscodeActivity) error {
	return s.DB.Create(&vscodeActivity).Error
}

// GenerateDailySummary generates a daily summary
func (s *SQLite) GenerateDailySummary(date time.Time) (Summary, error) {
	var summary Summary
	var totalTime int64
	s.DB.Model(&Activity{}).Where("date(start_time) = ?", date.Format("2006-01-02")).Select("SUM(duration)").Scan(&totalTime)
	summary = Summary{
		Date:      date,
		TotalTime: totalTime,
		Details:   "Generated summary", // Customize as needed
	}
	return summary, s.DB.Create(&summary).Error
}

// GetActivities retrieves all activities
func (s *SQLite) GetActivities() ([]Activity, error) {
	var activities []Activity
	return activities, s.DB.Find(&activities).Error
}
