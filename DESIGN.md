
## Variables

Variables in Lang can either be lists or bytes/numbers however, support for 32 bit numbers will be removed after [Feb. 2020](https://bitcoinsv.io/2019/04/17/the-roadmap-to-genesis-part-1/). It is strongly, dynamically typed.

  

#### Variable Def

For normal numbers or byte strings declarations using strings is fine.

```lisp

(defvar a "joe is cool")

```

```lisp

(defvar b "joe is not cool")

```

Lists can also be declared via the list function with mixed types.

```lisp

(defvar c (list 1 2 3 4 5 6 "foo"))

```
### Template Variables

Template variables do not actually hold any information duing runtime but instead are placeholders for information put in post-compilation. Sometimes it maybe useful to create templates using Lang. For example, if one wanted to create an output that can be reused with different keys this would how that could be executed.

  

#### Declaration

Template variables do not need to be declared at all. Instead, they can replace a spot of any existing variable.

```

(defoutput (p2pkh $pubkeyhash))

```

Will result in the following string at compile time.

```js
"76a914_pubkeyhash_88ac"
```

If compiled as binary the `OP_NOP`, 97, is used instead of the underscore character:


```js
118 169 20 [97 112 117 98 107 101 121 104 97] 115 104 97 136 172
```


  

#### Exception for lists

Due to the complicated nature of Lang lists there is no support at the moment to support templated lists.



## Functions
Functions in Lang are defined almost exactly like that in Lisp. They serve two purposes however:
    a.) Output Creation
    b.) Evaluation


### Output Creation

### Evaluation Function

Evaluation is more for the conveniance of the programmer. This is what happens when all your variables are predefined. For example, take the function call:

```lisp
(+ 4 10)
```

Though, this is a function
