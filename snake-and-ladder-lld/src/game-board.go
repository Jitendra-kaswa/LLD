package src

type IGameBoard interface {
	AddGamePiece(gp IGamePieces) error
	GetGamePieces() ([]IGamePieces, error)
	GetStartingPoint() (int, error)
	GetEndingPoint() (int, error)
	IsPointOutside(pos int) bool

	GetGamePieceAtStartPoint(pos int) ([]IGamePieces, error)
}

type GameBoard struct {
	startPoint    int
	endPoint      int
	gamePiecesMap map[int][]IGamePieces
	gamePieces    []IGamePieces
}

func NewGameBoard(startPos int, endPos int) IGameBoard {
	return &GameBoard{
		startPoint:    startPos,
		endPoint:      endPos,
		gamePiecesMap: make(map[int][]IGamePieces),
		gamePieces:    make([]IGamePieces, 0),
	}
}

func (gb *GameBoard) AddGamePiece(gp IGamePieces) error {
	gb.gamePieces = append(gb.gamePieces, gp)
	if _, exists := gb.gamePiecesMap[gp.GetStartPosition()]; !exists {
		gb.gamePiecesMap[gp.GetStartPosition()] = make([]IGamePieces, 0)
	}
	gb.gamePiecesMap[gp.GetStartPosition()] = append(gb.gamePiecesMap[gp.GetStartPosition()], gp)
	return nil
}

func (gb *GameBoard) GetGamePieceAtStartPoint(pos int) ([]IGamePieces, error) {
	if gamePieces, exists := gb.gamePiecesMap[pos]; exists {
		return gamePieces, nil
	}
	return make([]IGamePieces, 0), nil
}
func (gb *GameBoard) GetGamePieces() ([]IGamePieces, error) {
	return gb.gamePieces, nil
}

func (gb *GameBoard) GetStartingPoint() (int, error) {
	return gb.startPoint, nil
}
func (gb *GameBoard) GetEndingPoint() (int, error) {
	return gb.endPoint, nil
}
func (gb *GameBoard) IsPointOutside(pos int) bool {
	return pos < gb.startPoint || pos > gb.endPoint
}
