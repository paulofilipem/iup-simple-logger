# iUP Simple Logger - GoLang Version

Simple library that implements log writing to files

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Example

Create a file example.go

```
package main

import logger "github.com/paulofilipem/iup-simple-logger"

func main(){

    //Optional param, by default will created a file on runtime path with name "application.log"
    logger.SetFileLog("/tmp/logfilename.log")

    //Optional param, by default selected level is "DEBUG"
    logger.LogSelectedLevel = logger.INFO

    logger.Info("Test")
    logger.Debug("My debug message")
    logger.Error("Other test")

}
```

And run

```
go run example.go
```

## Contributing

Feel free :D

## Authors

* **Paulo Filipe Macedo @paulofilipem** 

## License

This project is licensed under the MIT License - see the [LICENSE.md](http://en.wikipedia.org/wiki/MIT_License) file for details
