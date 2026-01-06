package dto

import "time"

type GetUsersByUserIDResponse struct {
	UserID       string       `json:"user_id"`
	Status       string       `json:"status"`
	CreatedAt    time.Time    `json:"created_at"`
	PersonalInfo PersonalInfo `json:"personal_info"`
	Preferences  Preferences  `json:"preferences"`
	Addresses    []Address    `json:"addresses"`
}

type PersonalInfo struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	AvatarURL       string `json:"avatar_url"`
	IsEmailVerified bool   `json:"is_email_verified"`
}

type Preferences struct {
	Language       string                  `json:"language"`
	Currency       string                  `json:"currency"`
	MarketingOptIn bool                    `json:"marketing_opt_in"`
	Notifications  NotificationPreferences `json:"notifications"`
}

type NotificationPreferences struct {
	Email bool `json:"email"`
	SMS   bool `json:"sms"`
	Push  bool `json:"push"`
}

type Address struct {
	AddressID         string `json:"address_id"`
	Label             string `json:"label"`
	RecipientName     string `json:"recipient_name"`
	Street            string `json:"street"`
	City              string `json:"city"`
	State             string `json:"state"`
	ZipCode           string `json:"zip_code"`
	Country           string `json:"country"`
	IsDefaultShipping bool   `json:"is_default_shipping"`
	IsDefaultBilling  bool   `json:"is_default_billing"`
	Type              string `json:"type"`
}