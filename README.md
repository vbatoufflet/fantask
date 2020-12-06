# fantask: tasks group helper [![PkgGoDev][pkggo-badge]][pkggo-url]

Concurrent tasks controller for Go.

## What's for?

This tasks controller runs multiple tasks concurrently then wait for either
cancellation or one of the task to stop. If any of the conditions is met,
the main tasks execution context is canceled, allowing the remaining tasks to
stop.

[pkggo-badge]: https://pkg.go.dev/badge/batou.dev/fantask
[pkggo-url]: https://pkg.go.dev/batou.dev/fantask
