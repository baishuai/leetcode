public class Solution {
    public int countBattleships(char[][] board) {
        int count = 0;
        if (board.length <= 0)
            return count;
        char[] line = board[0];
        for (int i = 0; i < line.length; i++) {
            if (line[i] == 'X' && i == 0) {
                count++;
            } else if (i > 0 && line[i] == 'X' && line[i - 1] != 'X') {
                count++;
            }
        }
        for (int i = 1; i < board.length; i++) {
            line = board[i];

            for (int j = 0; j < line.length; j++) {
                if (line[j] == 'X') {
                    if (j == 0 && board[i - 1][j] != 'X') {
                        count++;
                    } else if (j > 0 && board[i - 1][j] != 'X' && line[j - 1] != 'X') {
                        count++;
                    }
                }
            }
        }
        return count;
    }
}