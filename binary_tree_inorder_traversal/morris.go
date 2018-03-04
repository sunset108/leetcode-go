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
    var pre, cur *TreeNode = nil, root
    for cur != nil {
        pre = cur.Left
        if pre != nil {
            for pre.Right != nil && pre.Right != cur {
                pre = pre.Right
            }
            if pre.Right == nil {
                pre.Right = cur
                cur = cur.Left
            } else {
                pre.Right = nil
                res = append(res, cur.Val)
                cur = cur.Right
            } 
        } else {
            res = append(res, cur.Val)
            cur = cur.Right
        }
    }
    return res
}