# Go Collections

Provides collection classes for Dictionary, List, Queue and Stack

## Using

```bash
go get github.com/patrickhuber/go-collections
```

## List

```go
import "github.com/patrickhuber/go-collections/list"

func lister(){
    l := list.New()
    l.Append(1)
    l.Insert(0, 2)
    _ := l.Get(1)
    l.Set(1, 10)
    l.Clear()
}
```

## Stack

```go
import "github.com/patrickhuber/go-collections/list"

func stacker(){
    s := stack.New()
    s.Push(1)
    s.Push(2)
    s.Push(3)
    for{
        if s.Length() == 0{
            break
        }
        s.Pop()
    }
}
```

## Queue

```go
import "github.com/patrickhuber/go-collections/list"

func queue(){
    q := queue.New()
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)
    for{
        if q.Length() == 0{
            break
        }
        q.Dequeue()
    }
}
```