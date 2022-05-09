# The Proxy Pattern

## Intent

The Proxy Pattern provides a surrogate or placeholder for another object to control access to it.

## Structure

![](../../resources/proxy_pattern_uml.png) 

## Terminology

| Term          | Meaning                                                                                                                   |
| ------------- | ------------------------------------------------------------------------------------------------------------------------- |
| Proxy         | Maintains a reference that lets the proxy access the real subject. Controls access to the real subject.                   |
| Subject       | Defines the common interface for RealSubject and Proxy so that a Proxy can be used anywhere a RealSubject is expected.    |
| RealSubject   | Defines the real object that the proxy represents.                                                                        |

## Explanation

Sometimes it is beneficial to hide or obfuscate the actual cost of creating or
using resources, and this is why the proxy pattern exists.  We call the
functionality we want on a proxy object then that object will control the
access and resources of the real object that we do want to access.

## Applicability 

There are a few different forms of the Proxy Pattern that can be useful:
 * _Remote Proxy_ - Provides local representation for an object in a diferent address space.
 * _Virtual Proxy_ - Creates expensive objects on demand.
 * _Protection Proxy_ - Controls access to the original object, useful for objects with different access rights.

## Notes

* As the Head First book is written in java, there is a lot of work going on in
  the background to support this, like rmiregistry and remote exceptions,
  autogenerating the stub and skeleton services that RMI uses.

* The Proxy pattern should be used to "create a representative object that
  control access to another object, which may be remote, expensive to create,
  or in need of securing."

* An example of a virtual proxy is an application that views CD covers:
    - A program will display "Loading CD cover, please wait..." while the bytes of the CV cover are being retrieved.
    - It will then remove the waiting message and then display the CD cover.
    - The proxy will show the loading CD cover message, and tell the actual object to get the cover from the internet/local file

* There's an interesting section on the similarities between the proxy pattern and the decorator pattern.

* Decorator adds behaviour to objects, and proxy represents objects - not adding any behaviour.

* A few examples of proxy patterns:
    - Firewall (control access to network resources)
    - Metrics/Smart Reference (provides actions whenever referenced)
    - Caching (provides temporary storage for expensive operations)
    - Synchronization (provides safe access from multiple threads)
    - Complexity Hiding (hides complexiy and controls access to complex set of classes)
    - Copy-On-Write (controls copying objects until they are needed by a client)
