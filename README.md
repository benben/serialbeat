# serialbeat

A beat for shipping serial output to logstash or elasticsearch. Based on [`serialbeat`](https://github.com/suda/serialbeat) by Benjamin Knofe but updated to `beats/v7`.

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

### Build

To build the binary for serialbeat run the command below. This will generate a binary
in the same directory with the name serialbeat.

```
make
```

Now you should have a `serialbeat` binary in the same directory.

## Configuration

See `serialbeat.reference.yml` for all possible configuration options.

## Development

### Init Project
To get running with serialbeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push serialbeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/suda/serialbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Test

To test serialbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean serialbeat source code, run the following command:

```
make fmt
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone serialbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/suda/serialbeat
git clone https://github.com/suda/serialbeat ${GOPATH}/src/github.com/suda/serialbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.


## License

See LICENSE
