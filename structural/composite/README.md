# The Composite Pattern

## Intent

Compose objects into tree structures to represent part-whole hierarchies.
Composite lets clients treat individual objects and compositions of objects
uniformly.

## Structure

![](../../data/composite_pattern_uml.png) 

## Terminology

| Term      | Meaning                                                                                                               |
| --------- | --------------------------------------------------------------------------------------------------------------------- |
| Leaf      | Represents individual objects in a composition, it has no children.                                                   |
| Composite | Represents an object that has one to many children, and stores child components and operations.                       |
| Component | Declares the interface for objects in the composition, defines an interface for accessing and managing components.    |

## Explanation

The book follows on nicely from the iterator pattern, as the composite
pattern is linked to the need for menus within menus, giving us a tree-like
structure.

From the definition, I think it means that every if an object is composed of
other objects, something should be able to handle it

This pattern allows us to build structure of objects in the form of trees
that contain both compositions of objects and individual objects as nodes.

> We can apply operations over both composites and individual objects. 

Since all the components in a composite must implement the interface, and
some of the methods don't belong to one class or the other, sometimes the
best you can do is throw a runtime exception.

A drawback of the composite pattern is that is does two things: manages
hierarchies and performs operations on those hierarchies. In principle, it
takes the Single Responsibility principle and trades it for transparency

## Applicability 

You can use the Composite Pattern if:
 * You want to represent part-whole hierarchies of objects.
 * You want to treat compositions of objects and individual objects uniformly. 
