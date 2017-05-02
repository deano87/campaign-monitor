## Campaign monitor test


### Test Notes

I chose GoLang to solve the first 4 requirements.
The reason for choosing it is because Go is simple to read & understand, has great integrated unit testing, it's extremely fast & there's a lot of community support

All test scenarios are built to cover 100% of the code written for each requirement

**Requirement 1:** <br />
Please note that Go doesn't have a NULL value. It has a nil value, but it is used in a slightly different way than what is expected in the requirement, which is why i only tested the string length

#### How to run the Go tests

1. Download & install GoLang: <br />
https://golang.org/dl/<br />
It is also recommended to read the Go installation wiki: <br />
https://golang.org/doc/install

2. The default path to the go library is ~/go, make sure /src exists too:

        mkdir -p ~/go/src

3. under ~/go/src, clone this repository by running:

        git clone https://github.com/deano87/campaign-monitor.git

4. navigate to one of the requirements (1-4 are the relevant ones)
5. run:

        go test -cover -v


**Requirement 5:** <br />
I used the Mocha framework for this requirement. To start the test simply navigate to requirement-5, run:

        npm install

in order to install the required modules, and than:

        npm test

to run the Mocha test
