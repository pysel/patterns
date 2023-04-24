package main

import "fmt"

// Component: an interface to represent both Leaf and Composite
type FileSystemNode interface {
	GetSize() uint64
}

// Leaf
type File struct {
	Size uint64
}

// Composite
type Directory struct {
	Children []FileSystemNode
}

func (f *File) GetSize() uint64 {
	return f.Size
}

func (d *Directory) GetSize() uint64 {
	var size uint64
	for _, child := range d.Children {
		size += child.GetSize()
	}
	return size
}

func main() {
	// Check README.md for better visualization
	root := &Directory{
		Children: []FileSystemNode{
			&Directory{
				Children: []FileSystemNode{
					&File{Size: 100},
					&File{Size: 200},
					&Directory{
						Children: []FileSystemNode{
							&File{Size: 300},
							&File{Size: 400},
						},
					},
				},
			},
			&Directory{
				Children: []FileSystemNode{
					&File{Size: 500},
					&File{Size: 600},
				},
			},
		},
	}

	fmt.Println("Size of root directory:", root.GetSize())
}
