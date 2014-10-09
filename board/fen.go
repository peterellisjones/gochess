package board

type fenParts struct {
	board          string
	sideToMove     string
	castlingRights string
	epSquare       string
	halfMoveClock  string
	fullMoveNumber string
}

func (parts fenParts) sideToMove() Side {
	return side.Parse(parts.sideToMove)
}

func getFenParts(fen string) fenParts {
	parts := fenParts{
		board:          "",
		sideToMove:     "w",
		castlingRights: "_",
		epSquare:       "_",
		halfMoveClock:  "0",
		fullMoveNumber: "1",
	}

	arr := strings.Split(fen, " ")
	if len(arr) > 0 {
		parts.board = arr[0]
	}

	if len(arr) > 1 {
		parts.sideToMove = arr[1]
	}

	if len(arr) > 2 {
		parts.castlingRights = arr[2]
	}

	if len(arr) > 3 {
		parts.epSquare = arr[3]
	}

	if len(arr) > 4 {
		parts.halfMoveClock = arr[4]
	}

	if len(arr) > 5 {
		parts.fullMoveNumber = arr[5]
	}

	return parts
}
