# Eliza vs Niall

[![Build Status](https://travis-ci.org/xyproto/eliza_vs_niall.svg?branch=master)](https://travis-ci.org/xyproto/eliza_vs_niall) [![License](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/xyproto/eliza_vs_niall/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/eliza_vs_niall)](https://goreportcard.com/report/github.com/xyproto/eliza_vs_niall)

Let the Eliza chatbot talk with the Niall chatbot, where Niall is trained by the given data on `stdin`, and watch the ensuing conversation.

## Usage

Train Niall with the contents of `input.txt` and output 10 lines of conversation:

    eliza_vs_niall 10 < input.txt

## Installation

Development version:

    go get -u github.com/xyproto/eliza_vs_niall

## Example output

Using the GPL 3 license as input data:

```
Eliza: How do you do. Please tell me your problem.
Niall: User product.
Eliza: How do you feel when you say that?
Niall: Then your rights from the extent that there is transferred to operate with this license.
Eliza: How do you feel when you say that?
Niall: You need not include claims.
Eliza: Why do you care whether I need not include claims?
Niall: The same work as a covered work.
Eliza: Let's change focus a bit... Tell me about your family.
Niall: Or.
Eliza: How does that make you feel?
Niall: And prominently visible.
Eliza: Let's change focus a bit... Tell me about your family.
Niall: Subject to find the contributors essential patent license.
Eliza: Very interesting.
Niall: Selling its corresponding source code saying where the parties with the terms and.
Eliza: Please tell me more.
Niall: Of it includes a prominent item in connection with.
Eliza: Very interesting.
Niall: However.
Eliza: How does that make you feel?
```

Using a poem about relationships as input data:

```
Eliza: How do you do. Please tell me your problem.
Niall: Please tell me your blue eyes.
Eliza: OK, but can you elaborate a bit?
Niall: Tips dance linger to have and night.
Eliza: How do you feel when you say that?
Niall: Into the rising sunrise.
Eliza: I see.
Niall: Tips dance linger to grow old.
Eliza: Let's change focus a bit... Tell me about your family.
Niall: For you i love to infinity.
Eliza: We should be discussing you, not me.
Niall: Please tell me your problem.
Eliza: Let's change focus a bit... Tell me about your family.
Niall: Heartbeats climbs a boiling point.
Eliza: Very interesting.
Niall: Only ever with you.
Eliza: Please tell me more.
Niall: Yet.
Eliza: Very interesting.
Niall: Only ever with you my muse from divinity.
Eliza: Why do you care whether I your muse from divinity?
```

## General information

* Version: 0.0.1
* License: MIT
* Alexander F. Rødseth <xyproto@archlinux.org>
