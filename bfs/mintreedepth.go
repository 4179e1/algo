package bfs

import (
	"github.com/4179e1/algo/bintree"
)

func MinBinTreeDepth(root *bintree.BinTree) int {
	if root == nil {
		return 0
	}

	var depth int = 1
	treeQueue := []*bintree.BinTree{root}

	for len(treeQueue) > 0 {

		// 遍历当前层
		size := len(treeQueue)
		for i := 0; i < size; i++ {
			// 每次pop一个
			item := treeQueue[0]
			treeQueue = treeQueue[1:]

			// 终结条件
			if (item.Left == nil) && (item.Right == nil) {
				return depth
			}

			// 如果当前层有子节点，会添加到队列尾部
			// 但只会在下一个for循环中处理
			if item.Left != nil {
				treeQueue = append(treeQueue, item.Left)
			}

			if item.Right != nil {
				treeQueue = append(treeQueue, item.Right)
			}
		}

		// 增加层数
		depth++
	}

	return depth
}
