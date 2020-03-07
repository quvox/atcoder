package main

import (
    "fmt"
)

type (
    UnionNode struct {
        parent *UnionNode
        depth  int
        size   int
        name   int
        value  interface{}
        id     int
    }

    UnionArray struct {
        array   []UnionNode
        idMap   map[int]int
        count   int
    }
)

func (u *UnionNode) findRoot() *UnionNode {
    if u.parent == nil { return u }
    return u.parent
}

func (u *UnionNode) getSize() int {
    root := u.findRoot()
    return root.size
}

func merge(t1, t2 *UnionNode) *UnionNode {
    if t1.depth >= t2.depth {
        t2.parent = t1
        t2.depth = 0
        t1.size += t2.size
        return t1
    } else {
        t1.parent = t2
        t1.depth = 0
        t2.size += t1.size
        return t2
    }
}

func belongSame(u1, u2 *UnionNode) bool {
    r1 := u1.findRoot()
    r2 := u2.findRoot()
    return r1.name == r2.name
}

func (ua *UnionArray) insertNewNode(idx, value int) {
    ua.array[idx] = UnionNode{value: value, size: 1, id: ua.count}
    ua.idMap[ua.count] = idx
    ua.count += 1
}

func (ua *UnionArray) getNode(name int) *UnionNode{
    idx := ua.idMap[name]
    return &ua.array[idx]
}

func (ua *UnionArray) merge(nameParent, nameChild int) {
    n1 := ua.getNode(nameParent)
    n2 := ua.getNode(nameChild)
    merge(n1, n2)
}


func initUnionNodeArray(size int) UnionArray{
    return UnionArray{array: make([]UnionNode, size), count: 0, idMap: map[int]int{}}
}


func main() {
    v0 := getInt(readLine())
    uarray := initUnionNodeArray(v0)
    for i:=0; i<v0; i++ {
        uarray.insertNewNode(i, i+1)
    }

    vn := getInt(readLine())
    for i:=0; i<vn; i++ {
        v := getIntArray(readLine())
        for k:=1; k<len(v); k++ {
            uarray.merge(v[0], v[i])
        }
    }

    for i:=0; i<len(uarray.array); i++ {
        fmt.Printf("%d ", uarray.array[i].findRoot().name)
    }
}
