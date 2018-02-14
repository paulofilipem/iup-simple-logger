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

    logger.Info("Test")

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

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
