# record terminal session for later callback

* uses asciinema player, but doesn't need webservice to record, no data upload
* records into `term.rec`, `term-stdout.json` and `term.html` in current
  directory

# Building

* requires go and libtsm-dev
* `go build`

# Run

If not given a command as paramter it will run /bin/sh

Example: `./go-termrecording ls --color`
