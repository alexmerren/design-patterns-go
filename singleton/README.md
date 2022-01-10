# The Singleton Pattern

## Intent

Ensure a class only has one instance, and provide a global point of access to it.

## Structure

![](../data/singleton_pattern_uml.png)

## Terminology

| Term      | Meaning                                                     |
| --------- | ----------------------------------------------------------- |
| Singleton | Creates a unique instance that clients can access globally. |


## Explanation

Sometimes we have classes that we only need one of (i.e. database connections,
file system) The singleton pattern provides a single point of access for
accessing and creating an instance of that class. 

## Caveats 

A big no no is using the singleton pattern in multiple threads. This is
combatted by making the class synchronized.

To improve performance, we can use "double checked
locking" in getInstance() method.

## Applicability

We can use the singleton pattern in these usecases:

 * Must be only one instance of a class, and have a single access point.
 * A singleton is extensible so many subclasses can each have a single instance.
