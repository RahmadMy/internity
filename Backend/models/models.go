package models

import "time"

type User struct {
	ID                  uint         `gorm:"primaryKey;autoIncrement"`
	Name                string       `gorm:"type:varchar(255);not null"`
	Email               string       `gorm:"type:varchar(255);not null;uniqueIndex"`
	Role                string       `gorm:"type:enum('admin','mentor','intern');not null"`
	Avatar              *string      `gorm:"type:varchar(255)"`
	InternshipStartDate *time.Time   `gorm:"type:date"`
	InternshipEndDate   *time.Time   `gorm:"type:date"`
	IsActive            bool         `gorm:"not null;default:true"`
	CreatedAt           time.Time    `gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt           time.Time    `gorm:"type:timestamp;not null;autoUpdateTime"`
	Attendances         []Attendance `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	News                []News       `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
}

type Attendance struct {
	ID                uint       `gorm:"primaryKey;autoIncrement"`
	UserID            uint       `gorm:"not null;index"`
	Date              time.Time  `gorm:"type:date;not null"`
	Status            string     `gorm:"type:enum('hadir','izin','sakit','alfa');not null"`
	ClockIn           *time.Time `gorm:"type:time"`
	ClockOut          *time.Time `gorm:"type:time"`
	LatitudeIn        *float64   `gorm:"type:decimal(10,8)"`
	LongitudeIn       *float64   `gorm:"type:decimal(11,8)"`
	LatitudeOut       *float64   `gorm:"type:decimal(10,8)"`
	LongitudeOut      *float64   `gorm:"type:decimal(11,8)"`
	JobDescription    *string    `gorm:"type:text"`
	DocumentationFile *string    `gorm:"type:varchar(255)"`
	LeaveReason       *string    `gorm:"type:text"`
	MedicalLetterFile *string    `gorm:"type:varchar(255)"`
	CreatedAt         time.Time  `gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt         time.Time  `gorm:"type:timestamp;not null;autoUpdateTime"`
	User              User       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type News struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"type:varchar(255);not null"`
	Content   string    `gorm:"type:text;not null"`
	ImagePath *string   `gorm:"type:varchar(255)"`
	AuthorID  uint      `gorm:"not null;index"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;autoUpdateTime"`
	Author    User      `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
}

type EmailOTP struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Email     string    `gorm:"type:varchar(255);not null;index"`
	OTPCode   string    `gorm:"type:varchar(6);not null"`
	ExpiresAt time.Time `gorm:"type:timestamp;not null"`
	IsUsed    bool      `gorm:"not null;default:false"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;autoCreateTime"`
}

func (EmailOTP) TableName() string {
	return "email_otps"
}
