# ADR - Morse Translator

This document aims to describe the organization and the decisions made around the project

### Folder organization
In project planning, I wondered for myself about the coupling of functionality. So, my approach for this development was turn the translator into a lib. This way, anyone could import this routine inside of your code.

An example of how user could use the lib, is the [playground](pkg\docs\playground.md).

Following the lib patterns in Golang, ideally, the code is in `pkg` folder. It was what I did. Inside, we have some another folders:
* `domain`: Where I store the "business" routines and entities;
    * `entities`: entities that maka part of core. Hardly will change and doesnt knows the exterior world;
        * `types`: auxiliary types used around domain;
    * `translation`: The use-cases services folder. Was named "translation" to be more friendly to users and to makes sense in imports paths;
---
### Code

Here, you will find details of decisions around the code

#### Translate use-case
Starting by service folder called `translation`, we have the actual functionality. As my implementations read a dictionary from file, this layer of file read is isolated, and the mais functions of `FromMorse` and `ToMorse` just knows the entry. This approach give us some advantages:
* `Easy to change the source`: If the dictionary won't be a file anymore, it's easy to change;
* `Performance`: The load of dictionary remains the instance. It can be reutilized through the functions;

#### Types
The types folder contains structures that represents data. For example, the `dictionary_types.go` file. It represents the default path of Latin file. The user doesn't need to know the path of default dictionary when is implementing, just the type. By the way, it opens a possibility to add more types and alphabet support.

#### Tests
The coverage is 100% actually. The tests follow the pattern of pure GOlang. We can find test by `_test` suffix and any another lib for testing was used, just the original.
The tests functions names have an approach of the real case. `WHEN > action > SHOULD > return` - When user send this, should return that.
