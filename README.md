# muscletail

Monitor files for occurences of specific strings over time. It's basically:

```bash
tail -F <file> | grep --line-buffered -F <string> | while read LINE; do
    ...
done
```

except it allows you to specify a threshold to compare against a [simple moving average](https://en.wikipedia.org/wiki/Moving_average#Simple_moving_average) with a customizable window length.

## Installation

So far only `go get github.com/lillesvin/muscletail`. Binaries will likely follow when this thing is a bit more mature.

## Usage

If you are ok with "OMGWTF?!" or "OHNO!" appearing every once in a while in a file but want to be alerted if they occur more frequently than every 5 seconds (over a window of 10 samples), then you configure a watch a little like this:

```toml
[[watch]]
id = "Too much OMGWTF?!"
file = "watch/this/logfile"
matches = [
    'OMGWTF?!',
    'OHNO!',
]
threshold = 5
window_length = 10
```

## Todo

 - [ ] Support defining trigger actions in the config per watch, it just logs now
 - [ ] Support regex in matches
 - [ ] Allow for more parameters than just "interval between occurences" to compare against
 - [ ] More analyses than just "simple moving average"

## Known Issue(s)

### It doesn't actually do anything on Windows

For doing the actual tailing I'm relying on [hpcloud/tail](https://github.com/hpcloud/tail) which unfortunately has a couple of [unresolved issues with Windows](https://github.com/hpcloud/tail/labels/Windows) and no one to fix them, so if you (or someone you know) can help consider lending them a hand.

## Run tests

`go test`
