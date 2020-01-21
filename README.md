# Eliza vs Niall

[![Build Status](https://travis-ci.org/xyproto/strife.svg?branch=master)](https://travis-ci.org/xyproto/strife) [![License](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/xyproto/strife/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/strife)](https://goreportcard.com/report/github.com/xyproto/strife)

Let the Eliza chatbot talk with the Niall chatbot, where Niall is trained by the given data on `stdin`, and watch the ensuing conversation.

## Usage

Be in the same directory as `niall.txt`, or supply training data on stdin, then:

    ./strife

Or supply the desired number of output lines:

    ./strife 10

## Installation

Requires Go 1.11 or later.

Development version:

    go get -u github.com/xyproto/strife

## Example output

```
Eliza: How do you do. Please tell me your problem.
Niall> Once a loyal dog.
Eliza> How do you feel when you say that?
Niall> I have a good relationship with my cat.
Eliza> When your cat, how do you feel?
Niall> Every other thursday.
Eliza> I see.
Niall> Every other thursday.
Eliza> Let's change focus a bit... Tell me about your family.
Niall> I feel when your cat.
Eliza> When do you usually feel when my cat?
Niall> I feel when my cat.
Eliza> Do you often feel when your cat?
Niall> How do you feel terrible.
Eliza> Why do you care whether I feel terrible?
Niall> In circles.
Eliza> Please tell me more.
Niall> How do is the first day.
Eliza> How do you suppose?
Niall> Once a loyal dog.
Eliza> How does that make you feel?
Niall> The time flies.
Eliza> How do you feel when you say that?
Niall> The time flies.
Eliza> Can you elaborate on that?
Niall> I feel when you are you feel when you often feel when you are you suppose.
Eliza> Do you often feel when me am me feel when me often feel when me am me suppose?
Niall> Now is wonderful.
Eliza> Please tell me more.
Niall> Now is to do you feel when you suppose.
Eliza> Why do you care whether I feel when me suppose?
Niall> Then i want to go for a very friendly cat.
Eliza> If you got to go for a very friendly cat, then what would you do?
Niall> I have problems opening doors.
Eliza> Now that you have problems opening doors, what will you do next?
Niall> I would you are you do you have problems opening doors.
Eliza> Does it please you to think that I'm me do me have problems opening doors?
Niall> Lets change the topic.
Eliza> Can you elaborate on that?
Niall> In a very friendly cat is a loyal dog.
Eliza> How do you feel when you say that?
```

## General information

* Version: 0.0.1
* License: MIT
* Alexander F. Rødseth <xyproto@archlinux.org>
