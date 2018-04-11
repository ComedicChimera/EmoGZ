# EmoGZ
*WIP*

A simple, esoteric, Emoji Based language.  EmoGZ (pronounced Emojis) was not made for practical use, but rather as an experiment to test limits of language design.

The EmoGZ interpreter is written in pure Go and uses no external libraries. 

### Example of a Hello World Program
```
💎 'h' 'e' 'l' 'l' 'o' ' ' 'w' 'o' 'r' 'l' 'd'
💯 10 🔫
🍊
  💎👀
  🍦 1
  😂
  🍦 0
  😺
  👌🍌🏁
🏁
```

## Installation

You can download it as a small compressed binary or as a buildable source.

### Windows
 - x86:
 - x64:

## Building EmoGZ

All EmoGZ files have the file extension `.ejz`.

To run EmoGZ code, simply pass the file you wish to run as a starting argument
to the **emo** binary.

```
emo hello.ejz
```

## EmoGZ Language Reference

To begin, we recommend using some form an emoji keyboard if you are actually
planning on writing EmoGZ.  It should also be noted that
EmoGZ is in not what one might describe as the most 'user friendly' language and
was designed for fun, not really as a full language.  Despite this, it
can be used to perform a surprisingly large number of functions even if it is
just a meme.

### Memory

EmoGZ's use of memory differs from normal language standards.  Firstly, there is only one
slot of 'working memory' called the **operand**.  However, there is also a region of
memory called the **cache** in which memory values can be stored and removed when they are not being operated upon.  *The cache starts are 0*

Neither the cache nor the operand is typed and for all intents and purposes the cache is
relatively infinite. (Limit: 3000 items)

## Literals

EmoGZ supports two kinds of literals: `characters` and `integers`.  Characters
were added out of necessity and for convenience, and they represent single
Unicode characters.  Integers are as they name would entail, numeric values.
But, while integers can be made negative there is not a way to enter an integer
as negative.

```c
characters: 'a' '*'
integers: 12 3453
```

## Operations

The following is a table of operations and their purposes.  This will
comprise the majority of the language reference.

| Emoji | Meaning |
| :---: | ------- |
| `😂` | Print the value in operand to the console |
| `🔫` | Change the sine of the operand |
| `💀`  | End the program |
| `💩` | Decrement operand |
| `😺` | Increment operand |
| `👌` | Evaluate the next region if the operand is greater than 0 |
| `💙` | Compare operand and literal(s) to determine if they are equal |
| `👀` | Get the operand's value |
| `💯` | Set the operand's value |
| `🍦` | Pop a value at a specific location in the cache into the operand |
| `💎` | Add a value or set of values to the cache |
| `🍊` | Begin a loop |
| `🍌` | Break a loop |
| `🏁` | End a region |
| `👑` | Get a single character of input |

*Note: All number specific operators will coerce characters to integers.
