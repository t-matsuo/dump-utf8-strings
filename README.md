# dump-utf8-strings
create hex dump from utf-8 text file with multibytes char

# Install

```
# wget https://github.com/t-matsuo/dump-utf8-strings/releases/download/0.1/dump-utf8-strings
# chmod 755 dump-utf8-strings
```

# Usage

```
# cat textfile
abcdあいうえお
efghかきくけこ

# ./dump-utf8-strings textfile
61 a
62 b
63 c
64 d
e3 [3>]
81 [->]
82 あ
e3 [3>]
81 [->]
84 い
e3 [3>]
81 [->]
86 う
e3 [3>]
81 [->]
88 え
e3 [3>]
81 [->]
8a お
0a [LF] (^J)
65 e
66 f
67 g
68 h
e3 [3>]
81 [->]
8b か
e3 [3>]
81 [->]
8d き
e3 [3>]
81 [->]
8f く
e3 [3>]
81 [->]
91 け
e3 [3>]
81 [->]
93 こ
0a [LF] (^J)
```
