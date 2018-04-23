# yang-tutorial-turingmachine

Project to study Golang.

Target: Turing Machine Implementation used in [pyang Yang Tutorial](https://github.com/mbj4668/pyang/wiki/InstanceValidation)

## Reference

* [InstanceValidation · mbj4668/pyang Wiki](https://github.com/mbj4668/pyang/wiki/InstanceValidation)
* [IETF pyang Tutorial](https://www.ietf.org/slides/slides-edu-pyang-tutorial-01.pdf) (pdf)
* [pyang/doc/tutorial at master · mbj4668/pyang · GitHub](https://github.com/mbj4668/pyang/tree/master/doc/tutorial)
* [DSDLMappingTutorial < Main < TWiki](http://www.yang-central.org/twiki/bin/view/Main/DSDLMappingTutorial)

## Build/Run
Install dependency tools at first.
* [GitHub \- favadi/protoc\-go\-inject\-tag: Inject custom tags to protobuf golang struct](https://github.com/favadi/protoc-go-inject-tag)
* [GitHub \- openconfig/goyang: YANG parser and compiler to produce Go language objects](https://github.com/openconfig/goyang)

```
$ go get github.com/openconfig/goyang
$ go get github.com/favadi/protoc-go-inject-tag
```

and make server and client
```
$ make
```

Run server
```
$ ./tm_server
```
Run client (in another terminal),  and type command like below.
```
$ ./tm_client -t turing-machi.xml -i turing-machine-rpc.xml
command: help
command: get
command: initialize
command: config
command: run
command: get
command: exit
```

Then client send gRPC message to server, server works as turing-machine.
```
# create transition function
input        | output
state symbol | state symbol headmove
   S0      1 |    S0           right
   S0      0 |    S1      1         
   S1      1 |    S1            left
   S1      0 |    S2           right
   S2      1 |    S2      0         
   S2      0 |    S3           right
   S3      1 |    S3      0         
   S3      0 |    S4           right

[...]

# Run
Step State | Tape                           | Next Write Move
   1  [S0] |  0 |<1>| 1 | 1 | 0 | 1 | 1 | 0 | [S0]         =>
   2  [S0] |  0 | 1 |<1>| 1 | 0 | 1 | 1 | 0 | [S0]         =>
   3  [S0] |  0 | 1 | 1 |<1>| 0 | 1 | 1 | 0 | [S0]         =>
   4  [S0] |  0 | 1 | 1 | 1 |<0>| 1 | 1 | 0 | [S1]     1     
   5  [S1] |  0 | 1 | 1 | 1 |<1>| 1 | 1 | 0 | [S1]        <= 
   6  [S1] |  0 | 1 | 1 |<1>| 1 | 1 | 1 | 0 | [S1]        <= 
   7  [S1] |  0 | 1 |<1>| 1 | 1 | 1 | 1 | 0 | [S1]        <= 
   8  [S1] |  0 |<1>| 1 | 1 | 1 | 1 | 1 | 0 | [S1]        <= 
   9  [S1] | <0>| 1 | 1 | 1 | 1 | 1 | 1 | 0 | [S2]         =>
  10  [S2] |  0 |<1>| 1 | 1 | 1 | 1 | 1 | 0 | [S2]     0     
  11  [S2] |  0 |<0>| 1 | 1 | 1 | 1 | 1 | 0 | [S3]         =>
  12  [S3] |  0 | 0 |<1>| 1 | 1 | 1 | 1 | 0 | [S3]     0     
  13  [S3] |  0 | 0 |<0>| 1 | 1 | 1 | 1 | 0 | [S4]         =>
  14  [S4] |  0 | 0 | 0 |<1>| 1 | 1 | 1 | 0 | END

[...]
```
