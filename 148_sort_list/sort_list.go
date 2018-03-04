/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    slow, fast, pre := head, head, head
    for fast != nil && fast.Next != nil {
        pre = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    pre.Next = nil
    return merge(sortList(head), sortList(slow))
}

func merge(head1 *ListNode, head2 *ListNode) *ListNode {
    dummy := &ListNode{-1, nil}
    cur := dummy
    for head1 != nil && head2 != nil {
        if head1.Val < head2.Val {
            cur.Next = head1
            head1 = head1.Next
        } else {
            cur.Next = head2
            head2 = head2.Next
        }
        cur = cur.Next
    }
    if head1 != nil {
        cur.Next = head1
    }
    if head2 != nil {
        cur.Next = head2
    }
    return dummy.Next
}