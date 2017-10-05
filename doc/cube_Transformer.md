# Transformer

#### Class: Transformer
* id self
* read file __Searcher__
* return html

#### Description
A class to transform a requested resource/file in a html.

#### Contracts
- +transformRequested

#### Knowledge Required
* resourceNameRequested : __string__ (argument)
* searcher : [Searcher](./cube_Searcher.md) (variable)
* rulesToApply : __aList__ (variable)

#### Message Protocol
* transformRequested(fileName string) string

#### Events Generated
None.
