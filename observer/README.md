# The Observer Pattern

## Chapters:

 - [General Information](#general-information)
 - [Explanation](#explanation)
 - [Structure](#structure)
 - [Applicability](#applicability)

## General Information

 * *Category*: Behavioral Pattern. 
 * *Intent*: Define a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically. 
 * *Other Names*: Dependents, Publish-Subscribe.

## Structure

![](../images/observer_pattern_uml.png) 

## Terminology

| Term                     | Meaning                                                                                                                                         |
| ------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| One-to-many Relationship | "A relationship that includes two entities A and B, where A may be linked to many elements of B, but B can only be linked to one element of A." | 
| Open-closed Principle    | "An entity can allow its behaviour to be extended without modifying its source code."                                                           |
| Subject                  | Provides an interface for attaching, detaching and notifying observer objects.  Has a list of observers that have registered to it.             |
| Concrete Subject         | Implements the Subject interface. Stores states that are interesting to the observers.                                                          |
| Observer                 | Provides an interface for Observers that should be notified when of changes in a subject.                                                       |
| Concrete Observer        | Implements the Observer interface. Maintains references to a single ConcreteSubject. Stores state that should stay consistent with the subject. |


## Explanation

The motivation behind the creation of the Observer pattern was to describe a
one-to-many relationship between objects that follows the open-closed principle.

An example of the Observer pattern being used would be spreadsheet software. 
When updating the information in the cells of a spreadsheet, anything that depends on the information (charts, images) will need to be updated.
The subject (the spreadsheet) will notify its observers (charts, images) when information changes, and they will update independently of each other. 

## Applicability 

The general usecase of the Observer pattern can be fit into three situations (NOTE: this is not comprehensive):

1. An abstraction that has two aspects, and one is dependent on the other. Encapsulating these in the Observer pattern allows reusability.
2. An object needs to be able to notify other objects, but not make any assumptions about who those objects are.
3. Making a change to one object requires changing other objects, and it is unknown how many objects need to be changed.

## Consequences

The subject only knows that it has observers, but does not know anything about the functions that the observers may do.

Because they are 'loosely coupled' they can belong to different layers of a system. 
A low level subject can communicate with a high level observer, and vice-versa.

Since observers of a subject have no knowledge of each other, they do not know the cost of changing subject.
