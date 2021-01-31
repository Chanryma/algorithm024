package main

func rotate(nums []int, k int)  {
    if len(nums) < 2 {
        return
    }
    for k > 0 {
        tmp := nums[0]
        nums[0] = nums[len(nums)-1]
        i := len(nums)-2
        for i > 0 {
            nums[i+1] = nums[i]
            i--
        }
        nums[1] = tmp
        k--
    }
}
