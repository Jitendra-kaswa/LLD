package src

type IGamePlayer interface {
	GetUserId() (int, error)
	GetPlayerId() (int, error)
	SetNewPosition(newPosition int) error
	GetCurrentPosition() (int, error)
}

type GamePlayer struct {
	id              int
	playerId        int
	currentPosition int
}

func NewGamePlayer(id, playerId int) IGamePlayer {
	return &GamePlayer{
		id:              id,
		playerId:        playerId,
		currentPosition: 0,
	}
}

func (gp *GamePlayer) GetUserId() (int, error) {
	if gp == nil {
		return 0, ErrUserNotExists
	}
	return gp.id, nil
}

func (gp *GamePlayer) GetPlayerId() (int, error) {
	if gp == nil {
		return 0, ErrUserNotExists
	}
	return gp.playerId, nil
}

func (gp *GamePlayer) SetNewPosition(newPosition int) error {
	if gp == nil {
		return ErrUserNotExists
	}
	gp.currentPosition = newPosition
	return nil
}

func (gp *GamePlayer) GetCurrentPosition() (int, error) {
	if gp == nil {
		return 0, ErrUserNotExists
	}
	return gp.currentPosition, nil
}
