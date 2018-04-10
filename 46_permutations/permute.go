var res [][]int

func permute(nums []int) [][]int {
    res = make([][]int, 0)
    state := make([]bool, len(nums))
    dfs(nums, state, []int{})
    return res
}

func dfs(nums []int, state []bool, r []int) {
    if len(nums) == len(r) {
        t := make([]int, len(r))
        copy(t, r)
        res = append(res, t)
        return
    }
    for i := 0; i < len(state); i++ {
        if !state[i] {
            state[i] = true
            dfs(nums, state, append(r, nums[i]))
            state[i] = false
        }
    }
}