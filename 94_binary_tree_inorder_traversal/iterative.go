/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    if root == nil {
        return res
    }
    stack := make([]*TreeNode, 0)
    for root != nil || len(stack) != 0 {
        if root != nil {
            stack = append(stack, root)
            root = root.Left
        } else {
            root = stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            res = append(res, root.Val)
            root = root.Right
        }
    }
    return res
}