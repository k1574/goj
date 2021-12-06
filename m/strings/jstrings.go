package jstrings

/* Split by sep until reaching endsep and return splitted in an array. */
func
SplitTilSep(s, sep, endsep string) []string {
        n := strings.LastIndex(s, endsep)
        var arg, str string
        if n != -1 {
                arg = s[:n]
                str = s[n+len(endsep):]
        } else {
                arg = s
                str = ""
        }

        return append(strings.Split(arg, sep), str)
}

