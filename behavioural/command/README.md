# The Command Pattern

## Intent

Encapsulate a request as an object, thereby letting you parameterize clients
with different requests, queue or log requests, and support undoable oeprations.

## Structure

![](../../resources/command_pattern_uml.png)

## Terminology

| Term            | Meaning                                                         |
| --------------- | --------------------------------------------------------------- |
| Caller          | Asks a command to carry out the request.                        |
| Command         | An interface for executing an operation.                        |
| Receiver        | An interface that knows how to perform operations.              |
| ConcreteCommand | Defines a relationship between a Reciever object and an action. |

## Explanation

Sometimes we issue requests as objects without knowing anything about the operation being requested.
The command pattern provides a layer of abstraction where the caller does not
know how involved the command object's execute() function is, but can call is regardless.

The Command pattern example could be using a menu in a text document,
supporting `open()`, `copy()`, and `paste()`.  Each of these commands have
different complications (open creates a file prompt then opens the file, paste
just get what is in the clipboard and inserts it into the file). The menu can
call these all the same.

## Applicability

The usecases of the Command pattern can be fit into a few usecases:
 * Create, queue, and execute requests at different times. 
 * Support undo-ing: Executing operation can store a state, and can reverse itself.
 * Transactions: Encapsulating a request as an object is perfect for representing transactions.
