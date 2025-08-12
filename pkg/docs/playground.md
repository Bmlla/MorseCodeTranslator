# Playground
To library test and debug, we can generate a binary of implementation. In the root path, we can find a `main.go` with implementation of library. And, we can find a `makefile`

### Generating binary
The binary changes between Operation Systems. So, for play the app, you can run directly from IDE or generate a binary for execution.

```sh
// In root project folder
make build
```
It will generate a file for execution in the same folder as the `makefile` has been executed   

---
### Executing
For execution, we can use 3 params:
* --text: Text to be translated;
* --mode: `from` or `to`. "From" means "translate to Morse" and vice-versa;
* --dict: [optional] Custom dictionary path;

Running with default parameters:
```sh
make run-demo-from

//output:
//"Running with text='HELLO WORLD in Morse Code', dict='default/latin', mode='from'"
//HELLO WORLD
```
```sh
make run-demo-to

//output:
//"Running with text='HELLO WORLD', dict='default/latin', mode='to'"
//.... . .-.. .-.. ---  .-- --- .-. .-.. -..
```
---
### Running out of Makefile
To run out of makefile commands, just:
```sh
// Here, I am in the same folder of binary
.\MorseCodeTranslator.exe --text ".... . .-.. .-.. ---  .-- --- .-. .-.. -.." --mode from --dict "optional/path/to/your/dict"

//output
//HELLO WORLD
```
```sh
// Here, I am in the same folder of binary
.\MorseCodeTranslator.exe --text "HELLO WORLD" --mode to --dict "optional/path/to/your/dict"

//output
//.... . .-.. .-.. ---  .-- --- .-. .-.. -..
```
