# The Template Method Pattern 

## Intent

Define the skeleton of an algorithm in an operation, deferring some steps to
subclasses.  Template Method lets subclasses redefine certain steps of an
algoithm without changing the algorithm's structure.

## Structure

![](../../data/template_method_pattern_uml.png)

## Terminology

| Term           | Meaning                                                                                   |
| -------------- | ----------------------------------------------------------------------------------------- |
| Abstract Class | Defines operations that concrete subclasses can implement.                                |
| Concrete Class | Implement the primitive operations to carry out subclass-specific steps of the algorithm. |

## Explanation

If you want to implement an algorithm that is very general over a few objects,
make a class that generalises the algorithm and lets the subclasses implement
the few methods that vary.

## Applicability

The Template Method pattern should be used:
 * To implement the invariant parts of an algorithm once and leave it up to
   subclasses to implement the behavior that can vary
 * When common behavior among subclasses should be factored and localised in a
   common class to avoid deuplication.
 * To control subclasses extensions. You can define a template method that
   calls "hook" operations at specific points, permitting extensions only at
   those points.
