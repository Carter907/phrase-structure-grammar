## My Phrase Structure Grammar

> this project arose from reading about grammars in my discrete math textbook. I was interested in the idea of creating my own phrase structure grammar.
> I'm not sure how this project will turn out, but it's worth a shot!

### Definitions
*Vocabulary*
> A vocabulary (or alphabet) V is a finite, nonempty set of elements called symbols. A word
(or sentence) over V is a string of finite length of elements of V. The empty string or null
string, denoted by 𝜆 (and sometimes by 𝜖), is the string containing no symbols. The set of all
words over V is denoted by V∗. A language over V is a subset of V∗.

*Phrase Structure Grammar*
> A phrase-structure grammar G = (V, T, S, P) consists of a vocabulary V, a subset T
of V consisting of terminal symbols, a start symbol S from V, and a finite set of productions P. The set V − T is denoted by N. Elements of N are called nonterminal symbols. Every
production in P must contain at least one nonterminal on its left side.
