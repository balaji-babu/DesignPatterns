## DesignPatterns:
This repo contains code for SOLID design principles and design patterns in go programming language.

## How to run this code?
Go to the respective directory and run the main.go [go run main (or) go build].

## How to install golang in work environment?
Please refer to [golang](https://go.dev/doc/install) official doc.

## Introduction:
* Design patterns are typically for object oriented based languages.
* In golang we don't have Inheritance and week encapsulation, so golang is not fully object oriented.

## Solid Design Principles:
* SOLID is an acronym for five sets of principles that was first coined by "Robert C Martin"
* The Five principles of object oriented designs are:
  * SIP - Single Responsibility Principle
  * OCP - Open-Closed Principle
  * LSP - Liskov Substitution Principle
  * ISP - Interface Segregation Principle
  * DIP - Dependency Injection Principle

## Single Responsible Principle:
> A class(type: struct) should have one, and only one reason to change.
* Ignoring the above rules, will lead to [GOD Object](https://medium.com/@carlos.ariel.mamani/the-god-object-or-god-class-anti-pattern-bfb8c15eb513) 
* Module should be highly cohesive and loosely coupled.

## Open Closed Principle:
> Entities should be open for extension, but closed for modification.
* Code doesn't have to be changed every time when the requirements change.
* In golang, code should be able to override a struct, using Stratergy Design Pattern

## Liskov Substitution Principle:
> LSP states that objects of a superclass should be replaceable with objects of its subclasses without breaking the application.
* eg: API which works correctly in base class, when extended by the derived class, it should work correctly in the derived class.
* LSP deals with Inheritance, which is primarily not applicable in golang, but we can apply LSP in go, using interface and polymorphism.

## Interface Segregation Principle:
> Clients should not be forced to depend upon interfaces that they do not use.
* The goal is to reduce the side effects and frequency of required changes by splitting the software into multiple, independent parts.
* The above statement is achievable if we define our interface for a specific task.

