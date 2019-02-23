package models

type Bot struct {
	ID          string `db:"id"`
	IP          string `db:"ip"`
	WhoAmI      string `db:"whoami"`
	Os          string `db:"os"`
	InstallDate string `db:"install_date"`
	Admin       bool   `db:"admin"`
	Av          string `db:"av"`
	CPU         string `db:"cpu"`
	Gpu         string `db:"gpu"`
	Version     string `db:"version"`
	LastCheckin string `db:"last_checkin"`
	LastCommand string `db:"last_command"`
	NewCommand  string `db:"new_command"`
}
