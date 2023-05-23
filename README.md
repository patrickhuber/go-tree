# go-tree

a nested hashmap implementation in go

- [Usage](#usage)
    - [Insert](#insert)
    - [InsertAll](#insertall)
    - [Find](#find)
    - [Remove](#remove)
    - [RemoveAll](#removeall)

# usage

You can go get the library using the following command

```bash
go get github.com/patrickhuber/go-tree
```

To use the library call the new function in the tree package

## Insert

```go
import "github.com/patrickhuber/go-tree"

func main(){
    t := tree.New[int, string]()
    t.Insert([]int{1}, "")
    t.Insert([]int{1,2}, "")
    t.Insert([]int{1,2,3}, "")
    t.Insert([]int{1,2,3,4}, "child")
}
```

## InsertAll

Instead of calling insert for each path, the InsertAll command will insert each segment of the path

```go
t := tree.New[int, string]()
t.InsertAll([]int{1,2,3,4}, "child")
```

## Find

Find retrieves the element at the given path and returns false if none can be found

```go
node, ok := t.Find([]int{1,2,3,4})
```

## Remove

Remove removes a single element from the tree

```go
node, ok := t.Remove([]int{1,2,3,4})
```

## RemoveAll 

Removes all elements in the path

```go
ok := t.RemoveAll([]int{1,2,3})
```