/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var sum int
func convertBST(root *TreeNode) *TreeNode {
    sum = 0
    dfs(root)
    return root
}

func dfs(node *TreeNode) {
    if node != nil {
        // 反序中序遍历该二叉搜索树，记录过程中的节点值之和，即可得到累加树
        dfs(node.Right)
        sum += node.Val
        node.Val = sum
        dfs(node.Left)
    }
}