This repo is a quick way to test Go's build constraints for conditional compilation.

### Examine the first line of foo.go, bar.go, and default.go.

You'll see a build constraint at the top of each file.  Read [Build Constraints documentation](http://golang.org/pkg/go/build/#hdr-Build_Constraints).

### Play with build tags

You can specify arbitrary build tags.  Read [how to specify build tags](http://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies) with the go command.

Try the following:

    % go build -tags 'foo'
    % ./tagtest
    
This should only compile foo.bar.

    % go build -tags 'bar'
    % ./tagtest
    % go build
    % ./tagtest
    % go build -tags 'foo bar'
    
The last one will give you an error since you are trying to compile two hello() in package main.

