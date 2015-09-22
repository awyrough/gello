// packages
package stringutil

// reverse returns its argument string reverse rune-wise left to right
func Reverse(s string) string {
    // transform string s to rune array r
    r := []rune(s)
    // i,j starting at 0, stopping when i = len(r) / 2, i++, j--
    for i,j := 0, len(r)-1; i < len(r) / 2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}