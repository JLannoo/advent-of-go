package day7

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Name     string
	Size     int
	IsFolder bool
	Children []*TreeNode
	Parent   *TreeNode
}

func (node *TreeNode) AddChild(child TreeNode) {
	if !node.IsFolder {
		panic("Cannot add child to non-folder node")
	}

	node.Children = append(node.Children, &child)
}

func (node *TreeNode) Recurse(fn func(TreeNode)) {
	fn(*node)

	for _, child := range node.Children {
		child.Recurse(fn)
	}
}

func (node *TreeNode) GetTotalSize() int {
	if !node.IsFolder {
		return node.Size
	}

	var total int

	for _, child := range node.Children {
		total += child.GetTotalSize()
	}

	node.Size = total
	return total
}

func (node *TreeNode) GetFullPath() string {
	if node.Parent == nil {
		return node.Name
	}

	return node.Parent.GetFullPath() + "/" + node.Name
}

type Tree struct {
	Root          TreeNode
	CurrentFolder *TreeNode
}

func (tree *Tree) ChangeDirectory(name string) {
	if name == ".." {
		tree.GoUp()
		return
	}

	if name == "/" {
		tree.GoToRoot()
		return
	}

	for _, child := range tree.CurrentFolder.Children {
		if child.Name == name && child.IsFolder {
			tree.CurrentFolder = child
			return
		}
	}

	panic("Could not find folder " + name)
}

func (tree *Tree) GoUp() {
	tree.CurrentFolder = tree.CurrentFolder.Parent
}

func (tree *Tree) GoToRoot() {
	tree.CurrentFolder = &tree.Root
}

func (tree *Tree) CreateDirectoryContents(instruction Command) {
	for _, line := range instruction.Output {
		split := strings.Split(line, " ")
		isDir := split[0] == "dir"

		if isDir {
			name := split[1]

			tree.CurrentFolder.Children = append(tree.CurrentFolder.Children, &TreeNode{
				Name:     name,
				IsFolder: true,
				Parent:   tree.CurrentFolder,
			})
		} else {
			size, err := strconv.Atoi(split[0])
			if err != nil {
				panic(err)
			}

			name := split[1]

			tree.CurrentFolder.Children = append(tree.CurrentFolder.Children, &TreeNode{
				Name:     name,
				IsFolder: false,
				Parent:   tree.CurrentFolder,
				Size:     size,
			})
		}
	}
}

func (tree *Tree) GetDirectoriesWithSize() []TreeNode {
	directories := []TreeNode{}
	sizes := []int{}

	tree.Root.Recurse(func(node TreeNode) {
		if node.IsFolder {
			directories = append(directories, node)
			sizes = append(sizes, node.GetTotalSize())
		}
	})

	return directories
}

func (tree *Tree) GetDirectoriesSmallerThan(max int) []TreeNode {
	directories := []TreeNode{}

	tree.Root.Recurse(func(node TreeNode) {
		size := node.GetTotalSize()

		if size < max && node.IsFolder {
			directories = append(directories, node)
		}
	})

	return directories
}
