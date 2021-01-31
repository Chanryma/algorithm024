package main

func twoSum(nums []int, target int) []int {
    i := 0
    for i < len(nums) - 1 {
        j := i + 1
        for j < len(nums) {
            if nums[i] + nums[j] == target {
                return []int {i, j}
            }
            j++
        }
        i++
    }
    return nil
}

func twoSum1(nums []int, target int) []int {
    m := make(map[int]int, len(nums))
    for i, n := range nums {
        m[n] = i
    }
    for i, n := range nums {
        t := target - n
        ti, found := m[t]
        if found && ti != i {
            return []int {i, ti}
        }
    }
    return nil
}
