package squares_of_sorted_array

func SortedSquares(nums []int) []int {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	right := len(nums) - 1
	left := 0
	var largest int
	sq := make([]int, len(nums))

	for i, _ := range nums {
		if abs(nums[left]) < abs(nums[right]) {
			largest = nums[right]
			right--
		} else {
			largest = nums[left]
			left++
		}
		sq[len(nums)-1-i] = largest * largest
	}
	return sq
}
