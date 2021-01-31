package main

func removeDuplicates(nums []int) int {
    if len(nums) < 2 {
        return len(nums)
    }
    p1 := 0
    p2 := 1
    for p2 < len(nums) {
        if nums[p1] != nums[p2] {
            p1++
            nums[p1] = nums[p2]
        }
        p2++
    }

    return p1+1
}
