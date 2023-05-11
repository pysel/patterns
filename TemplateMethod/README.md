# Template Method Pattern

Type: `Behavioral`

## Overwiew

The Template Method pattern is a behavioral design pattern that defines a skeleton of an algorithm in a base class and allows subclasses to override specific steps without changing the overall structure. This pattern enables code reusability and separation of concerns, as common functionality is implemented in the base class, while unique variations are handled in derived classes.

To use the Template Method pattern, create an abstract base class with a template method, declare its steps as virtual or abstract, and override these steps in concrete derived classes as needed
