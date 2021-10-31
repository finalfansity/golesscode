package main

import "fmt"

func main()  {
	a := [][]byte{
		{'X','X','X','X'},
		{'X','O','O','X'},
		{'X','X','O','X'},
		{'X','O','X','X'},
	}
	solve(a)
}

func solve(board [][]byte)  {
	var (
		m = len(board)
		n = len(board[0])
		result = make([][]byte, m)
	)
	if m==n && m==1{
		return
	}
	for i:=0;i<m;i++{
		result[i]=make([]byte, n)
		for j:=0;j<n;j++{
			result[i][j]='X'
		}
	}

	// 从边界上遍历是否有O，有则广搜，将相连的区域变为O
	ds := [][2]int{{0,1},{1,0},{0,-1},{-1,0}}
	initD := 0
	x,y := 0,0
	for !(x==0&&y==0&& initD ==3){
		if board[x][y]=='O' && result[x][y]!='O'{
			//广搜并修改result
			queue := make([][2]int,0)
			queue = append(queue, [2]int{x,y})
			result[x][y]='O'
			search := [][2]int{{1,0},{-1,0},{0,1},{0,-1}}
			for len(queue)>0 {
				cx,cy := queue[0][0],queue[0][1]
				queue=queue[1:]
				for _,s := range search {
					nx,ny := cx+s[0],cy+s[1]
					if nx>=0 && nx<m && ny>=0 && ny<n && board[nx][ny]=='O' && result[nx][ny]!='O'{
						queue = append(queue, [2]int{nx,ny})
						result[nx][ny]='O'
					}
				}
			}
		}
		if (x==0 && y==n-1)||(x==m-1 && y==n-1)||(x==m-1 && y==0){
			initD++
		}
		x += ds[initD][0]
		y += ds[initD][1]
	}

	copy(board, result)
	fmt.Printf("%s\n",board)
}