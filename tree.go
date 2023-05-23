package tree

import "fmt"

func New[T any]() Tree[T] {
	var t T
	return &tree[T]{
		Root: newNode("", t),
	}
}

func newNode[T any](key string, value T) *Node[T] {
	return &Node[T]{
		Key:      key,
		Value:    value,
		Children: map[string]*Node[T]{},
	}
}

func newEmptyNode[T any](key string) *Node[T] {
	return &Node[T]{
		Key:      key,
		Children: map[string]*Node[T]{},
	}
}

type Tree[T any] interface {
	// Finds the node at the given path
	Find(path []string) (*Node[T], bool)
	// Inserts the node at the given path and fails if the parent path does not exist
	Insert(path []string, item T) (*Node[T], error)
	// InsertAll works like Insert but it builds out the tree path if it does not exist
	InsertAll(path []string, item T) (*Node[T], error)
}

type tree[T any] struct {
	Root *Node[T]
}

type Node[T any] struct {
	Key      string
	Value    T
	Children map[string]*Node[T]
}

var (
	ErrNotExist = errNotExist()
)

func errNotExist() error {
	return fmt.Errorf("does not exist")
}

func (t *tree[T]) Find(path []string) (*Node[T], bool) {
	if t.Root == nil {
		return nil, false
	}
	current := t.Root
	for _, segment := range path {
		child, ok := current.Children[segment]
		if !ok {
			return nil, false
		}
		current = child
	}
	return current, true
}

func (t *tree[T]) Insert(path []string, item T) (*Node[T], error) {
	parent := path[:len(path)-1]
	n, ok := t.Find(parent)
	if !ok {
		return nil, wrapErr(ErrNotExist, "can not find item at path")
	}
	last := path[len(path)-1]
	newNode := newNode(last, item)
	n.Children[last] = newNode
	return newNode, nil
}

func (t *tree[T]) InsertAll(path []string, item T) (*Node[T], error) {
	if t.Root == nil {
		return nil, wrapErr(ErrNotExist, "missing root")
	}
	current := t.Root
	for i, segment := range path {
		child, ok := current.Children[segment]
		if ok {
			current = child
			continue
		}

		// last segment we insert the item
		if i == len(path)-1 {
			child = newNode(segment, item)
		} else {
			child = newEmptyNode[T](segment)
		}

		current.Children[segment] = child
		current = child
	}
	return current, nil
}

func wrapErr(err error, message string, args ...any) error {
	msg := fmt.Sprintf(message, args...)
	return fmt.Errorf("%w %s", err, msg)
}
