# The Facade Pattern

## Intent

Provide a unified interface to a set of interfaces in a subsystem. Facade
defines a higher-level interface that makes the subsystem easier to use.

## Structure

![](../data/facade_pattern_uml.png)

## Terminology

| Term              | Meaning                                                                               |
| ----------------- | ------------------------------------------------------------------------------------- |
| Facade            | Knows which subsystem class is responsible for requests, and delegates appropriately. |
| Subsystem classes | Does not know of the facade, but knows how to handle work assigned by the facade.     |

## Explanation

## Applicability

The facade pattern has a couple usecases:
 * Simplifying the interface to a complex system, that is composed of lots of classes.
 * Specifying an entry point to each level of layered software.

