var (
	n int
	edges [][]int
	inDeg []int
)

func findOrder(numCourses int, prerequisites [][]int) (ans []int) {
    n = numCourses
	edges = make([][]int,n)
	for i:= range edges{
		edges[i] = make([]int,0)
	}
	inDeg = make([]int,n)
	for _,pre := range prerequisites {
		ai,bi := pre[0],pre[1]
		// 加边模板
		addEdge(bi,ai)
		
	}
	learned :=  topSort()
    if len(learned) == n {
        ans = learned
    }
    return
}


func addEdge(x,y int) {
	edges[x]=append(edges[x],y)
	inDeg[y]++
}
// 返回学过的课程
func topSort() []int{
	learned := []int{}
	// 拓扑排序基于BFS，需要队列
	q := []int{}
	// 从所有零入度点出发
	for i:=0;i<n; i++ {
		if inDeg[i] ==0 {
			q=append(q,i)
		}
	}
	//执行BFS
	for len(q) > 0 {
		x := q[0]	// 取队头（这门课学过了）
		q = q[1:]
		learned = append(learned, x)
		// 考虑x的所有出边
		for _,y:=range edges[x] {
			inDeg[y]-- // 去掉约束关系
			if inDeg[y] ==0 {
				q=append(q,y)
			}
		}
	}
	return learned
	
	
}