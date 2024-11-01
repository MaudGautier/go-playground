# Go playground

This project is a small playground for me to experiment and better understand some behaviors of Go.

## Getting started

```bash
go run <script_name.go>
```

## Experiments

Explanations are written directly in test scripts:

- [`append_slice`](./experiments/append_slice.go) – Shows that `append`s on slices copies the original slice when the
  capacity equals the length, and mutates it in place otherwise.
- [`strings_and_runes`](./experiments/strings_and_runes.go) – Shows how strings are parsed: parsed as Unicode code
  points returned as runes, and not ASCII letters. However, when using `strings.Fields`, we get actual strings.
- [`channels_communication_order`](./experiments/channels_communication_order.go) – Shows that the ordering in channel
  communication does not correspond to order in which goroutines were launched, but to the order in which they finish
  their execution and send it to the channel.
- [`unbuffered_channel_waits_for_receiver_to_send`](./experiments/unbuffered_channel_waits_for_receiver_to_send.go) –
  Shows that nothing is sent to a channel until there is a receiver ready to receive it.
- [`wait_group`](./experiments/wait_group.go) – Shows that children goroutines are killed as soon as the main routine
  completes (even if the children goroutines had not finished executing) and that this can be prevented with a wait
  group.
- [`channel_deadlock`](./experiments/channel_deadlock.go) – Shows a deadlock when the receiver is in the same routine as
  the sender (if the number of values sent to the channel is greater than its buffer size).
