# Event Bus

A multi-provider, single-subscriber callback/events dispatcher.

## What's It For?

Under most circumstances, you wouldn't want or need this. It's not optimal - nor does it intend to be such - though it *is* thread-safe.

## Why Not Use Channels?

Channels are nice and all, but they're a bit heavy-weight for what this was intended to be used for. They also require more setup and management. For a full-featured and robust solution, a system based on them makes a lot more sense. This did not intend to be full-featured, nor robust.

## What's This Package Good For, Then?

Some testing apparati need a way to dictate asynchronous operations of simulated hardware. This just happened to be a simple solution for that.
