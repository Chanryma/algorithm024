package main

func merge(nums1 []int, m int, nums2 []int, n int)  {
    for n > 0 {
        nums1[m+n-1] = nums2[n-1]
        n--
    }
    sort.Ints(nums1)
}
