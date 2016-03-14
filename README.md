gen-yml
=======
Yaml generator from json.


Dependencies
------------

    $ go get github.com/codegangsta/cli
    $ go get github.com/codenamoo/typeconv
    $ go get gopkg.in/yaml.v2

or

    $ godep restore


Build
-----

    $ go build main.go

or

    $ make build

or

    $ ./build


Example
-------

    $ bin/gen-yml -d '{"foo":"bar"}'
    foo: bar

