package player

type Player struct {
	room int
}

func (p *Player) SetRoom(room int) {
	p.room = room
}
