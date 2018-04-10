func reverse(x int) int {
    var maxInt int = int(^uint32(0) >> 1)
    var res int64
    for x != 0 {
        res = res * 10 + int64(x) % 10
        x /= 10
    }
    if (res > 0 && res > int64(maxInt)) || (res < 0 && res < int64(^maxInt)) {
        return 0
    }
    return int(res)
}