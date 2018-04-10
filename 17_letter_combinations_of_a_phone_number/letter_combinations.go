var res []string

func letterCombinations(digits string) []string {
    res = make([]string, 0)
    if len(digits) == 0 {
        return res
    }
    dfs(digits, []byte{})
    return res
}

func dfs(d string, c []byte) {
    if len(d) == 0 {
        res = append(res, string(c))
        return
    }
    l := toLetter(d[0])
    for i := byte(0); ((l == 'p' || l == 'w') && i < 4) || (l != 'p' && l != 'w' && i < 3); i++ {
        dfs(d[1:], append(c, l + i))
    }
}

func toLetter(b byte) byte {
    switch b {
    case '2':
        return 'a'
    case '3':
        return 'd'
    case '4':
        return 'g'
    case '5':
        return 'j'
    case '6':
        return 'm'
    case '7':
        return 'p'
    case '8':
        return 't'
    default:
        return 'w'
    }
}