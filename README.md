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

A Lail program is zero or more statements and expressions. A program's result is it's last expression. Statements are separated by `;`.

```
let lang = "fr";
let greeting = if (lang == "en") {
    "Hello"
} else if (lang == "fr") {
    "Salut"
} else {
    "Im speechless"
}

greeting + " Lail";
```
