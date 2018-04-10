/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    if root == nil {
        return
    }
    cur := root
    for cur != nil {
        tmp := cur.Left
        if tmp != nil {
            for tmp.Right != nil {
                tmp = tmp.Right
            }
            tmp.Right = cur.Right
            cur.Right = cur.Left
            cur.Left = nil
        }
        cur = cur.Right
    }
}