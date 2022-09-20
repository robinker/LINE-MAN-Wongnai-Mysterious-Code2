package decoder


func DecodeZigZagPattern(cipherText string, numRails int) string {
	var rails [][]string
	for i := 0; i < numRails; i++ {
		rails = append(rails, make([]string, len(cipherText)))
	}
	zigzagPattern := drawZigZagPattern(cipherText, numRails)
	rails = solvZigZagPattern(cipherText, numRails, zigzagPattern)
	return readZigZagPattern(rails, cipherText, numRails)
}

func drawZigZagPattern(cipherText string, numRails int) [][]string {
	var rails [][]string
	for i := 0; i < numRails; i++ {
		rails = append(rails, make([]string, len(cipherText)))
	}
	row := 0
	dir := -1
	for col, _ := range cipherText {
		rails[row][col] = "*"
		if row == numRails-1 || row == 0  {
			dir *= -1
		}
		row += dir
	}
	return rails
}

func solvZigZagPattern(cipherText string, numRails int, rails [][]string) [][]string {
	index := 0
	for row := 0; row < numRails; row++ {
		for col := range cipherText {
			if rails[row][col] == "*" && index < len(cipherText)  {
				rails[row][col] = string(cipherText[index])
				index += 1
			}
		}
	}
	return rails
} 

func readZigZagPattern(rails [][]string, cipherText string, numRails int) string {
	result := ""
	row := 0
	dir := -1
	for col := range(cipherText) {
		result += rails[row][col]
		if row == numRails-1 || row == 0  {
			dir *= -1
		}
		row += dir
	}
	return result
}