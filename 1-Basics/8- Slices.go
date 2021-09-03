package main

import "fmt"

func main() {

    sampleSlice := make([]string, 3)
    fmt.Println("emp:", sampleSlice)

    sampleSlice[0] = "a"
    sampleSlice[1] = "b"
    sampleSlice[2] = "c"
    fmt.Println("set:", sampleSlice)
    fmt.Println("get:", sampleSlice[2])

    fmt.Println("len:", len(sampleSlice))

    sampleSlice = append(sampleSlice, "d")
    sampleSlice = append(sampleSlice, "e", "f")
    fmt.Println("appended:", sampleSlice)

    copiedSlice := make([]string, len(sampleSlice))
    copy(copiedSlice, sampleSlice)
    fmt.Println("copiedSlice:", copiedSlice)

    subSlice := sampleSlice[2:5]
    fmt.Println("subSlice1:", subSlice)

    subSlice = sampleSlice[:5]
    fmt.Println("subSlice2:", subSlice)

    subSlice = sampleSlice[2:]
    fmt.Println("ssubSlice3:", subSlice)

    anotherSliceDeclaration := []string{"g", "h", "i"}
    fmt.Println("dcl:", anotherSliceDeclaration)

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}