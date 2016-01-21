Greenpeace
----------

Small wrapper for Kubernetes that exposes secrets read from files undes /secrets as
environemnt variables and then executes desired binary.

It also does Kubernetes style variable expansion in your command arguments, so
variables in form `$(MY-VARIABLE)` are expanded.

Usage:
`greenpeace bash -c echo $(TEST)`


If /secrets/test exists, content of that file will be exposed as $TEST and
expanded in commands/arguments if used in form `$(TEST)`.
