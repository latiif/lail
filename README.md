# 🌘 Lail

Lail is a [recursive acryonym](https://en.wikipedia.org/wiki/Recursive_acronym) that stands for **L**ail is **A**nother **I**nterpreted **L**anguage _and it also means [night in Arabic](https://en.wiktionary.org/wiki/%D9%84%D9%8A%D9%84#Noun)_.

## Why Lail

Lail is still in its early stages, and I am yet to decide on a use-case for it, it is just a hobby project that I started by following the amazing book [Writing An Interpreter In Go](https://interpreterbook.com/).

## Lail Philosophy

1. Lail should be minimal with an intuitive and sensible syntax and behaviour
1. Lail should provide building blocks to extend itself.
1. Built-in functions should be kept at minimal with a justifiable raison d'être.

## Language Specification

As Lail is still under development, the language specification _can_ change.

### Program

A Lail program is zero or more statements and expressions. A program's result is its last expression. Statements are separated by `;`.

![lail example](https://i.imgur.com/I0FpaMy.png)

* Functions are first class citizens.

* Dot Notation allows for chaining functions and for more readable code.

* All functions are anonymous functions.

* Assignment can be done by `let` or directly with `=`. Assignment is an expression.

* Last expression in a function is its return value.

* Identifiers can include any unicode letter plus emojis.
