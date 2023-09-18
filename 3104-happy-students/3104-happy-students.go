func countWays(nums []int) int {
    // everyone remains happy
    // can we do greedy?
    // ordering does not matter here, so sort in increasing order.
    // then if lowest is zero, then you cannot, not select any one
    // if lowest if +ve, not selecting anyone is good.
    
    // at index i, decide to select him or not.
    // if k, is val, greater than/equal selected count, then you can make him happy only by selecting him.
    // if k, is val, less than selected count, 
    
    // 0,2,3,3,6,6,7,7
    //  should, 
    // should, should, should, should
    // all
    
    
    // iterate and maintain selected, 
        sort.Ints(nums);
    res := 0
    if nums[0] > 0{
        res += 1
    }
    for i:= 1; i < len(nums); i++ {
        if nums[i] > i && nums[i-1] < i{
            res += 1
        }
    }
    if nums[len(nums) - 1] < len(nums) {
        res += 1
    }
    return res;
    
}