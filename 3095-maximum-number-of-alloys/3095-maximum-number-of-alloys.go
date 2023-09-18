
func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
    // check stock -> metal with high stock should be given preference
    //              -> remaining stocks could be way cheaper in terms of cost - so we cannot greedily say, which metal 
    // should be given priority
    // Similarly cost alone cannot say.
    
    // combination of stock and cost, and requirements only can figure which to use.
    // find how many alloys we can create for each machine given the budget. and return max.
    // for each machine, use stock first if not available buy it.

    // looks like a binary search
    // 
    res := 0
    t := budget
    for i := 0 ; i < n; i++{
        t += stock[i]
    }

    for i := 0 ; i < k; i++{
        // using machine i
        
        res = max(res, maxAlloys(0, t,  composition[i], stock, cost, budget))
       
    }
    return res
}

func max(a int, b int) int{
    if a > b{
        return a;
    }else {
        return b;
    }
}

func maxAlloys(start int, end int, comp []int, stock []int, cost []int, b int) int {

    if(start == end){
        return start
    }
    mid := (start + end + 1)/2
    // fmt.Printf(" start %d, end %d, mid %d", start, end, mid)

    if checkPossible(mid, comp, stock, cost, b){
       return maxAlloys(mid, end,  comp, stock, cost, b)
    } else {
        return maxAlloys(start,mid-1,  comp, stock, cost, b)
    }
}

func checkPossible(count int, comp []int, stock []int, cost []int, b int) bool {
    for i := 0; i < len(comp); i++{
        req := comp[i]*count - stock[i]

        if  req > 0 {
            if b < req * cost[i] {
                return false
            } else {
                b -= req * cost[i];
            }
        }
    }
    return true;
}
    
    