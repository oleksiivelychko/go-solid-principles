# go-solid-principles

📌 **Single Responsibility** _(a class (module) should have one, and only one, reason to change, i.e. every class should have only one responsibility)_:
- [wrong way](srp/bad/main.go)
- [right way](srp/good/main.go)

📌 **Open-Closed** _(software entities should be open for extension, but closed for modification)_:
- [wrong way](ocp/bad/main.go)
- [right way](ocp/good/main.go)

📌 **Liskov Substitution** _(functions that use pointers or references to base classes must be able to use objects of derived classes without knowing it)_:
- [wrong way](lsp/bad/main.go)
- [right way](lsp/good/main.go)

📌 **Interface Segregation** _(clients should not be forced to depend upon interfaces that they do not use)_:
- [wrong way](isp/bad/main.go)
- [right way](isp/good/main.go)

📌 **Dependency Inversion** _(depend upon abstractions, not concretions)_:
- [wrong way](dip/bad/main.go)
- [right way](dip/good/main.go)

---
※ References:
- [SOLID](https://en.wikipedia.org/wiki/SOLID)
- [Understanding SOLID Principles in Golang: A Guide with Examples](https://medium.com/@vishal/understanding-solid-principles-in-golang-a-guide-with-examples-f887172782a3)
