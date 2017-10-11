# Searcher

#### Class: Searcher
* id self

#### Description
A class to search for a _name_ on some _folders_.

#### Contracts
- +lookFor

#### Knowledge Required
* filesFound : __aList__ (variable)
* storage : [Storage](./cube_Storage.md) (variable)
* chooser : [Chooser](./cube_Chooser.md) (variable)
* file : __string__ (variable)
* name : __string__ (argument)

#### Message Protocol
* lookFor(name string) _string_

#### Events Generated
None.