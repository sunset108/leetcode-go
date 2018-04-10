func minSwapsCouples(row []int) int {
    n := len(row)
    var res int
    for i := 0; i < n; i += 2 {
        var couple int
        if row[i] & 1 == 0 {
            couple = row[i] + 1
        } else {
            couple = row[i] - 1
        }
        if row[i + 1] != couple {
            for j := i + 2; j < n; j ++ {
                if row[j] == couple {
                    tmp := row[i + 1]
                    row[i + 1] = row[j]
                    row[j] = tmp
                    res++
                    break
                }
            }
        }
    }
    return res
}