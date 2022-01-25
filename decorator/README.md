# The Decorator Pattern

## Intent

Attach additional responsibilities to an object dynamically.  Decorators
provide a flexible alternative to subclassing for extending functionality.

## Structure

![](../data/decorator_pattern_uml.png)

## Terminology

| Term                  | Meaning                                                                                    |
| --------------------- | ------------------------------------------------------------------------------------------ |
| Component             | An interface for objects that can have responsibilities added dynamically                  |
| Concrete Component    | An object that adheres to the component interface                                          |
| Decorator             | Object that maintains a reference to a component object, and uses the components interface |
| Concrete Decorator    | A implementation of the Decorator                                                          |

## Explanation

When trying to add new functionality that conforms to the same interface (might be coincidence, might not), 
then you can wrap the component in a decorator, and the decorator implements the new functionality.

## Applicability

You can use the decorator pattern:

 * To add responsibilities to individual objects dynamically and transparently, without atffecting other objects.
 * For responsibilities that can be withdrawn.
 * When extension by subclassing is impractical. When adding lots of independent functionality, the subclasses increase massively.
