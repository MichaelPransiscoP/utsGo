package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetAllRooms handles the Get All Rooms endpoint
func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connectGORM()
	defer db.Close()

	var rooms []model.Room
	db.Find(&rooms)

	response := models.RoomsResponse{
		Status: "success",
		Data: struct {
			Rooms []models.Room `json:"rooms"`
		}{
			Rooms: rooms,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetDetailRoom handles the Get Detail Room endpoint
func GetDetailRoom(w http.ResponseWriter, r *http.Request) {
	db := connectGORM()
	defer db.Close()

	params := mux.Vars(r)
	roomID, _ := strconv.Atoi(params["id"])

	var room models.Room
	db.First(&room, roomID)

	var participants []models.Participant
	db.Where("id_room = ?", roomID).Find(&participants)

	response := models.RoomDetailResponse{
		Status: "success",
		Data: struct {
			Room         models.Room          `json:"room"`
			Participants []models.Participant `json:"participants"`
		}{
			Room:         room,
			Participants: participants,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// InsertRoom handles the Insert Room endpoint
func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := connectGORM()
	defer db.Close()

	var room models.Room
	json.NewDecoder(r.Body).Decode(&room)

	var game models.Game
	db.First(&game, room.IDGame)

	var participantsCount int64
	db.Model(&models.Participant{}).Where("id_room = ?", room.ID).Count(&participantsCount)

	if participantsCount < int64(game.MaxPlayer) {
		db.Create(&room)
		response := models.InsertRoomResponse{
			Status: "success",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		response := models.InsertRoomResponse{
			Status: "failed",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// LeaveRoom handles the Leave Room endpoint
func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	db := connectGORM()
	defer db.Close()

	params := mux.Vars(r)
	participantID, _ := strconv.Atoi(params["id"])

	var participant models.Participant
	db.First(&participant, participantID)

	if participant.ID != 0 {
		db.Delete(&participant)
		response := models.LeaveRoomResponse{
			Status: "success",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		response := models.LeaveRoomResponse{
			Status: "failed",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
