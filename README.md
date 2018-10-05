# go-minruby
MiniRuby interpreter written in Go

# Usage

```bash
$ make
2018/10/05 17:14:27 Hello kenju-wagatsuma! Welcome to go-minruby!
(=･ω･=)>> 1 + 1
2
(=･ω･=)>> p(100 * 33)
3300
(=･ω･=)>> true
true
(=･ω･=)>> !false
true
(=･ω･=)>> !(10 > 5)
false
(=･ω･=)>> (5 + 10 * 2 + 15 / 3) * 2 + -10
50
(=･ω･=)>> x = 5
(=･ω･=)>> y = x
(=･ω･=)>> x + y
10
(=･ω･=)>> msg = "hello"
(=･ω･=)>> msg
hello
(=･ω･=)>> if true; 10; end
10
(=･ω･=)>> if (1 > 3); 10; else 20; end
20
(=･ω･=)>> name = "Mr. Foo"
(=･ω･=)>> msg = "Hello, " + name + "!"
(=･ω･=)>> msg
Hello, Mr. Foo!
```

# References

- https://github.com/mame/cookpad-hackarade-minruby
- https://github.com/kenju/go-monkey
