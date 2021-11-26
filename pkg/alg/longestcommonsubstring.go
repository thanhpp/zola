package alg

func LongestCommonSubstring(str1, str2 string) int {
	str1Len := len(str1)
	str2Len := len(str2)

	// Create a table to store lengths of longest common
	// substrings. A value of the table is the length of
	// longest common substring ending with the character
	// at table[i][j].
	table := make([][]int, str1Len)
	for i := range table {
		table[i] = make([]int, str2Len)
	}

	// row and column of the table
	row := 0
	col := 0

	// loop through the table and fill the values
	for row = 0; row < str1Len; row++ {
		for col = 0; col < str2Len; col++ {
			if str1[row] == str2[col] {
				if row == 0 || col == 0 {
					table[row][col] = 1
				} else {
					table[row][col] = table[row-1][col-1] + 1
				}
			} else {
				table[row][col] = 0
			}
		}
	}

	// Find the maximum value in the table
	maxLen := 0
	for row = 0; row < str1Len; row++ {
		for col = 0; col < str2Len; col++ {
			if table[row][col] > maxLen {
				maxLen = table[row][col]
			}
		}
	}

	// Return the longest common substring length
	return maxLen
}
