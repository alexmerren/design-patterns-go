# The Iterator Pattern

## Intent

The Iterator Pattern provides a way to access the elements of an aggregate
object sequentially without exposing its underlying representation.

## Structure

![](../data/iterate_pattern_uml.gif) 

## Terminology

| Term              | Meaning                                                                                   |
| ----------------- | ----------------------------------------------------------------------------------------- |
| Aggregate         | Defines an interface for creating an Iterator object.                                     |
| Iterator          | Defines an interface for accessing and traversing elements.                               |
| ConcreteAggregate | Implements the Iterator creation interface to return an instant of the ConcreteIterator.  |
| ConcreteIterator  | Implements the iterator interface.                                                        |

## Explanation

The purpose of the iterator pattern is two fold: Give a standard for going
through collections of objects, and being able to iterate through collections
in different ways, depending on the goal. Ideally, this would be done without
exposing how the actual collection is stored, for example, fixed-size array vs.
vector vs. stack, etc.

The pattern abstracts an iterator for any collection into two methods:
hasNext() and getNext(), you can imagine what these two methods do.

## Applicability 

The Iterator Pattern can be used:
 * To acess an aggregate object's contents without exposing internal representation
 * To support multiple traversals of aggregate objects
 * Toprovide a uniform interface for traversing aggregate structures
