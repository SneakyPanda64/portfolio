package models

type User struct {
	Country     string `json:"country"`
	LastVisited int64  `json:"last_visited"`
}
