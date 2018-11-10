# fantask: tasks group helper [![GoDoc][godoc-badge]][godoc-url]

Simple tasks group helper for Go.

## What's for?

This tasks group helper will run multiple tasks concurrently then wait for a
cancellation or one of the task to stop. If any of this condition is met,
the main context will be canceled, allowing the remainding tasks to stop.

[godoc-badge]: https://godoc.org/github.com/vbatoufflet/fantask?status.svg
[godoc-url]: https://godoc.org/github.com/vbatoufflet/fantask
