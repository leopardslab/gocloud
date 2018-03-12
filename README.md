![GoCloud Logo](assets/logo.png)

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/0fce581810a6420aaca4ba6757c54529)](https://www.codacy.com/app/cloudlibz/gocloud?utm_source=github.com&utm_medium=referral&utm_content=cloudlibz/gocloud&utm_campaign=Badge_Grade)
[![Build Status](https://travis-ci.org/cloudlibz/gocloud.svg?branch=master)](https://travis-ci.org/cloudlibz/gocloud)
[![Gitter](https://img.shields.io/badge/chat-on%20gitter-ff006f.svg?style=flat-square)](https://gitter.im/cloudlibz/gocloud)

# gocloud: Cloud services library.

GoCloud is a golang library which hides differences between different cloud providers' (AWS,GCP,Openstack etc) APIs and allows you to manage different cloud resources through a unified and easy to use API.

## Service Types

**Compute** -- Allows you to manage cloud and virtual servers.

**Compute Storage** -- Allows you to manage Compute storage.

**Container** -- Allows users to install and deploy containers onto container based virtualization platforms.

**Load balancer** -- Allows you to manager Load Balancer service.

**DNS** -- Allows you to manage DNS service.

## Service Providers

### AWS

* EC2 (Compute)
* EC2 Storage
* Amazon Elastic Container Service (Container)
* Elastic Load Balancing
* AWS Route53 (DNS)

### Google

* Google Compute
* Google Compute Storage
* Google Container Service (Container)
* Google Elastic Load Balancing
* Google DNS

Currently, implementations for other cloud providers are being worked on.

## Installation instructions for Linux (Ubuntu)
1. Install golang.  
   ```
   $ sudo apt-get update -y
   $ sudo apt-get install golang
   ```
   
2. Set GOPATH environment variable. Run `gedit ~/.bashrc`.  
  Copy the following in your .bashrc file:
  ```
  export GOPATH=$HOME/gopath
  export GOBIN=$HOME/gopath/bin
  ```
     
3. Test your installation by copying the following piece of code in a file. Save the file as *gotest.go*. Run the file using the command `go run gotest.go`. If that command returns “Hello World!”, then Go is successfully installed and functional.
```golang
package main
import "fmt"
func main() {
    fmt.Printf("Hello World!\n")
}
```

4. Now we need to fetch the gocloud repository and other necessary packages. Run the following commands in order:
```
$ go get github.com/cloudlibz/gocloud
$ go get golang.org/x/oauth2
$ go get cloud.google.com/go/compute/metadata
```

5. Download your AWS and Google access credentials and store them in a file in your <b>HOME</b> directory.  
   Save your AWS credentials in a file named *amazoncloudconfig.json* and your Google Cloud credentials in a file named *googlecloudconfig.json*.  
   You can also set your credentials as environment variables.  
   For AWS:  
   ```
   export AWSAccessKeyID =  "xxxxxxxxxxxx"
   export AWSSecretKey = "xxxxxxxxxxxx"
   ```
   For Google Cloud Services:
   ```
   export PrivateKey =  "xxxxxxxxxxxx"
   export Type =  "xxxxxxxxxxxx"
   export ProjectID = "xxxxxxxxxxxx"
   export PrivateKeyID = "xxxxxxxxxxxx"
   export ClientEmail = "xxxxxxxxxxxx"
   export ClientID = "xxxxxxxxxxxx"
   export AuthURI = "xxxxxxxxxxxx"
   export TokenURI = "xxxxxxxxxxxx"
   export AuthProviderX509CertURL = "xxxxxxxxxxxx"
   export ClientX509CertURL =  "xxxxxxxxxxxx"
   ```

6. You are all set to you gocloud!

## Development setup

```
$ git clone https://github.com/cloudlibz/gocloud
$ cd gocloud
```

## Unit tests

```
$ cd gocloud
$ go test -v
```
