package models

import "time"

//Merchant struct
type Merchant struct {
	ID          int        `json:"id" gorm:"AUTO_INCREMENT"`
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	ImageURL    string     `json:"image_url"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

//MerchantData struct
type MerchantData struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	ImageURL    string      `json:"image_url"`
	Description string      `json:"description"`
	CreatedAt   time.Time   `json:"created_at"`
	Outlet      interface{} `json:"outlet,omitempty"`
}

//MerchantSVD struct
type MerchantSVD struct {
	FullName   string            `json:"full_name"`
	BrandName  string            `json:"brand_name"`
	Email      string            `json:"email"`
	Password   string            `json:"password"`
	Phone      string            `json:"phone"`
	CategoryID int               `json:"category_id"`
	Details    MerchantSVDDetail `json:"details"`
	Config     MerchantSVDConfig `json:"config"`
}

//MerchantSVDDetail struct
type MerchantSVDDetail struct {
	Address                string `json:"address"`
	DirectorName           string `json:"director_name"`
	DirectorIdentityNumber string `json:"director_identity_number"`
	DirectorIdentityType   string `json:"director_identity_type"`
}

//MerchantSVDConfig struct
type MerchantSVDConfig struct {
	HasShift  bool        `json:"has_shift"`
	MDRFixfee interface{} `json:"mdr_fixfee"`
}
