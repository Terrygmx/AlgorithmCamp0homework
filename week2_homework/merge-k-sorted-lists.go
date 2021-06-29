/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    n := len(lists)
    if n== 0 {return nil}
    if n ==1 {return lists[0]}
    mid := n/2
    left := mergeKLists(lists[0:mid])
    right := mergeKLists(lists[mid:])
    return mergeTwoLists(left, right)
}



func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1==nil {return l2}
    if l2==nil {return l1}
    cur := &ListNode{0,nil}
    dummy := cur
    for l1 != nil || l2 != nil {
        if l1 == nil {
            cur.Next = l2
            l2=l2.Next
        }else if l2 == nil {
            cur.Next = l1
            l1=l1.Next
        }else{
            if l1.Val < l2.Val {
                cur.Next = l1
                l1=l1.Next
            }else{
                cur.Next = l2
                l2=l2.Next
            }
        }
        cur=cur.Next
    }
    return dummy.Next

}