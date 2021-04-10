package model

type Customer struct {
	FirstName   string               `json:"first_name,omitempty"`
	LastName    string               `json:"last_name,omitempty"`
	Age         uint32               `json:"age,omitempty"`
	Balance     float64              `json:"balance,omitempty"`
	Debt        float64              `json:"debt,omitempty"`
	Preferences *Preferences         `json:"preferences,omitempty"`
	Friends     []*Customer          `json:"friends,omitempty"`
	Addresses   map[string]*Location `json:"addresses,omitempty"`
}
type Preferences struct {
	DarkMode     bool          `json:"dark_mode,omitempty"`
	Language     string        `json:"language,omitempty"`
	Notification *Notification `json:"notification,omitempty"`
}
type Notification struct {
	Sms   bool `json:"sms,omitempty"`
	Email bool `json:"email,omitempty"`
	Push  bool `json:"push,omitempty"`
}
type Location struct {
	Address   string  `json:"address,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	IsActive  bool    `json:"is_active,omitempty"`
}
