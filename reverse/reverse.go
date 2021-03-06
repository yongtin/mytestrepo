package main

func ReverseMatch(s_ary []string) (result_ary []string) {
    for _, str := range s_ary {
        goodstr := ReverseCheck(str, s_ary)
        if len(goodstr) > 0 {
            result_ary = append(result_ary, goodstr)
        }
    }
    return result_ary
}

func ReverseCheck(s string, s_ary []string) (rating string) {
    var reverse_s string = ""
    reverse_s = Reverse(s)
    for _, this_s := range s_ary {
        if reverse_s == this_s {
            return s
        }
    }
    return ""
}

/* copying from web */
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}