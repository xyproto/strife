# Eliza vs Niall

[![Build Status](https://travis-ci.org/xyproto/eliza_vs_niall.svg?branch=master)](https://travis-ci.org/xyproto/eliza_vs_niall) [![GoDoc](https://godoc.org/github.com/xyproto/eliza_vs_niall?status.svg)](http://godoc.org/github.com/xyproto/eliza_vs_niall) [![License](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/xyproto/eliza_vs_niall/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/eliza_vs_niall)](https://goreportcard.com/report/github.com/xyproto/eliza_vs_niall)

Two chatbots face each other. There can only be one. Eliza meets Niall.

## Installation

Development version:

    go get -u github.com/xyproto/eliza_vs_niall

## Usage

For letting each bot say 10 lines of text, and training Niall with the text in input.txt:

    eliza_vs_niall 10 < input.txt

## Example output

With no training text for Niall, except the output from Eliza:

```
Eliza: How do you do. Please tell me your problem.
Niall: Please tell me your problem.
Eliza: How do you feel when you say that?
Niall: How do.
Eliza: What is it you're really asking?
Niall: Please tell me your problem.
Eliza: I see.
Niall: How do you do.
Eliza: Perhaps you can answer your own question.
Niall: How do.
Eliza: How do you suppose?
Niall: How do you do.
Eliza: Perhaps you can answer your own question.
Niall: How do you do you do you do.
Eliza: What is it you're really asking?
Niall: Please tell me your problem.
Eliza: Please tell me more.
Niall: How do.
Eliza: How do you suppose?
Niall: How do you do.
Eliza: What is it you're really asking?
```

With Niall being given a poem about relationships as the initial input:

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

* Version: 0.1
* License: MIT
* Alexander F. Rødseth <xyproto@archlinux.org>
