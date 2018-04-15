# yang-tutorial-turingmachine

Project to study Golang.

Target: Turing Machine Implementation used in [pyang Yang Tutorial](https://github.com/mbj4668/pyang/wiki/InstanceValidation)

## Reference

* [InstanceValidation · mbj4668/pyang Wiki](https://github.com/mbj4668/pyang/wiki/InstanceValidation)
* [IETF pyang Tutorial](https://www.ietf.org/slides/slides-edu-pyang-tutorial-01.pdf) (pdf)
* [pyang/doc/tutorial at master · mbj4668/pyang · GitHub](https://github.com/mbj4668/pyang/tree/master/doc/tutorial)
* [DSDLMappingTutorial < Main < TWiki](http://www.yang-central.org/twiki/bin/view/Main/DSDLMappingTutorial)

## Build/Run

```
$ go build yttm.go
$ ./yttm -t turing-machine-config.xml -i turing-machine-rpc.xml
[...]
# create transition function
input        | output
state symbol | state symbol headmove
    0      1 |     0           right
    0      0 |     1      1         
    1      1 |     1            left
    1      0 |     2           right
    2      1 |     2      0         
    2      0 |     3           right
    3      0 |     4           right
    3      1 |     3      0         
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
[...]
```
