package bintree

type BinTree struct {
	Value int
	Left  *BinTree
	Right *BinTree
}

func New(value int) *BinTree {
	return &BinTree{
		Value: value,
	}
}

func NewFromSlice(l []int) *BinTree {
	var root *BinTree

	var treeQueue []**BinTree

	treeQueue = append(treeQueue, &root)

	for _, item := range l {
		pti := treeQueue[0]
		treeQueue = treeQueue[1:]
		if item < 0 {
			continue
		}
		if *pti == nil {
			*pti = New(item)
			treeItem := *pti
			//left child
			treeQueue = append(treeQueue, &(treeItem.Left))
			//right child
			treeQueue = append(treeQueue, &(treeItem.Right))
		}
	}

	return root
}
