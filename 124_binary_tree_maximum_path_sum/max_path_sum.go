/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var max int

func maxPathSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    max = ^int((^uint(0) >> 1))
    cal(root)
    return max
}

func cal(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := cal(root.Left)
    right := cal(root.Right)
    if left < 0 {
        left = 0
    }
    if right < 0 {
        right = 0
    }
    if max < root.Val + left + right {
        max = root.Val + left + right
    }
    if left < right {
        return root.Val + right
    } else {
        return root.Val + left
    }
}