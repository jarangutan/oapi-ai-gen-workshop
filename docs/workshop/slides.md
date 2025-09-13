---
transition: fade
colorPreset: dark
class: invert
_class: invert
paginate: true
_paginate: false
---

<!--
footer: ShellHacks 2025
-->

# Building an API in Go using Copilot and Code Gen

Jose Aranguren, ![width:26px invert](https://raw.githubusercontent.com/gilbarbara/logos/refs/heads/main/logos/github-icon.svg) jarangutan

<!-- # Agenda
1. Introduction
2. Overview
3. What are APIs
4. OpenAPI Schemas
5. Getting started with Go
6. Code Generators
7. Lets make an API!
-->

---

![bg right:50%](./assets/chester.png)

## whoami

Hi I'm Jose :-)

I like:

- Gardening
- Sourdough
- My dog --->
- Reading documentation
- Concerts

---

## The team

TODO: Fix names woops

- Yasmine Abdrabo
- William Penneker
- Fabiana Lincon TODO: Fix

We are here to help! Raise your hand if you get stuck!

---

## What are the prerequisites?

Go 1.24+ and your favorite IDE

That's it! You can build a lot with just Go

If you need Go, head to <https://go.dev/doc/install> and install for your operating system

---

## What we're going to do together

- Use OpenAI to build us an OpenAPI spec
- Use Code Generation to build us an API
- Learn some Go to code up some handlers
- Build us a database
- Run our API

You'll be coding so get your IDE open!

---

# Did you misspell OpenAI?

No!

OpenAPI is a formal standard for describing HTTP APIs

---

<!--
_footer: https://swagger.io/docs/specification/v3_0/basic-structure/
-->

In the real world, your users need to know what your API does and what it returns

```yaml
openapi: 3.0.4
info:
  title: Sample API
  description: Optional multiline or single-line description in [CommonMark](http://commonmark.org/help/) or HTML.
  version: 0.1.9
servers:
  - url: http://api.example.com/v1
  - url: http://staging-api.example.com
paths:
  /users:
    get:
      summary: Returns a list of users.
      description: Optional extended description in CommonMark or HTML.
      responses:
        "200": # status code
          description: A JSON array of user names
          content:
            application/json:
              schema:
                type: array
                items:
                  type: string
```

---

## But what is an API?

<!--
_footer: https://www.postman.com/what-is-an-api/
-->

> An API, which stands for application programming interface, is a set of protocols that enable different software components to communicate and transfer data.

---

# While we're at it, what is Go?

<!--
_footer: https://go.dev/ and https://gobyexample.com/hello-world
-->

Go is a programming language developed by Google that is simple, built for concurrency, and with a strong standard library

```go
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```

---

# But I don't know Go :-(

That's OK! While we are coding I will be referring to

<https://gobyexample.com/>

That way, you can got something you can go back to when learning more later

---

# Lets get Started :-)

```bash
git clone https://github.com/jarangutan/oapi-ai-gen-workshop.git
git checkout base
go mod tidy
```

---

# You made it :-D

Awesome job!

---

# What's next?

Checkout the code in the main repo! There are tons of notes explaining bits and pieces of the code

There's also a version of this API that uses sqlite as a database with the same server we just built

# But I want more

- Read [Go's Standard Library](https://github.com/golang/go)
- Read [Go's Blog](https://blog.golang.org/)
  - Specially on slices ([1](https://blog.golang.org/slices-intro), [2](ttps://blog.golang.org/slices) and [strings](https://blog.golang.org/strings)
- Read [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests/)
- Read others Go code on Github

---

# Most importantly

Go build!

---

# Q and A

---

# Thank you
