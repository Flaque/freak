# Freak

Freak is a teeny tiny tracker that I use to track my panic attacks.

## Usage

```s
$ freak 1
$ freak 0
$ freak 10
```

## Install

```
go get -u github.com/flaque/freak
```

## What's a panic attack?

"Panic attacks" mean different things to different people, so I can only really describe my own condition. But basically it's hightened fear / paranoia that generally latches onto a particular subject.

If you've ever been in a life or death situation like a car accident, you've probably felt a huge rush of addrenaline. You probably felt a weird mix of heat and cold as your body rushed blood towards your chest and legs.

My panic attacks are a lot like that, only more/less random and not really caused by anything. They're generally followed by an intense urge to leave whatever room I'm in and a bit of general fear towards the general situation I'm in.

Because my attacks are very "fear" based, I generally make a distinction between panic and anxiety. I've never been diagnosed with general anxiety (though I have been diagnosed with panic disorder) and the experiences of folks who have do not align with my own experiences. As far as I can tell, my attacks aren't made worse by excessive stress, though having raw data may disprove that.

In general, panic attacks are just annoying, not a particular concern. I've had them all my life and I've learned how to deal with them. Unless there's a particularly bad one, most panic attacks are like an annoying itch than a crippling concern.

In my life, I've helped communicate how bad of an attack I'm having with a 0-10 scale. A 0 being not having a panic attack and a 10 being the worst panic attack I've experienced.

In general:

| level | description                             |
| ----- | --------------------------------------- |
| 0     | nothing                                 |
| 2     | brief spike, no worries                 |
| 4     | I'll probably go for a walk             |
| 7+    | I'll probably take emergency medication |

## Theories

The following are the things I'm trying to validate with this raw data.

I've suspected for some time that these attacks come in clusters: a few weeks where I have a huge amount of panic attacks and then several months where I have few or relatively minor panic attacks.

I suspect that these attacks happen at the worst possible moments, for example on a sunday night when I should get to bed so I can get up early for work or right before I'm about to go into a meeting. Though without data, I don't have a way to know if this is true or if I'm just misremembering the most annoying experiences I've had.

I suspect that these attacks are reduced when I run more, so I'd like to cross reference this with my run trackers.

## Can I use this?

Of course. It's made purposely general so it could be used for lots of purposes. You're welcome to submit PRs though I may not take them if they're not super helpful.

## Why's it called freak?

I was debating "panic" but that was too close to Golang's `panic()`.

When I was a kid, I didn't know what to call my panic attacks, so I called them freakouts. It wasn't until I got diagnosed that I called them panic.
