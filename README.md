# FlexJSON 

An example of how to deal gracefully with dynamic JSON field values.

When parsing third-party JSON, it is sometimes the case that field values can be of multiple types, depending on the context. This is bad data design, but it does occur in the wild. This repo shows a method of dealing with dynamic typing during JSON marshalling and unmarshalling. 

The TL;DR is to us a wrapper object for the sub-objects in your JSON that may be of multiple types. 

During unmarshalling, the raw data can be examined to determine the type of the field value. The wrapper object is responsible for storing multiple types. In this example, these types are either an integer ID for an object, or the object itself. This is a common case.

In order to marshal Go objects back into JSON, the wrapper also needs to store information about how it was originally unmarshalled. In this example, the `partial` flag indicates whether the whole object or only an ID was encountered in the original JSON. The custom marshal method outputs the correct data value in each case.

## Installation

```
go get github.com/neocortical/flexjson
cd ~/path/to/flexjson && go install .
```

## Running the Example

```
flexjson
```

The main function reads each example string in flexjson/examples.go, unmarshals to Go objects, reports the results, remarshals to JSON, and prints the output to the screen.
