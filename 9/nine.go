package main

import (
	"fmt"
	"os"
	"strconv"
)

var s = ""

func main() {
	data := readData("data")
	var total int64
	tail := len(data) - 1
	var diskLocation int64

	for i := 0; i <= tail; i++ {
		if i%2 == 0 {
			fmt.Println("S")
			sum := sumOfIndicesBetweenInclusive(diskLocation, diskLocation+int64(data[i]-1))
			id := int64(i) / 2
			total += sum * id
			diskLocation += int64(data[i])
		} else {
			openSpace := data[i]
			for openSpace != 0 {
				if openSpace >= data[tail] {
					sum := sumOfIndicesBetweenInclusive(diskLocation, diskLocation+int64(data[tail]-1))
					id := int64(tail) / 2
					total += sum * id
					diskLocation += data[tail]
					//not necessary but keeps consistant
					openSpace -= data[tail]
					data[tail] = 0
					tail = tail - 2
				} else {
					sum := sumOfIndicesBetweenInclusive(diskLocation, diskLocation+openSpace-1)
					id := int64(tail) / 2
					total += sum * id
					diskLocation += openSpace
					data[tail] = data[tail] - openSpace
					openSpace = 0
				}
			}
		}

	}

	fmt.Println(total)

}

func logstuff(i, j int64) {

	for range i {
		s = s + strconv.Itoa(int(j))
	}
}

func readData(filename string) []int64 {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	//remove eof
	bytes = bytes[:len(bytes)-1]

    data := make([]int64, len(bytes))
	for i, s := range bytes {
        data[i] = int64(s)
	}
    return data
}

func sumOfIndicesBetweenInclusive(left, right int64) int64 {
	tot := (right - left + 1) * (left + right) / 2
	fmt.Printf("%d to %d tot %d\n", left, right, tot)
	return tot
}
