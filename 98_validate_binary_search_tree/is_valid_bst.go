/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    stack := make([]*TreeNode, 0)
    pre, cur := (*TreeNode)(nil), root
    for cur != nil || len(stack) != 0 {
        if cur != nil {
            stack = append(stack, cur)
            cur = cur.Left
        } else {
            cur = stack[len(stack) - 1]
            if pre != nil && pre.Val >= cur.Val {
                return false
            }
            pre = cur
            stack = stack[:len(stack) - 1]
            cur = cur.Right
        }
    }
    return true
}