package _6

import "bytes"

func convert(s string, numRows int) string {
	rows, zigRowIndex, zagRowIndex := make([][]byte, numRows), 0, numRows-2
	for i := range s {
		if zigRowIndex >= numRows && zagRowIndex <= 0 {
			// next zigzag
			zigRowIndex = 0
			zagRowIndex = numRows - 2
		}

		// s[i] is one byte for one rune since the s contains only letters , or .
		if zigRowIndex < numRows {
			rows[zigRowIndex] = append(rows[zigRowIndex], s[i])
			zigRowIndex++
		} else if zagRowIndex > 0 {
			rows[zagRowIndex] = append(rows[zagRowIndex], s[i])
			zagRowIndex--
		}
	}

	buf := bytes.Buffer{}
	for _, row := range rows {
		buf.Write(row)
	}
	return buf.String()
}
