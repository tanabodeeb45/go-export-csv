package models

import (
	"time"
)

type AnalyticsJobFormRecord struct {
	ID              uint                    `gorm:"primaryKey"`
	ComputedKey     string                  `gorm:"index:,unique;not null"`
	EntryID         uint                    `gorm:"not null"`
	Entry           JobDeploymentEntry      `gorm:"foreignKey:EntryID;constraint:OnDelete:CASCADE"`
	EntryDate       time.Time               `gorm:"type:date;default:CURRENT_DATE;not null"`
	OutletID        int                     `gorm:"type:int;default:0;not null"`
	Outlet          Outlet                  `gorm:"foreignKey:OutletID;constraint:OnDelete:CASCADE"`
	OutletName      string                  `gorm:"type:varchar(255);default:'';not null"`
	UserID          int                     `gorm:"default:0;not null"`
	UserName        string                  `gorm:"type:varchar(255);default:'';not null"`
	UserMobile      string                  `gorm:"type:varchar(255);default:'';not null"`
	RouteName       string                  `gorm:"type:varchar(255);default:'';not null"`
	RecordID        uint                    `gorm:"not null"`
	Record          JobDeploymentFormRecord `gorm:"foreignKey:RecordID;constraint:OnDelete:CASCADE"`
	CampaignID      uint                    `gorm:"not null"`
	Campaign        Campaign                `gorm:"foreignKey:CampaignID;constraint:OnDelete:CASCADE"`
	FormID          uint                    `gorm:"not null"`
	Form            CampaignForm            `gorm:"foreignKey:FormID;constraint:OnDelete:CASCADE"`
	CampaignName    string                  `gorm:"type:varchar(255)"`
	FormName        string                  `gorm:"type:varchar(255)"`
	ItemUID         string                  `gorm:"not null"`
	ItemType        string                  `gorm:"type:varchar(255)"`
	ItemLabel       string                  `gorm:"type:varchar(255)"`
	ItemProperty    *string                 `gorm:"type:varchar(255);default:null"`
	ItemProductSku  *string                 `gorm:"type:varchar(255);default:null"`
	ItemProductName *string                 `gorm:"type:varchar(255);default:null"`
	AnalyticsKey    *string                 `gorm:"type:varchar(255);default:null"`
	Value           *float64                `gorm:"type:numeric;default:null"`
	Content         string                  `gorm:"type:text;not null"`
	Timestamp       int64                   `gorm:"not null"`
	CreatedAt       time.Time               `gorm:"autoCreateTime"`
	UpdatedAt       time.Time               `gorm:"autoUpdateTime"`
}
