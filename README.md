# Go Collections

Provides collection classes for Dictionary, List, Queue and Stack

## Using

```bash
go get github.com/patrickhuber/go-collections
```

## List

List of Any

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

Generic List

```go
import "github.com/patrickhuber/go-collections/generic/list"

func lister(){
    l := list.New[int]()
    l.Append(1)
    l.Insert(0, 2)
    _ := l.Get(1)
    l.Set(1, 10)
    l.Clear()
}
```

Concurrent List

```go
import "github.com/patrickhuber/go-collections/concurrent/list"

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
import "github.com/patrickhuber/go-collections/stack"

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

Generic Stack

```go
import "github.com/patrickhuber/go-collections/generic/stack"

func stacker(){
    s := stack.New[int]()
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
import "github.com/patrickhuber/go-collections/queue"

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

Generic Queue

```go
import "github.com/patrickhuber/go-collections/generic/queue"

func queue(){
    q := queue.New[int]()
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

## Dictionary

```go
import "github.com/patrickhuber/go-collections/dictionary"

func dict(){
    d := dictionary.New()
    d.Set("test", "a")
    d.Set("other", "b")
    _ := d.Get("test")
    d.Remove("other")
    d.Clear()
}
```

Generic Dictionary

```go
import "github.com/patrickhuber/go-collections/generic/dictionary"

func dict(){
    d := dictionary.New[string, string]()
    d.Set("test", "a")
    d.Set("other", "b")
    _ := d.Get("test")
    d.Remove("other")
    d.Clear()
}
```