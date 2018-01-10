# serialbeat

A beat for shipping serial output to logstash or elasticsearch.

## Synopsis

Create a configuration file with the name serialbeat.yml and add

    serialbeat:
      device: /dev/ttyACM0
      baud: 38400
      delimiter: "\n"
    output.console:
      pretty: true

And run `./serialbeat -c serialbeat.yml -e -d "*"` to see your serial output.

## Installation

    mkdir -p ${GOPATH}/src/github.com/benben/serialbeat
    git clone git@github.com:benben/serialbeat.git ${GOPATH}/src/github.com/benben/serialbeat
    cd ${GOPATH}/src/github.com/benben/serialbeat
    glide i
    make

Now you should have a `serialbeat` binary in the same directory.

## Configuration

See `serialbeat.reference.yml` for all possible configuration options.

## Contributing

Check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

## License

See LICENSE
