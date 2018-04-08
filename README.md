# graph-theory
Go program to build a non-deterministic finite automiton from a regular expression, can use this NFA to check if regex matches a given string

The program accepts either an Infix or Postfix expression.  If it is infix the expression is converted to postfix first and the string entered is checked against the postfix expression.  A response of true or false is then output to the console.

Capable of dealing with the '*', '.', '|' and '+' special characters

## Running the code
To run the code in this repository, the files must first be compiled.
The [Go compiler](https://golang.org/dl/) must first be installed on your machine.
Once that is installed, the code can be compiled and run by following these steps.

1. Clone this repository using [Git](https://github.com/sshhane/graph-theory).

```bash
> git clone https://github.com/sshhane/graph-theory.git
```
2. Change into the folder.
```bash
> cd graph-theory
```
3. Compile and run with the following command:
```bash
> go run rega.go
```

### Prerequisites

The [Go compiler](https://golang.org/dl/) must first be installed on your machine.

## Sample output

```
Infix / Postfix expression to NFA conversion
===================
Enter:
1) for Infix
2) for Postfix
1
Enter Infix expression:
a.b.c
Enter string to test:
abc
String abc  matches nfa:  true
```

```
Infix / Postfix expression to NFA conversion
===================
Enter:
1) for Infix
2) for Postfix
2
Enter Postfix expression:
ab.c
Enter string to test:
sss
String sss  matches nfa:  false
```

```
Infix / Postfix expression to NFA conversion
===================
Enter:
1) for Infix
2) for Postfix
2
Enter Postfix expression:
a+b.
Enter string to test:
aaaaab
String aaaaab  matches nfa:  true
```


## Acknowledgments

* https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
* https://web.microsoftstream.com/video/946a7826-e536-4295-b050-857975162e6c
* https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
* https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d
* https://gist.github.com/PurpleBooth/109311bb0361f32d87a2

