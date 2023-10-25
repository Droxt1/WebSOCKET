package ws

import (
	"context"
	// Assuming this is your WebSocket package
)

type RoomService struct {
	roomRepo RoomRepository
}

func NewRoomService(roomRepo RoomRepository) *RoomService {
	return &RoomService{
		roomRepo: roomRepo,
	}
}

func (s *RoomService) CreateRoom(ctx context.Context, name string) (*Room, error) {
	// Add any business logic or validation here
	// For example, you can check if a room with the same name already exists

	room := &Room{Name: name}
	createdRoom, err := s.roomRepo.CreateRoom(ctx, room)
	if err != nil {
		return nil, err
	}

	return createdRoom, nil
}

// Add other room-related methods here
