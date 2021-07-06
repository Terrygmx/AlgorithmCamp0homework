var (
	m,n int
	visit [][]bool
    dx = [4]int{1,-1,0,0}
	dy = [4]int{0,0,1,-1}
)


func solve(board [][]byte)  {
	m = len(board)
	n = len(board[0])
	visit = make([][]bool, m)
    for i:=range visit {
        visit[i] = make([]bool,n)
    }

	for i:= range board {
			for j := range board[i] {
				// 只有四个边上的O才是dfs起点
				if board[i][j] == 'O' && !visit[i][j] && (j==0 || j==n-1 || i==0 || i==m-1) {
					dfs(board,i,j)
				}
			}
		
	}
	for i:= range board {
		for j:= range board[i] {
			if board[i][j] == 'O' && !visit[i][j] {
				board[i][j] = 'X'
			}
		}
	}
	return
}

func dfs(bd [][]byte, x,y int) {
	visit[x][y]=true
	// 四个方向
	for i:=0;i<4;i++ {
		nx := x+dx[i]
		ny := y+dy[i]
		// 访问数组前，判断合法性
		if nx<0 || ny<0||nx>=m||ny>=n {
			continue
		}
		if bd[nx][ny] == 'O' && !visit[nx][ny] {
			dfs(bd,nx,ny)
		}
	}
}