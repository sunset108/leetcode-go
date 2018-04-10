/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{}
    dummy.Next = head
    pre, cur := dummy, head
    for i := 1; i < n; i++ {
        cur = cur.Next
    }
    for cur.Next != nil {
        cur = cur.Next
        pre = pre.Next
    }
    pre.Next = pre.Next.Next
    return dummy.Next
}