# Command Pattern

Loosely based on the [GoF](https://en.wikipedia.org/wiki/Design_Patterns) [Command Pattern](https://en.wikipedia.org/wiki/Command_pattern).

* `Command`: a function that takes untyped string arguments and returns an error. e.g. a CLI tool entry point.
* `Registry`: a set of commands that can be invoked by name.
* `Runner`: executes a `Command` in a `Registry` with optional instrumentation such as Prometheus metrics.