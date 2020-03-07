package main

import (
    "fmt"
    "math"
)

type (
    Bucket struct {
        maxValue   int
        maxSize    int
        bucketSize int
        array      []int
        bucket     []int
    }
)

func bucketInit(size int) Bucket {
    b := Bucket{maxSize: size}
    b.array = make([]int, size)
    b.bucketSize = int(math.Ceil(math.Sqrt(float64(size))))
    b.bucket = make([]int, b.bucketSize)
    return b
}

func (b *Bucket) insert(idx, value int) {
    b.array[idx] = value
    if value > b.maxValue { b.maxValue = value }
}

func (b *Bucket) scanMinInBucket(idx int) {
    bucketIdx := idx/b.bucketSize
    b.bucket[bucketIdx] = b.array[idx]
    for i:=bucketIdx*b.bucketSize; i<(bucketIdx+1)*b.bucketSize; i++ {
        if i >= b.maxSize { break }
        if b.array[i] < b.bucket[bucketIdx] {
            b.bucket[bucketIdx] = b.array[i]
        }
    }
}

func (b *Bucket) scanMinAll() {
    bucketNum := int(math.Ceil(float64(b.maxSize)/float64(b.bucketSize)))
    for i := 0; i < bucketNum; i++ {
        b.scanMinInBucket(i*b.bucketSize)
    }
}

func (b *Bucket) getMinimum(start, end int) int {
    startBucketIdx := int(math.Ceil(float64(start)/float64(b.bucketSize)))
    endBucketIdx := int(math.Floor(float64(end)/float64(b.bucketSize)))
    min := b.maxValue + 1
    for i:=start; i<startBucketIdx*b.bucketSize; i++ {
        if b.array[i] < min {
            min = b.array[i]
        }
    }
    for bi:=startBucketIdx; bi<=endBucketIdx; bi++ {
        if b.bucket[bi] < min {
            min = b.bucket[bi]
        }
    }
    for i:=(endBucketIdx+1)*b.bucketSize; i<end; i++ {
        if b.array[i] < min {
            min = b.array[i]
        }
    }
    return min
}


func main() {
    s := readLine()
    values := getIntArray(s)

    bucket := bucketInit(len(values))
    for i, v := range values {
        bucket.insert(i, v)
    }
    bucket.scanMinAll()

    fmt.Printf("orig: %v\n", bucket.array)
    fmt.Printf("bucket: %v\n", bucket.bucket)

    fmt.Printf("2 - 13: %v\n", bucket.getMinimum(1, 12))
    fmt.Printf("5 - 5: %v\n", bucket.getMinimum(4, 4))
    fmt.Printf("7 - 11: %v\n", bucket.getMinimum(6, 10))
}
