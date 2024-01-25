package models

type Config struct {
	HoursPerDay int
}

type Downtime struct {
	SLA      string
	Daily    string
	Weekly   string
	Monthly  string
	Quartely string
	Yearly   string
}
