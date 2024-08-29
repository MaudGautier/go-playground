# Go playground

This project is a small playground for me to experiment and better understand some behaviors of Go.

## Getting started

```bash
go run <script_name.go>
```

## Experiments

Explanations are written directly in test scripts:

- `append_slice` - Shows that `append`s on slices copies the original slice when the capacity equals the length, and
  mutates it in place otherwise.
