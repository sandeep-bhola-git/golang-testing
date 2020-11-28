package sort

func BubbleSortDesc(elements []int) {
	for i := 0; i < len(elements); i++ {
		for j := 0; j < len(elements)-1; j++ {
			if elements[i] > elements[j] {
				elements[i], elements[j] = elements[j], elements[i]
			}
		}
	}
}
