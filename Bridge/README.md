# Bridge Design Pattern

The Bridge pattern separates abstraction from implementation by having the abstraction contain a reference to the implementation. This enables the abstraction to delegate its operations to the implementation. This design promotes flexibility, as changing the implementation or extending the abstraction becomes simpler. The pattern is especially useful when there are multiple variations of abstractions and implementations, preventing an exponential growth in the number of classes.

## This implementation

Without bridge pattern, we would have to have 6 concrete implementations: `(Windows||Linux||Arch)(Audio||Video)`.
Using bridge pattern allows us to have an `abstraction` of OS, with 3 `refined abstractions`:

* Windows
* Linux
* ArchLinux

and an `implementer` interface - `MediaFormat` with 2 `concrete implementations`:

* AudioPlayer
* VideoPlayer

It does not save much in this case, only 1 (6 - 5) implementations, but with growth of either implementers or abstractions, the benefits of Bridge Design Pattern grow exponentially.
