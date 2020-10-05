package bintree

type BinTree struct {
	Value interface{}
	Left  *BinTree
	Right *BinTree
}

func New(value interface{}) *BinTree {
	return &BinTree{
		Value: value,
	}
}

func NewFromSlice(l []interface{}) *BinTree {
	var root *BinTree

	var treeQueue []**BinTree

	treeQueue = append(treeQueue, &root)

	for _, item := range l {
		pti := treeQueue[0]
		treeQueue = treeQueue[1:]
		if item == nil {
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
