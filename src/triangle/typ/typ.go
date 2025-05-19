package typ

func triangleType(nums []int) string {
	if nums[0]+nums[1] <= nums[2] || nums[1]+nums[2] <= nums[0] || nums[0]+nums[2] <= nums[1] {
		return "none"
	}
	if nums[0] == nums[1] {
		if nums[0] == nums[2] {
			return "equilateral"
		}
		return "isosceles"
	} else if nums[0] == nums[2] || nums[1] == nums[2] {
		return "isosceles"
	}
	return "scalene"
}
