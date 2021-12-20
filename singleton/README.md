# The Singleton Pattern

The singleton pattern ensures that there is only one instance of an object.

This is useful in applications like thread pools, caches, loggers, etc.

We use singleton objects instead of global variables as they are used when
needed and not made if they are not needed in the program.

## Definition

The Singleton Pattern ensures that a class has only one instance, and
providers a global point of access to it.

A big no no is using the singleton pattern in multiple threads. This is
combatted by making the class synchronized.

This can be done in go by doing the [following](https://stackoverflow.com/questions/18880575/what-is-the-golang-equivalent-of-a-java-synchronized-block).

To improve performance, we can use "double checked locking" in getInstance() method.

## Structure

![](../data/singleton_pattern_uml.png)

## Explanation

## Applicability
