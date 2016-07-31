package player

import (
    "fmt"
)

// IsGameOver returns true if the game is over and the name of the winner
func IsGameOver(board [][]string) (bool, string){
    for i := 0; i < 3; i++ {
        if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][0] != ""{
            return true, board[i][0]
        } else if board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[0][i] != ""{
            return true, board[0][i]
        }
    }
    if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != ""{
        return true, board[0][0]
    }
    if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[2][0] != ""{
        return true, board[0][2]
    }
    // search for cat's game
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i][j] == "" {
                return false, ""
            }
        }
    }
    return true, "cat"
}

type Move struct {
    X int
    Y int
}

// getAllMoves returns an array of all possible remaining moves
func getAllMoves(board [][]string) []Move{
    moves := []Move{}

    over, _ := IsGameOver(board)

    if over {
        return moves
    }

    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i][j] == "" {
                move := Move{i,j}
                moves = append(moves, move)
            }
        }
    }

    return moves
}

// doMove returns a new board with the move executed by the player
func doMove(board [][]string, move Move, player string) [][]string{
    newBoard := [][]string{{"","",""},{"","",""},{"","",""}}
    copy(newBoard[0], board[0])
    copy(newBoard[1], board[1])
    copy(newBoard[2], board[2])
    newBoard[move.X][move.Y] = player
    return newBoard
}

func getNextPlayer(player string) string {
    if player == "x" {
        return "o"
    } else {
        return "x"
    }
}

func evaluateBoard(board [][]string, player string, depth int) int {
    val := 0

    over, winner := IsGameOver(board)

    if over && winner == player {
        return 10 + depth
    } else if over && winner == getNextPlayer(player) {
        return -depth - 10
    }

    return val + depth
}

func alphaBetaHelper(board [][]string, alpha int, beta int, player string, depth int) (int, Move) {
    newAlpha := -beta
    newBeta := -alpha

    actionMove := Move{-1,-1}

    over, _ := IsGameOver(board)

    if over || depth == 0{
        return evaluateBoard(board, player, depth), actionMove
    }
    for _, move := range getAllMoves(board) {

        nextBoard := doMove(board, move, player)
        val, _ := alphaBetaHelper(nextBoard, newAlpha, newBeta, getNextPlayer(player), depth - 1)
        val = -val
        if val > newAlpha {
            newAlpha = val
            actionMove = move
            if newAlpha > newBeta {
                return newAlpha, actionMove
            }
        }
    }

    return newAlpha, actionMove
}

func GetNextMove(board [][]string) Move{
    alpha := -10000
    beta := 10000
    score, move := alphaBetaHelper(board, alpha, beta, "o", 10)
    fmt.Println("I got a score of", score, "for", move)
    return move
}