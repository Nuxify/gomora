package interfaces

type PlayerServiceInterface interface {
	GetScores(player1Name string, player2Name string) (string, error)
}
