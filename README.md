

#go-klvgen

Note: This is a port of [github.com/kevana/klvgen](http://github.com/kevana/klvgen) written in Go.

This is a proof-of-concept tool that generates a MISB 601.2 compliant metadata packet stream for interpolation with video streams. It includes a running timestamp and several static data fields.

##Typical use case 
We want information such as the timestamp and gps coordinates of a drone's camera feed, but we don't want to worry about timestamps getting out of sync with the video or matching  metadata to the right video stream.

##Compilation

    go build

##Usage
If run with no arguments, klvgen will start with default values for connection parameters and fields.

    $ ./klvgen

For a full listing of available parameters, run klvgen with the `-h` or `--help` flags.

    $ ./klvgen -h
    $ ./klvgen --help
