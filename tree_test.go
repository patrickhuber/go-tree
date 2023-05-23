package tree_test

import (
	"fmt"
	"testing"

	"github.com/patrickhuber/go-tree"
	"github.com/stretchr/testify/require"
)

func TestCanInsert(t *testing.T) {
	tree := tree.New[string, string]()

	pathToBuild := []string{"root", "test"}
	err := buildPath(tree, pathToBuild)
	require.Nil(t, err)
}

func TestCanFind(t *testing.T) {
	tree := tree.New[string, string]()
	pathToBuild := []string{"root", "parent", "child", "grandchild"}
	buildPath(tree, pathToBuild)
	n, ok := tree.Find(pathToBuild)
	require.True(t, ok)
	require.NotNil(t, n)
}

func TestCanInsertAll(t *testing.T) {
	tree := tree.New[string, string]()
	key := "child"
	path := []string{"grand", "parent", key}
	expected := "gary"
	n, err := tree.InsertAll(path, expected)
	require.Nil(t, err)
	require.NotNil(t, n)
	require.Equal(t, expected, n.Value)
	require.Equal(t, key, n.Key)

	n, ok := tree.Find(path)
	require.True(t, ok)
	require.NotNil(t, n)
	require.Equal(t, expected, n.Value)
	require.Equal(t, key, n.Key)
}

func TestCanInsertAllWithOverlap(t *testing.T) {
	tree := tree.New[string, string]()
	key := "child"
	path := []string{"grand", "parent", key}
	expected := "gary"
	n, err := tree.InsertAll(path, expected)
	require.Nil(t, err)
	require.NotNil(t, n)
	require.Equal(t, expected, n.Value)
	require.Equal(t, key, n.Key)

	key = "2nd"
	expected = "sam"
	n, err = tree.InsertAll([]string{"grand", "parent", key}, expected)
	require.Nil(t, err)
	require.NotNil(t, n)
	require.Equal(t, expected, n.Value)
	require.Equal(t, key, n.Key)
}

func TestCanRemove(t *testing.T) {
	tree := tree.New[string, string]()
	key := "child"
	path := []string{"grand", "parent", key}
	expected := "gary"
	_, err := tree.InsertAll(path, expected)
	require.Nil(t, err)

	err = tree.Remove(path)
	require.NotNil(t, err)

	p, ok := tree.Find(path[:len(path)-1])
	require.True(t, ok)
	require.NotNil(t, p)
}

func TestCanRemoveAll(t *testing.T) {
	tree := tree.New[string, string]()
	key := "child"
	path := []string{"grand", "parent", key}
	_, err := tree.InsertAll(path, "gary")
	require.Nil(t, err)
	tree.RemoveAll(path[0:0])
}

func buildPath(tr tree.Tree[string, string], pathToBuild []string) error {
	var path []string
	for _, segment := range pathToBuild {
		path = append(path, segment)
		n, err := tr.Insert(path, segment)
		if err != nil {
			return err
		}
		if n == nil {
			return fmt.Errorf("unable to insert node at segment %s", segment)
		}
	}
	return nil
}
