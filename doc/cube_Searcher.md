# Searcher

#### Class: Searcher
* id self
* read file __fileReader__
* return resource

#### Description
A class to search and read the resource requested.

#### Contracts
- +searchRequested
- -placesToSearch
- -loadPlacesToSearch

#### Knowledge Required
* resourceNameRequested : __string__ (argument)
<!-- * resource : __object__ (variable) -->
* placesToSearch : __aList__ (variable)
* propertiesResourceName : __string__ (variable)

#### Message Protocol
* searchRequested(fileName string) file
* placesToSearch(path string) self
* loadPlacesToSearch() self

#### Events Generated
None.


<!-- #### Class: EngineParser
* read file <- __Searcher__
* validate file
* parse file
* return html

#### Class: Catalogue
* get html resource <- __EngineParser__
* order htmls
* return html -->
