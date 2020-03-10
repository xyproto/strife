# Eliza vs Niall

[![Build Status](https://travis-ci.com/xyproto/strife.svg?branch=master)](https://travis-ci.com/xyproto/strife) [![License](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/xyproto/strife/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/strife)](https://goreportcard.com/report/github.com/xyproto/strife)

Niall, a pretty limited, but receptive chatbot, seeks therapy from Eliza, the old and traditional theraputic chatbot.

Weirdness ensues.

There is no purpose to this apart from seeing what happens.

Requires Go 1.11 or later.

## Usage

Be in the same directory as `niall.txt`, or supply training data on stdin, then:

    ./strife

Or supply the desired number of output lines:

    ./strife 10

## Installation

Development version:

    go get -u github.com/xyproto/strife

## Example output

```
Eliza: How do you do. Please tell me your problem.
Niall> Everything is not very friendly.
Eliza> How do you feel when you say that?
Niall> I feel terrible.
Eliza> When you feel terrible, what do you do?
Niall> My cat is not very friendly towards me.
Eliza> When your cat is not very friendly towards you, how do you feel?
Niall> In circles.
Eliza> Let's change focus a bit... Tell me about your family.
Niall> What shall i feel terrible.
Eliza> Why do you ask?
Niall> Every other thursday.
Eliza> Let's change focus a bit... Tell me about your family.
Niall> Why do you say that maybe knows me what do such a loyal dog.
Eliza> Why don't you tell me the reason why do me say that maybe knows you what do such a loyal dog?
Niall> My cat.
Eliza> Why do you say that your cat?
Niall> My parrot.
Eliza> I see, your parrot.
Niall> I feel terrible.
Eliza> When do you usually feel terrible?
Niall> Lets change focus a talkative parrot.
Eliza> How do you feel when you say that?
Niall> Then i have a run around in circles.
Eliza> Why do you tell me that you've a run around in circles?
Niall> I have problems opening doors.
Eliza> Now that you have problems opening doors, what will you do next?
Niall> Can anybody help me the next thing.
Eliza> Please tell me more.
Niall> Why dont know me.
Eliza> Why don't you tell me the reason why dont know you?
Niall> Why do you say that youve a really they all know you.
Eliza> We should be discussing you, not me.
Niall> Lets change focus a bit gray.
Eliza> I see. And what does that tell you?
Niall> Why do you usually feel terrible.
Eliza> Why do you care whether I usually feel terrible?
Niall> Please tell me.
Eliza> Can you elaborate on that?
Niall> My cat is horrible.
Eliza> When your cat is horrible, how do you feel?
...
```

## General information

* Version: 0.0.1
* License: MIT
* Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
