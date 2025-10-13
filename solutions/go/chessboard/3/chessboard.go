package chessboard

type File []bool
type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
    var count int = 0
	for _, value := range cb[file] {
        if value {
            count++
        }
    }
    return count
}

func CountInRank(cb Chessboard, rank int) int {
    if rank < 1 || rank > 8 {
    	return 0
    }
    var count int = 0
    for _, file := range cb {
        if file[rank - 1] {
            count++
        }
    }
	return count
}

func CountAll(cb Chessboard) int {
	var count int = 0
    for _, file := range cb {
        count += len(file)
    }
	return count
}

func CountOccupied(cb Chessboard) int {
	var count int = 0
    for k, _ := range cb {
        count += CountInFile(cb, k)
    }
	return count
}
