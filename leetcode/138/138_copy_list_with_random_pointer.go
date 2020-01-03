package _138

/*
__author__ = 'lawtech'
__date__ = '2018/8/27 下午7:04'
*/

type RandomListNode struct {
	Label  int
	Next   *RandomListNode
	Random *RandomListNode
}

func Clone(pHead *RandomListNode) *RandomListNode {
	if pHead == nil {
		return pHead
	}

	CloneNodes(pHead)
	ConnectRandomNodes(pHead)
	return ReconnectNodes(pHead)
}

func CloneNodes(pHead *RandomListNode) {
	pNode := pHead
	for pNode != nil {
		pCloned := &RandomListNode{
			Label:  pNode.Label,
			Next:   pNode.Next,
			Random: pNode.Random,
		}
		pNode.Next = pCloned
		pCloned = pCloned.Next
	}
}

func ConnectRandomNodes(pHead *RandomListNode) {
	pNode := pHead
	for pNode != nil {
		pCloned := pNode.Next
		if pNode.Random != nil {
			pCloned.Random = pNode.Random.Next
		}
		pNode = pCloned.Next
	}
}

func ReconnectNodes(pHead *RandomListNode) *RandomListNode {
	pNode := pHead
	pClonedHead := pNode.Next
	pClonedNode := pNode.Next
	pNode.Next = pClonedHead.Next
	pNode = pNode.Next

	for pNode != nil {
		pClonedNode.Next = pNode.Next
		pClonedNode = pClonedNode.Next
		pNode.Next = pClonedNode.Next
		pNode = pNode.Next
	}

	return pClonedHead
}