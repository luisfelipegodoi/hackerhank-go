package main

// Node struct
type Node struct {
	data  int
	left  *Node
	right *Node
}

func levelOrder(root *Node) [][]int {

	if root == nil {
		return [][]int{}
	}

	queue := []*Node{}
	queue = append(queue, root)
	curNum, nextLevelNum, res, tmp := 1, 0, [][]int{}, []int{}

	for len(queue) != 0 {
		if curNum > 0 {
			node := queue[0]
			if node.left != nil {
				queue = append(queue, node.left)
				nextLevelNum++
			}

			if node.right != nil {
				queue = append(queue, node.right)
				nextLevelNum++
			}

			curNum--
			tmp = append(tmp, node.data)
			queue = queue[1:]
		}

		if curNum == 0 {
			res = append(res, tmp)
			curNum = nextLevelNum
			nextLevelNum = 0
			tmp = []int{}
		}
	}

	return res
}

func main() {

}
