/*
Adapter pattern

Type: structural (combines capabilities of two independent interfaces)

Adapter pattern is used the same way as adapters in real life. To be able to use some object as another.
Example: you have a USB port on your machine, but an Ethernet cable. Solution: use adapter

In software design, adapter allows you to join functionalities of incompatible interfaces.

As an example, we have 2 "classes":
    1. Sparrow implementing Bird trait
    2. ToyDuck that can squeeeeeak

Suppose we want to use Sparrow as ToyDuck: in this case, to invoke squeeeeeak method on Sparrow.
For that, we use adapter (1)
*/


trait Bird {
    fn fly(&self);
    fn make_sound(&self);
}

trait Duck {
    fn squeak(&self);
}

struct ToyDuck {}
struct Sparrow {}

impl Bird for Sparrow {
    fn fly(&self) {
        println!("I am flying");
    }

    fn make_sound(&self) {
        println!("Chirp chirp");
    }
}

impl Duck for ToyDuck {
    fn squeak(&self) {
        println!("Squeeeeeeak");
    }
}

// (1)
// This struct holds a bird object and has a squeak method implemented.
// Now it is possible to say that we are able to use BirdAdapter with a Sparrow inside as a ToyDuck
struct BirdAdapter {
    bird: Sparrow
}

impl Duck for BirdAdapter {
    fn squeak(&self) {
        self.bird.make_sound()
    }
}

fn main() {
    // define classes
    let sparrow: Sparrow = Sparrow {};
    let toy_duck: ToyDuck = ToyDuck {};

    let pseudo_duck: BirdAdapter = BirdAdapter { bird: sparrow };

    // using Sparrow as Duck
    pseudo_duck.squeak();
    toy_duck.squeak();
}