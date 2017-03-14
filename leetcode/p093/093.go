package p093


/**
Given a string containing only digits, restore it by returning all possible valid IP address combinations.

For example:
Given "25525511135",

return ["255.255.11.135", "255.255.111.35"]. (Order does not matter)
 */

func restoreIpAddresses(s string) []string {

	ss := []byte(s)
	rbs := make([]byte, len(ss)+3)
	res := make([]string, 0)

	atoi := func(s, d int) bool {
		seg := 0
		for i := s; i < s+d; i++ {
			seg = seg*10 + int(ss[i]-'0')
		}
		return (d == 1 && seg < 10) || (d == 2 && seg >= 10 && seg < 100) || seg >=100 && seg <= 255
	}
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				for l := 1; l <= 3; l++ {
					if i+j+k+l != len(ss) {
						continue
					}
					if !atoi(0, i) || !atoi(i, j) || !atoi(i+j, k)|| !atoi(i+j+k, l){
						continue
					}
					copy(rbs[0:i], ss[0:i])
					rbs[i] = '.'
					copy(rbs[i+1:i+j+1], ss[i:i+j])
					rbs[i+j+1] = '.'
					copy(rbs[i+j+2:i+j+k+2], ss[i+j:i+j+k])
					rbs[i+j+k+2] = '.'
					copy(rbs[i+j+k+3:], ss[i+j+k:])
					res = append(res, string(rbs))
				}
			}
		}
	}
	return res
}
