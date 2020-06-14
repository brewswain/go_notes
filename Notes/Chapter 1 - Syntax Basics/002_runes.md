# Runes

While strings are used to represent a _series_ of text characters, Go uses **_runes_** to represent single characters.
String literals are wrapped by double quotation marks(""), but rune literals are written with singly quotation marks('').
This explains why Golang is so opinionated over when you should use double quotation marks.

Interestingly, Go programs can use almost any character from any language, as Go uses the unicode standard for storing runes.
Runes are kepts as numeric codes, and not the character themselves. This means that if you pass a rune to fmt.Println(), the numeric code itself will get outputted, not the original character.

Also, escape sequences can be used just like with string literals ('\t', etc)
