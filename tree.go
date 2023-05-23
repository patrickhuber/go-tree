package tree

import "fmt"

func New[TKey comparable, TValue any]() Tree[TKey, TValue] {
	var key TKey
	var val TValue
	return &tree[TKey, TValue]{
		Root: newNode(key, val),
	}
}

func newNode[TKey comparable, TValue any](key TKey, value TValue) *Node[TKey, TValue] {
	return &Node[TKey, TValue]{
		Key:      key,
		Value:    value,
		Children: map[TKey]*Node[TKey, TValue]{},
	}
}

func newEmptyNode[TKey comparable, TValue any](key TKey) *Node[TKey, TValue] {
	return &Node[TKey, TValue]{
		Key:      key,
		Children: map[TKey]*Node[TKey, TValue]{},
	}
}

type Tree[TKey comparable, TValue any] interface {

	// Find finds the node at the given path
	Find(path []TKey) (*Node[TKey, TValue], bool)

	// Inserts the node at the given path and fails if the parent path does not exist
	Insert(path []TKey, item TValue) (*Node[TKey, TValue], error)

	// InsertAll works like Insert but it builds out the tree path if it does not exist
	InsertAll(path []TKey, item TValue) (*Node[TKey, TValue], error)

	// Remove removes the leaf node. If there is an error it will be of type PathError
	// If the node can not be found nil and false are returned
	Remove(path []TKey) error

	// RemoveAll removes the node and all children it contains
	RemoveAll(path []TKey) error
}

type tree[TKey comparable, TValue any] struct {
	Root *Node[TKey, TValue]
}

type Node[TKey comparable, TValue any] struct {
	Key      TKey
	Value    TValue
	Children map[TKey]*Node[TKey, TValue]
}

var (
	ErrNotExist = errNotExist()
	ErrPath     = errPath()
)

func (t *tree[TKey, TValue]) Find(path []TKey) (*Node[TKey, TValue], bool) {
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

func (t *tree[TKey, TValue]) Insert(path []TKey, item TValue) (*Node[TKey, TValue], error) {
	parent := parent(path)
	n, ok := t.Find(parent)
	if !ok {
		return nil, wrapErr(ErrNotExist, "can not find item at path")
	}
	last := path[len(path)-1]
	newNode := newNode(last, item)
	n.Children[last] = newNode
	return newNode, nil
}

func (t *tree[TKey, TValue]) InsertAll(path []TKey, item TValue) (*Node[TKey, TValue], error) {
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
			child = newEmptyNode[TKey, TValue](segment)
		}

		current.Children[segment] = child
		current = child
	}
	return current, nil
}

func (t *tree[TKey, TValue]) Remove(path []TKey) error {
	return t.remove(path, false)
}

func (t *tree[TKey, TValue]) remove(path []TKey, all bool) error {
	if len(path) == 0 {
		return wrapErr(ErrPath, "path is empty")
	}
	parent := parent(path)
	p, ok := t.Find(parent)
	if !ok {
		return wrapErr(ErrNotExist, "no node at path %v", path)
	}
	key := path[len(path)-1]
	_, ok = p.Children[key]
	if !ok {
		return wrapErr(ErrNotExist, "no node at path %v", path)
	}
	// check if the node has children
	if len(p.Children) > 0 && !all {
		return wrapErr(ErrPath, "node must have no children %v. Use RemoveAll instead", path)
	}
	delete(p.Children, key)
	return nil
}

func (t *tree[TKey, TValue]) RemoveAll(path []TKey) error {
	return t.remove(path, true)
}

func parent[TKey comparable](path []TKey) []TKey {
	return path[:len(path)-1]
}

func wrapErr(err error, message string, args ...any) error {
	msg := fmt.Sprintf(message, args...)
	return fmt.Errorf("%w: %s", err, msg)
}

func errNotExist() error {
	return fmt.Errorf("does not exist")
}

func errPath() error {
	return fmt.Errorf("invalid path")
}
