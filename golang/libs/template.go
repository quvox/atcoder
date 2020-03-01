package main

import (
  "bufio"
  "fmt"
  "os"
  "sort"
  "strconv"
  "strings"
)

const (
  InputSize = 100000 // 10^5
)
var (
  rdr = bufio.NewReaderSize(os.Stdin, InputSize)
)

func readLine() string {
  buf := make([]byte, 0, InputSize)
  for {
    l, p, e := rdr.ReadLine()
    if e != nil {
      panic(e)
    }
    buf = append(buf, l...)
    if !p {
      break
    }
  }
  return string(buf)
}


func getStringArray() []string {
  return strings.Split(readLine(), " ")
}


func nextLineValue() int {
  v, _ := strconv.Atoi(readLine())
  return v
}

func nextLineValues() []int {
  return getIntArray(readLine())
}

func getInt(s string) int {
  v, _ := strconv.Atoi(s)
  return v
}

func getIntArray(s string) []int {
  var v []int
  valstr := strings.Split(s, " ")
  for _, c := range valstr {
    va, _ := strconv.Atoi(c)
    v = append(v, va)
  }
  return v
}

func digitTotal(v int) int {
  total := 0
  d := strings.Split(strconv.Itoa(v), "")
  for _, c := range d {
    dv, _ := strconv.Atoi(c)
    total += dv
  }
  return total
}

func sortDesc(v *[]int) {
  sort.Sort(sort.Reverse(sort.IntSlice(*v)))
}

func removeDupIntSlice(values *[]int) []int {
  m := make(map[int]bool)
  uniq := []int{}
  for _, v := range *values {
    if !m[v] {
      uniq = append(uniq, v)
      m[v] = true
    }
  }
  return uniq
}


func main() {
  s := readLine()
  for {
    l := len(s)
    if l == 0 { break }
    if l-len("dream") > -1 && s[l-len("dream"):] == "dream" {
      s = s[:l-len("dream")]
    } else if l-len("erase") > -1 && s[l-len("erase"):] == "erase" {
      s = s[:l-len("erase")]
    } else if l-len("dreamer") > -1 && s[l-len("dreamer"):] == "dreamer" {
      s = s[:l-len("dreamer")]
    } else if l-len("eraser") > -1 && s[l-len("eraser"):] == "eraser" {
      s = s[:l-len("eraser")]
    } else {
      fmt.Println("NO")
      return
    }
  }
  fmt.Println("YES")
}
