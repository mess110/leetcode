def backtrack(board)
  return false if !valid?(board)
  cell_index = first_empty_cell_index(board)
  return true if cell_index < 0
  (1..9).each do |n|
    board[cell_index] = n.to_s
    return true if backtrack(board)
  end
  board[cell_index] = "."
  return false
end

board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
board.flatten!
backtrack(board)