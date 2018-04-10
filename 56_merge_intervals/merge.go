/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func merge(intervals []Interval) []Interval {
    n := len(intervals)
    if n <= 1 {
        return intervals
    }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].Start < intervals[j].Start
    })
    res := make([]Interval, 0)
    cur := intervals[0]
    for i := 1; i < n; i++ {
        if cur.End >= intervals[i].Start {
            if cur.End < intervals[i].End {
                cur.End = intervals[i].End
            }
        } else {
            res = append(res, cur)
            cur = intervals[i]
        }
    }
    res = append(res, cur)
    return res
}