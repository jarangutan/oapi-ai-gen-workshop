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

TODO: Put the lil guy gopher somewhere in the powerpoint where it would be cutest

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

Go 1.24+, git and your favorite IDE

That's it! You can build a lot with just Go

TODO: nicer
If you need Go, head to <https://go.dev/doc/install> and install for your operating system
If you need git, head to <https://git-scm.com/downloads> and install for your operating system

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

<!--
TODO: Clean up the text here and add more context in the slide itself :-)

- In the real world, your users need to know what your API does and what it returns
- Designing API first gets you thinking of what it is you are trying to do
- You can hand this to your team members and they'll know what you're up to
-->

---

<!--
_footer: https://swagger.io/docs/specification/v3_0/basic-structure/
-->

## Example

```yaml
openapi: 3.0.4
info:
  title: Sample API
  description: Optional multiline or single-line description in [CommonMark](http://commonmark.org/help/) or HTML.
  version: 0.1.9
servers:
  - url: http://api.example.com/v1
paths:
  /users:
    get:
      summary: Returns a list of users.
      description: Optional extended description in CommonMark or HTML.
      operationId: getUsers
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

<!--
TODO: Example of an interface
Your website needs to talk to a database. You wouldn't want your website to talk to the database directly.
An API sits between your Frontend and the database to facilitate them talking to each other. This API is where your "backend" starts

An API isn't just an HTTP API, an interface for a library can also be an API.

TODO: clean up definitions
An API (Application Programming Interface) is a set of rules that lets different software programs communicate and exchange data.

An API is a way for one piece of software to talk to another by following agreed-upon rules for sharing information.

TODO:
good sites
https://github.com/resources/articles/software-development/what-is-an-api

https://www.geeksforgeeks.org/software-testing/what-is-an-api/ (The picture is good so consider nabbing it)

-->

TODO:

> APIs act as bridges between different pieces of software, enabling them to communicate, share data, and work together.

---

# While we're at it, what is Go?

<!--
_footer: https://go.dev/ and https://gobyexample.com/hello-world
-->

<!--
TODO: less jank
Explain what the hek a package main is
Normally you start you app with a main.go file with the package main as package main is what Go looks for to run your program
Programs start running in package main.

-->

TODO: Put gopher here or next slide

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

TODO: Don't actually use gobyexample but instead let them know this site exists. Maybe mention the keywords they can use to reference back to

That's OK! While we are coding I will be referring to

<https://gobyexample.com/>

That way, you can got something you can go back to when learning more later

---

# Lets get Started :-)

TODO: Nicer
TODO: Mention if you get stuck or lost, that's ok, you will still have working code at the end. (main branch)
Go to the project repo: qrco.de/shgows (link to the repo)

```bash
git clone https://github.com/jarangutan/oapi-ai-gen-workshop.git
git checkout base
go mod tidy
```

---

# We made it :-D

Awesome job!

TODO:
If you got stuck or couldn't finish, that's ok! commit your changes and checkout main and do a make watch or go run cmd/api/main.go

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

# Thank you o/
