package sort

//SortInsertion function to implement insertion sort
func SortInsertion(data []int) {
	n := len(data)
	var j int
	for i := 0; i < n; i++ {
		j = i
		for j > 0 && data[j] < data[j-1] {
			data[j], data[j-1] = data[j-1], data[j]
			j--
		}
	}
}
