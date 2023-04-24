# Composite Design Pattern

Type: Structural

The Composite pattern is a structural design pattern that enables you to represent complex, hierarchical structures using a unified interface. It allows you to treat individual objects and compositions of objects uniformly by composing them into tree-like structures. This pattern is particularly useful when you need to manage objects with part-whole relationships.

In the Composite pattern, there are three main components:

* Component: An interface that defines common operations for both simple and complex objects. This interface ensures that all objects in the hierarchy can be treated uniformly.
* Leaf: A class that represents simple, individual objects without any nested elements. These are the building blocks of the structure.
* Composite: A class that represents complex objects, which can contain other objects, either simple (Leaves) or complex (other Composites). Composites manage child objects through the Component interface.

By using the Composite pattern, you can simplify operations on hierarchical structures and manipulate them without knowing the specific type of object (Leaf or Composite) you're working with.

## Implementation

I implemented it by imagining a hierarchy of files and directories in a filesystem. Directories play a role of `composites` (they may contain other directories-composites or files-leafs) and files are `leafs`.

We represent both Directories and Files with `FileSystemNode` interface.

In `main.go`, the root variable represents this file structure:

```mathematica
Root (Directory)
│
├── Dir1 (Directory)
│   ├── File1 (File)
│   ├── File2 (File)
│   └── Dir1_1 (Directory)
│       ├── File3 (File)
│       └── File4 (File)
│
└── Dir2 (Directory)
    ├── File5 (File)
    └── File6 (File)
```
