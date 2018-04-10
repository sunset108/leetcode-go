/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)
    if root == nil {
        return res
    }
    queue := []*TreeNode{root}
    for len(queue) != 0 {
        n := len(queue)
        r := make([]int, n)
        for i := 0; i < n; i++ {
            r[i] = queue[i].Val
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
            }
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            }
        }
        res = append(res, r)
        queue = queue[n:]
    }
    return res
}