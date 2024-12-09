package main

import (
	"fmt"
	"os"
	"slices"
)

var s = ""
var emptyFile = file{}

const emptyId = -1
func main(){
    part2()
}

func part2() {
	data := readData("data")
	files := make([]file, len(data))
	for index, length := range data {
		if length == 0 {
			continue
		}
		if index%2 == 0 {
			files[index] = file{
				id:     index / 2,
				length: int(length),
			}
		} else {
			files[index] = file{
				id:     -1,
				length: int(length),
			}
		}
	}
	for index := len(files) - 1; index >= 0; index-- {
		nextFile := files[index]
		if nextFile.id != emptyId {
			match, matchedFile := findFirstMatching(nextFile.length, files[:index])
			if match != -1 {
				files[index] = file{length: nextFile.length, id: emptyId}
				files[match] = nextFile
				if nextFile.length != matchedFile.length {
					files = slices.Insert(files, match+1, file{id: emptyId, length: matchedFile.length - nextFile.length})
					index++
				}
			}

		}
	}

    fmt.Println(calculateHash(files))
}

func calculateHash(files []file)int{

	diskLocation := 0
	var total int
	for _, file := range files {
		if file.id > 0 {
			total += sumOfIndicesBetweenInclusive(int(diskLocation), int(file.length+diskLocation-1))*int(file.id)
		}
		diskLocation += file.length
	}
    return total
}
func print(files []file) {

	for _, file := range files {
		for range file.length {
			if file.id == -1 {
				fmt.Print(".")
			} else {
				fmt.Print(file.id)
			}
		}
	}
    fmt.Println()
}

func findFirstMatching(minimalLength int, remainingFiles []file) (int, file) {
	for i, file := range remainingFiles {
		if file.id == -1 && minimalLength <= file.length {
			return i, file
		}
	}
	return -1, emptyFile
}

func part1() {
	data := readData("data")
	var total int
	tail := len(data) - 1
	var diskLocation int

	for i := 0; i <= tail; i++ {
		if i%2 == 0 {
			sum := sumOfIndicesBetweenInclusive(diskLocation, diskLocation+int(data[i]-1))
			id := int(i) / 2
			total += sum * id
			diskLocation += int(data[i])
		} else {
			openSpace := data[i]
			for openSpace != 0 {
				if openSpace >= data[tail] {
					sum := sumOfIndicesBetweenInclusive(diskLocation, diskLocation+int(data[tail]-1))
					id := int(tail) / 2
					total += sum * id
					diskLocation += data[tail]
					//not necessary but keeps consistant
					openSpace -= data[tail]
					data[tail] = 0
					tail = tail - 2
				} else {
					sum := sumOfIndicesBetweenInclusive(diskLocation, diskLocation+openSpace-1)
					id := int(tail) / 2
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

func readData(filename string) []int {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	//remove eof
	bytes = bytes[:len(bytes)-1]

	data := make([]int, len(bytes))
	for i, s := range bytes {
		data[i] = int(s - '0')
	}
	return data
}



func sumOfIndicesBetweenInclusive(left, right int) int {
	tot := (right - left + 1) * (left + right) / 2
	return tot
}

type file struct {
	id     int
	length int
}
