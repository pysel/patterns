# Chain of Responsibility Design Pattern

The Chain of Responsibility pattern is a behavioral design pattern that allows multiple objects (handlers) to process a request in a sequential order. Each handler in the chain either handles the request or passes it to the next handler in the sequence. This pattern promotes loose coupling between sender and receiver objects, enables dynamic modification of the chain, and encourages separation of concerns by assigning specific responsibilities to individual handlers.

## Implementation

A `Server` struct is a simulation of a request/response server. A user sends a request to this server and expects to get
a response if a `reqString` is of the right format: `handlerType|arbitraryData`.

## Chain of responsibility flow

`Server` <-> `Processor` <-> `Handler`
                  |
             `Extractor`
