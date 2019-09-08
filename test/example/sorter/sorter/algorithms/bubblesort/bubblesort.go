package bubblesort


func BubbleSort(values []int) {
	flag := true
	for i := 0; i <= len(values)-1; i++ {
		flag = true
		total := len(values)-1
		for j:=0; j < total ;j++  {
			for j:=0; j < total -i -1 ;j++  {
				if values[j] > values[j+1] {
					values[j],values[j+1] = values[j+1],values[j]
					flag = false
				}
			}
			if flag == true {
				break
			}
		}
	}
}