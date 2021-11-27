package utils

func AverageQuery(fields []string) string {
	var count string
	for _, field := range fields {
		if field == fields[len(fields)-1] {
			count += " AVG(" + field + ")" + field
			break
		}
		count += " AVG(" + field + ")" + field + ","
	}
	return count
}
