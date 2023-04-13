package buscaBinaria

func buscaBinaria(array []int, element int) int{
	start := 0
	end := len(array)-1
	for start <= end {
		mid:= (start+end)/2
		if array[mid] == element {
			return mid
		}else if array[mid] < element{
			start = mid + 1
		}else{
			end = mid - 1
		}
	}
}