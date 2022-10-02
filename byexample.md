# Tokens

```
"abc" string
'a' char
1234.5678 number (double)
abc identifier
$ variable
$$ constant
=> function
= change/define
() parens
{} braces
[] brackets
, comma
- substract 
+ add
* multiple
/ devide
? if
: else
@ `this`
@@ import
@~ export
; semicolon
. period
! not
& bitwise and
| bitwise or
&& boolean and
|| boolean or
<= smaller or same
>= bigger or same
< smaller
> bigger
% mod
== same
!= not same
-0 null
<- return
# comment
<< print/stdout
>> input/stdin
```

# By example

## Hello world

```
<<"Hello world!";
```

## Values

```
"abc" string
'a' char
1234.5678 number 
-0 null
! false
!! true
```

## Variables

```
$abc = "abc"; # variable
$$numbers = 1234.5678; # constant
```

## If/else

```
bl1 == bl2 ? then_function() : else_function();
```

## Arrays

```
$array = ['a','b','c'];
```

## Functions

```
$cool_function = ($msg) => {
    <<$msg;
}
$cool_function("Cool message!");
$cool_function("Aanother one!");
```

## Recursive function (only loops)

```
$recursive_function = () => {
    @(); #calls itself
}
```

## Return

```
$fancy_messages = ($msg) => {
    <- $msg + "!";
}
<<$fancy_messages("Hello there"); # Hello there!
```

## Object

```
$object = {
    cool_value = "hello",
    no_value,
    null_function = () => {
        <- -0
    }
}
<<$object.cool_value;
<<$object.no_value;
<<$object.null_function();
```

## Importing packages

```
$m = @@"math";
$m.power(10,20);
```

## Exporting packages

package.ct
```
$hello_world_function = () => {
    <- "Hello world!";
}
@~$hello_world_functio;
```
main.ct
```
$function = @@"package";
<<$function();
```