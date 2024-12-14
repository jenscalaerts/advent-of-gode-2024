package parsing

import "strconv"

//In the context of AOC I'd rather have the program panic then handling input error

func Atoi(s string) int{
    i, err:= strconv.Atoi(s)
    if err != nil {
        panic(err)
    }
    return i
}
