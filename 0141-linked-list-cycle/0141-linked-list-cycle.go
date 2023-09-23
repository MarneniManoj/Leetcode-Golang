/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

    m := make(map[*ListNode]bool)
    for head!=nil {
        fmt.Println(head.Val)
        if _, ok := m[head]; ok {
            return true
        }
        m[head] = true
        head = head.Next
    }
    return false

}