package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
}

// Game represents the games table
type Game struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	MaxPlayer int    `json:"max_player"`
}

// Room represents the rooms table
type Room struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	RoomName string `json:"room_name"`
	IDGame   int    `json:"id_game"`
}

// Participant represents the participants table
type Participant struct {
	ID        int `gorm:"primaryKey" json:"id"`
	IDRoom    int `json:"id_room"`
	IDAccount int `json:"id_account"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// RoomsResponse represents the JSON response for Get All Rooms endpoint
type RoomsResponse struct {
	Status string `json:"status"`
	Data   struct {
		Rooms []Room `json:"rooms"`
	} `json:"data"`
}

// RoomDetailResponse represents the JSON response for Get Detail Rooms endpoint
type RoomDetailResponse struct {
	Status string `json:"status"`
	Data   struct {
		Room Room `json:"room"`
	} `json:"data"`
}

// InsertRoomResponse represents the JSON response for Insert Room endpoint
type InsertRoomResponse struct {
	Status string `json:"status"`
}

// LeaveRoomResponse represents the JSON response for Leave Room endpoint
type LeaveRoomResponse struct {
	Status string `json:"status"`
}
