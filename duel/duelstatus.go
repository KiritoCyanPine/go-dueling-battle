package duel

// Status defins the position of the dual and the surrent events
type Status string

const (
	// InProgress is the status of the dual
	InProgress Status = "IN_PROGRESS"
	// Player1Wins is the status of the dual
	Player1Wins Status = "PLAYER_1_WINS"
	// Player2Wins is the status of the dual
	Player2Wins Status = "PLAYER_2_WINS"
)
