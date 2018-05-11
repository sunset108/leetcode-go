/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type nodeHeap []*ListNode

func (h nodeHeap) Len() int {
  return len(h)
}

func (h nodeHeap) Swap(i, j int) {
  h[i], h[j] = h[j], h[i]
}

func (h nodeHeap) Less(i, j int) bool {
  return h[i].Val < h[j].Val
}

func (h *nodeHeap) Push(i interface{}) {
  *h = append(*h, i.(*ListNode))
}

func (h *nodeHeap) Pop() interface{} {
  node := (*h)[len(*h)-1]
  *h = (*h)[:len(*h)-1]
  return node
}

func mergeKLists(lists []*ListNode) *ListNode {
  if len(lists) == 0 {
    return nil
  }
  dummy := &ListNode{}
  cur := dummy
  h := &nodeHeap{}
  for i := 0; i < len(lists); i++ {
    if lists[i] != nil {
      *h = append(*h, lists[i])
    }
  }
  heap.Init(h)
  for h.Len() > 0 {
    n := heap.Pop(h).(*ListNode)
    if n.Next != nil {
      heap.Push(h, n.Next)
    }
    cur.Next = n
    cur = cur.Next
  }
  return dummy.Next
}
