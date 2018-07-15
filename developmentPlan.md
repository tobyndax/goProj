# Design document 
This document gives an overview of the design of parts of the project not yet implemented.
# Script Importer
The script importer language should most likely be a node-based text language. 
Nodes might be a good idea due to storyboarding for games are usually describerd as nodes.
 
## Nodes
### Basic node structure
Each node begins with a $BeginNode Number statement: 
$BeginNode 12 
$EndNode 12

Nodes can be referenced by the number for instance with the $GoTo 12 statement.
Most nodes should have a $GoTo number statement.

Nodes execute from the first statement to the last sequentially, unless a GoTo statements is encountered.
The stack is in that case cleared.  

### Members
#### GoTo Num
$Advances the story to Node Num
#### Dialog 
$Dialog This is a dialog
Prints the stated text upon entering the node. Each dialog statement ends with a newline. 
Multiple $Dialog statements per node is allowed.
#### Options List
$Options Option1:Statement1;Option2:Statement2
Ex: $Options Yes:$GoTo 12;No:$Chain $Store Alive 1;$GoTo 13
#### Function
$Function func 
Executes the function func. 
Func is a inputless function. But is allowed to pull variables from the state machine. 
#### Store
$Store Key Value
Permanently store the value Value to member Key. These values can later be used by functions or dialog. 
In dialog statements the stored values can be used by typing ${Key}. For instance $Dialog ${Name}, what are you doing? Would output the users name followed by the text. 
#### Chain
$Chain Statement1;Statement2;... 
Chains multiples statements and executes them sequentially. 
There can only be one $GoTo statement and it must be the last statement. 

