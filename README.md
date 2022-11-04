# andictl

## Overview
This is tool for generating golang  basic project with no dependency on any framework.
## Getting Started
### Installation
Installation from source 
```
git clone git@github.com:andiwork/andictl.git
go build .
rm /usr/local/bin/andictl
cp andictl /usr/local/bin/andictl
chmod +x /usr/local/bin/andictl
```

## Tools used
- Gorm
- Survey
- Cobra
- Viper
- Mockgen
- Air
- GoReleaser
- Copier : https://github.com/jinzhu/copier

## Known Issues
* [ ] Do not create ./pkg/__model__/type.go file if  exist
* [ ] Fix ctrl + c bug. Intercept the os.Signal
* [ ] Replace factory Register to FactoryRegister to avoid duplication

## Todos
* [ ] Handle Gorm mysql uuid ID
* [ ] Generate model with more fields
* [ ] Exclude tests package in generate model package list
* [ ] download air
* [ ] Update unit tests
* [ ] Cron generator
* [ ] Pack final package
* [ ] Add mage from https://magefile.org
* [ ] Add asdf from https://asdf-vm.com
* [ ] Add GetAllRegistered function in factory generator

## License

© James Kokou GAGLO, 2021~time.Now

Released under the [Apache License Version 2.0](https://www.apache.org/licenses/LICENSE-2.0.txt)
