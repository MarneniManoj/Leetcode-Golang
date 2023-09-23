/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

    // maintain 2 pointers
    // one iterates at 1 speed and other in 2 speed.
    // if both collides then a cycle.
    speed := head
    slow := head

    for (speed != nil && speed.Next!= nil) {
        speed = speed.Next.Next
        slow = slow.Next
        if speed == slow {
            return true
        }
    }

    return false

}