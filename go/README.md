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
  * Composite Pattern
  * Decorator Pattern
  * Facade Pattern
  * Flyweight Pattern
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