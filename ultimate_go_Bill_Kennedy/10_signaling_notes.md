# Section 10: Channels, Signaling  and Orchestration in Go

## Questions to ask

Data -- Does the signal include data?

* YES => Do we need guaranteed delivery?
  * YES => use **unbuffered** **channel** (note: The price we pay for guaranteed delivery is *latency*)
  * NO => use **buffered** **channel** with buffer size > 1
  * Yes, but delayed guarantee is OK => use **buffered channel with buffer size = 1**
* NO => either buffered or unbuffered channels may be used (e.g., a "cancel" signal)
  * First choice => context
  * Second choice => unbuffered
  * Suspicious choice => buffered

## Channel States

The state of a channel can be `nil`, `open` or `closed`

Behavior of send and receive operations on a channel depends on the channel's state:

| Operation | NIL    | OPEN    | CLOSED      |
| --------- | ------ | ------- | ----------- |
| Send      | Blocks | Allowed | ***PANIC*** |
| Receive   | Blocks | Allowed | Allowed     |

Notes on channels

* **Close** is a state change, no a resource clean up operation
* Channels are always created in an **open** state
* Once a channel is **closed**, it cannot be re-opened

## Channel Pattern #1 -- Wait for Task

## Channel Pattern #2 -- Wait for Result

## Channel Pattern #3 -- Wait for Finish

