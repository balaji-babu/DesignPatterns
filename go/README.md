## Introduction:
* Design Pattern is a well described solution for a common software problems. Its a blueprint that can be customized to solve a particular design problem in code.
* Design Patterns are divided into three categories
  * Creational 
  * Structural
  * Behavioral

## Creational Design Pattern:
> Deals with object creation mechanisms, trying to create objects in a suitable manner. The basic form of object creation could result in a design problems or added complexity to the design. Creational design patterns solves this problem by controlling the object creation.
* Types:
  * Singleton Pattern
    * > It ensures that only one struct of its kind exist and provides a single access through out the code.
  * Factory Patten 
    * > Create an object without exposing the creation logic to the client. The purpose of this approach is to abstract the user form the things like dependencies that needed to create the class. The user only need to interface to get this value.
    * Responsible solely for the wholesale (not piecewise, unlike Builder) creation of objects.
  * Builder Pattern
    * > Builder pattern is a design pattern that allows for the step-by-step[piece-by-piece] creation of complex objects using the correct sequence of actions. The construction is controlled by a director object that only needs to know the type of object it is to create. 
    * Some objects are simple and can be created in a single constructor call, other objects require a lot of ceremony to create.
    * Having a factory function with 10 arguments is not productive, instead opt for piecewise (piece-by-piece) construction
    * > **Builder Facets:** In some situation we need to create more than one builder class. Segregating the builder into multiple builders, gives us control over the properties of a struct.
  * Prototype Pattern
    * > Prototype pattern provides a mechanism to copy the original object to a new object and them modify it according to our needs. This ensures the increase performance.
    * Complicated objects aren't designed from scratch they reiterate existing designs.
    * An existing design[Partially or Fully Constructed] design is a prototype
    * Make a copy of the prototype and customize it, requires a deep copy support.

## Structural Design Pattern:
> Deals with designing by identifying a simple way to realize relationships between entities. Using inheritance and composition to create a large object from small objects.
* Types:
  * Adapter Pattern
    > * Adapter pattern enables collaboration between objects with incompatible interfaces. Adapter pattern allows to adapt behavior between objects.
    * A construct which adapts an existing interface X to conform to the required interface Y.
    * Can separate the interface or data conversion code from the primary business logic of the program.
    * Open-Closed principle. Can introduce new types of adapters to the program without modifying existing code, by working with the adapters through the interface.
  * Bridge Pattern
    > * Bridge pattern allows to split a large class or a group of closely related classes, into two separate hierarchies(abstraction and implementation) that can be developed independently of one another.
     * Can define the needed code in separate structs and interfaces to break the dependency to code each business logic for similar places.
     * Could reduce the code coupling.
  * Composite Pattern
    > * It allows us to group objects from the same family in a tree-like structure and lets client treat individual objects and compositions uniformly.
    * Commonly used to solve the cases when we need to handle tree structures because it,s easy to iterate each one of its child/inner objects.
    * To define the object we can use a struct type or a interface type.
  * Decorator Pattern
    > * Essentially allow us to wrap existing functionality and append or prepend our own custom functionality on top of them.
    * Augument an object with additional functionality.
    * Do not want to rewrite or alter existing code [OCP]
    * If want to keep new functionality separate [SRP]
    * Need to be able to interact with existing structures.
    * Sol: Embed the decorated object and provide additional functionality.
  * Facade Pattern
    > * Encapsulates a complex subsystem behind a simple interface. It hides much of the complexity and makes the subsystem easy to use.
    * If we need to use the complex subsystem directly, we can do that too, there is no necessity to use the facade system all the time.
    * It decouples a client implementation from the complex subsystem.
  * Flyweight Pattern
    > * Flyweight is used to manage the state of an object with high variation. The pattern allows us to share common parts of the object state among multiple objects, instead of each object storing it. 
    * It helps to reduce the overall memory usage and the object initializing overhead. The pattern helps create interclass relationships and lower memory to a manageable level.
    * Avoid redundancy when storing data.

  * Proxy Pattern

## Behavioral Design Pattern:
> Deals with better interaction between objects and to provide lose coupling and flexibility to extend easily, without breaking the code.
* Types:
  * Template Method Pattern
  * Visitor Pattern
  * Chain of Responsibility Pattern
  * Command Pattern
  * Iterator Pattern
  * Mediator Pattern
  * Memento Pattern
  * Observer Pattern
  * State Pattern
  * Strategy Pattern