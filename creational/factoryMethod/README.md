# The Factory Method Pattern

## Intent

Define an interface for creating an object, but let subclasses decide which
class to instantiate. Factory Method lets a class defer instantiation to
subclasses.

## Structure

![](../../resources/factory_method_pattern_uml.png)

## Terminology

| Term              | Meaning                                                                   |
| ----------------- | ------------------------------------------------------------------------- |
| Product           | Defines the interface of objects that the factory method creates          |
| Creator           | Declares the factory method, that returns objects of type Product         |
| Concrete Product  | Implements the Product interface                                          |
| Concrete Creator  | Overrides the factory method to return an instant of a ConcreteProduct    |

## Explanation

The Factory Method pattern is very straight-forward: Have a factory that can be
extended to have different subclasses define what that factory makes.

## Applicability

You can use the factory method pattern:
 * When a class can't anticipate the class of objects it must create.
 * When a class wants its subclasses to specify the objects it creates.
 * When classes delegate responsibility to one of several helper subclasses.
