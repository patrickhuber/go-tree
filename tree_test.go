package tree_test

import (
	"fmt"
	"testing"

	"github.com/patrickhuber/go-tree"
	"github.com/stretchr/testify/require"
)

func TestCanInsert(t *testing.T) {
	tree := tree.New[string]()

	pathToBuild := []string{"root", "test"}
	err := buildPath(tree, pathToBuild)
	require.Nil(t, err)
}

func TestCanFind(t *testing.T) {
	tree := tree.New[string]()
	pathToBuild := []string{"root", "parent", "child", "grandchild"}
	buildPath(tree, pathToBuild)
	n, ok := tree.Find(pathToBuild)
	require.True(t, ok)
	require.NotNil(t, n)
}

func TestCanInsertAll(t *testing.T) {
	tree := tree.New[string]()
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

func buildPath(tr tree.Tree[string], pathToBuild []string) error {
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
