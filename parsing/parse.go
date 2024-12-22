package parsing

import (
	"bufio"
	"os"
	"strconv"
)

//In the context of AOC I'd rather have the program panic then handling input error

func Atoi(s string) int{
    i, err:= strconv.Atoi(s)
    if err != nil {
        panic(err)
    }
    return i
}

func ReadLines(location string)[]string{
    file, err := os.Open(location)
    defer file.Close()
    if err != nil{
        panic(err)
    }
    scanner := bufio.NewScanner(file)
    content := []string{}
    for scanner.Scan(){
       content = append(content, scanner.Text()) 
    }
    return content
}
