var res [][]int

func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    res = make([][]int, 0)
    dfs(candidates, 0, target, []int{})
    return res
}

func dfs(candidates []int, index, target int, r []int) {
    if target == 0 {
        t := make([]int, len(r))
        copy(t, r)
        res = append(res, t)
        return
    }
    for i := index; i < len(candidates) && candidates[i] <= target; i++ {
        if i == 0 || candidates[i] > candidates[i - 1] {
            dfs(candidates, i, target - candidates[i], append(r, candidates[i]))
        }
    }
}