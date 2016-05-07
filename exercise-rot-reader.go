package main

import (
    "io"
//    "os"
    "strings"
    "fmt"
)

type rot13Reader struct {
    r io.Reader
}

func (rot13Reader) Read(b []byte) (int, error) {
    length := len(b)
    if length == 0 {
        return length, io.EOF
    }
    for i := 0; i < length; i++ {
        b[i] = getRot13(b[i])
    }
    return len(b), nil
}

func getRot13(b byte) byte {
    var constant byte
    if (b >= 'A' && b <= 'Z') {
        constant = 'A'
    } else if (b >= 'a' && b <= 'z') {
        constant = 'a'
    } else {
        return b
    }
    rotated := (b - constant + 13) % 26
    return rotated + constant
}

func main() {
    test_case := "Lbh penpxrq gur pbqr!"

    s := strings.NewReader(test_case)
    r := rot13Reader{s}
    fmt.Println(r)
    fmt.Println(r)
    // io.Copy(os.Stdout, &r)
}
