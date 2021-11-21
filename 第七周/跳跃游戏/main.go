func canJump(nums []int) bool {
	l := len(nums)
	if l < 1 {
		return false
	}
	if l == 1 {
		return true
	}
	// 初始失败，无法跳到下一个
	if nums[0] == 0 {
		return false
	}

	lag := -1
	for i := l - 2; i >= 0; i-- {
		if lag > 0 {
			if i+nums[i] > lag {
				lag = -1
			}
			continue
		}
		if nums[i] == 0 {
			lag = i
			continue
		}
	}
	if lag < 0 {
		return true
	}
	return false
}
