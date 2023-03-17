/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    // two trees are same if all the nodes in each
    // tree hold same value on its data

    // if node of both trees are nil
    if p == nil && q == nil {
        return true
    } else if p == nil && q != nil {
        // when one tree's node is nil while other tree's node is not
        return false
    } else if p != nil && q == nil {
        // this is opposite of the previous case
        return false
    }
    // when node of both trees are not nil, 
    // compare their values and return
    return ((p.Val == q.Val) && 
    (isSameTree(p.Left, q.Left) &&
    isSameTree(p.Right, q.Right)))
    
}
