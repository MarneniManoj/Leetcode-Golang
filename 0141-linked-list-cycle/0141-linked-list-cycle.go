/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

    m := make(map[*ListNode]bool)
    ok := false

    for ; !ok && head != nil ; _, ok = m[head] {
        m[head] = true
        head = head.Next
    }

    return ok

}