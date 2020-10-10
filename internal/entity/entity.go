package entity

type Entity struct {
	Room int
}

func NewEntity(room int) *Entity {
	return &Entity{Room: room}
}

func (e *Entity) SetRoom(room int) {
	e.Room = room
}

func (e *Entity) GetRoom() int {
	return e.Room
}
