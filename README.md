# X-buffers Bench

This repository contains differente test cases in order to assess the viability of using either JSON, Flat Buffers or Protocol Buffers as a serializastion mechanisim, the IDL files can be found [here](./model) along with the generated files

Also the test run can be found [here](./allocation/)


# Test conditions

All test are done with this structure as base

```go
type Customer struct {
  FirstName   string               
  LastName    string               
  Age         uint32               
  Balance     float64              
  Debt        float64              
  Preferences *Preferences         
  Friends     []*Customer          
  Addresses   map[string]*Location 
}
```

For each scenario is run under 3 levels of complexity:

- **Level 1:** it is the customer object with 1 address.
- **Level 2:** it is the customer object with 10 addresses, 10 friends and each friend contains 10 addresses.
- **Level 3:** it is the customer object with 100 addresses, 200 friends and each friend contains 100 addresses.

# Results

## Marshalling 

**Level 1**

```
BenchmarkAllocationNativeL1-4             354464              3442 ns/op            1075 B/op         14 allocs/op
BenchmarkAllocationProtoL1-4              467773              2188 ns/op            1088 B/op         15 allocs/op
BenchmarkAllocationFBSL1-4                697542              1690 ns/op            1304 B/op         10 allocs/op
BenchmarkMarshalGogoProtoL1-4             1244660             944.4 ns/op           627 B/op          9 allocs/op

```

![unmarshak_level1](../xbuffer/resources/img/marshal_level1.png)

**Level 2**
```
BenchmarkAllocationNativeL2-4               6504            177418 ns/op           52443 B/op        619 allocs/op
BenchmarkAllocationProtoL2-4                9921            128793 ns/op           43157 B/op        717 allocs/op
BenchmarkAllocationFBSL2-4                 10000            116492 ns/op           62473 B/op        326 allocs/op
BenchmarkMarshalGogoProtoL2-4              17931             67348 ns/op           31112 B/op        433 allocs/op
```

![unmarshak_level1](../xbuffer/resources/img/marshal_level2.png)


**Level 3**

```
BenchmarkAllocationNativeL3-4                 39          29045645 ns/op         9810394 B/op      84209 allocs/op
BenchmarkAllocationProtoL3-4                  55          21001219 ns/op         6334196 B/op     103901 allocs/op
BenchmarkAllocationFBSL3-4                    78          13158922 ns/op         8333920 B/op      42052 allocs/op
BenchmarkMarshalGogoProtoL3-4                 104         11383685 ns/op         4968503 B/op      62899 allocs/op
```

![unmarshak_level1](../xbuffer/resources/img/marshal_level3.png)

## Unmarshalling 

**Level 1**

```
BenchmarkUnmarshalnNativeL1-4             134371             10778 ns/op            1000 B/op         26 allocs/op
BenchmarkUnmarshalProtoL1-4               728599              1592 ns/op             733 B/op         13 allocs/op
BenchmarkUnmarshalFBSL1-4                 510885903                2.365 ns/op           0 B/op          0 allocs/op
BenchmarkUnmarshalGogoProtoL1-4           1863813               628.5 ns/op           528 B/op         11 allocs/op
```

![unmarshak_level1](../xbuffer/resources/img/unmarshal_level1.png)

**Level 2**

```
BenchmarkUnmarshalnNativeL2-4              2886              353497 ns/op          32025 B/op       1000 allocs/op
BenchmarkUnmarshalProtoL2-4                13471             91999 ns/op           31446 B/op        705 allocs/op
BenchmarkUnmarshalFBSL2-4                  444298354         2.490 ns/op           0 B/op            0 allocs/op
BenchmarkUnmarshalGogoProtoL2-4            34784             34081 ns/op           22206 B/op        465 allocs/op
```

![unmarshak_level1](../xbuffer/resources/img/unmarshal_level2.png)

**Level 3**

```
BenchmarkUnmarshalnNativeL3-4                 22          50498996 ns/op         5849893 B/op     148547 allocs/op
BenchmarkUnmarshalProtoL3-4                   78          14708673 ns/op         4989886 B/op     105029 allocs/op
BenchmarkUnmarshalFBSL3-4                     527564692   2.248 ns/op             0 B/op          0 allocs/op
BenchmarkUnmarshalGogoProtoL3-4               193         5784670 ns/op           3625775 B/op    64031 allocs/op
```

![unmarshak_level1](../xbuffer/resources/img/unmarshal_level3.png)
