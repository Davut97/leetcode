package main

func main() {
	s := []string{"flower", "flow", "flight"}
	longestCommonPrefix(s)
}
func longestCommonPrefix(strs []string) string {
	for i := 0; ; i++ {
		for _, str := range strs {
			if i == len(str) || str[i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
}
