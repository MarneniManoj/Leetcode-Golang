func sumIndicesWithKSetBits(nums []int, k int) int {
    res := 0
    for i := 0; i < len(nums); i++{
        if k == numBits(i) {
            res += nums[i]
        }
    }

    return res
}

func numBits(i int) int {
    setbits := 0
    
    for k:=0; k < 12; k++ {
        
        if (i >> k & 1) == 1 {
            setbits += 1
        }
    }
    return setbits;
}