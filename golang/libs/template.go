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
  InputSize = 500000 // 10^5
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

func removeDupStringSlice(values *[]string) []string {
  m := make(map[string]bool)
  uniq := []string{}
  for _, v := range *values {
    if !m[v] {
      uniq = append(uniq, v)
      m[v] = true
    }
  }
  return uniq
}

func countUniqIntInSlice(values *[]int) map[int]int {
  uniq := make(map[int]int)
  for _, v := range *values {
    if _, ok := uniq[v]; !ok {
      uniq[v] = 1
    } else {
      uniq[v] += 1
    }
  }
  return uniq
}

func countUniqStringInSlice(values *[]string) map[string]int {
  uniq := make(map[string]int)
  for _, v := range *values {
    if _, ok := uniq[v]; !ok {
      uniq[v] = 1
    } else {
      uniq[v] += 1
    }
  }
  return uniq
}


func main() {
  n := nextLineValue()
  s := readLine()
  inputString := make([]string, n)
  for i:=0; i<n; i++ {
    inputString[i] = string(s[i])
  }
  uniq := countUniqStringInSlice(&inputString)

  q := nextLineValue()
  queries := make([]string, q)
  for i:=0; i<q; i++ {
    queries[i] = readLine()
  }

  for i:=0; i<q; i++ {
    query := strings.Split(queries[i], " ")
    num, _ := strconv.Atoi(query[1])
    if query[0] == "1" {
      inputString[num-1] = query[2]
    } else {
      numEnd, _ := strconv.Atoi(query[2])
      toCheck := inputString[num-1:numEnd]
      fmt.Println(len(uniq))
    }
  }
}
