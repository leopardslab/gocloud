![GoCloud Logo](assets/logo.png)

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/0fce581810a6420aaca4ba6757c54529)](https://www.codacy.com/app/shlokgilda/gocloud?utm_source=github.com&utm_medium=referral&utm_content=shlokgilda/gocloud&utm_campaign=Badge_Grade)
[![Build Status](https://travis-ci.org/shlokgilda/gocloud.svg?branch=master)](https://travis-ci.org/shlokgilda/gocloud)
[![Gitter](https://img.shields.io/badge/chat-on%20gitter-ff006f.svg?style=flat-square)](https://gitter.im/shlokgilda/gocloud)
[![docs](https://camo.githubusercontent.com/df8e028288079a740c10e6cfaad2fa0e0c96014d/687474703a2f2f696d672e736869656c64732e696f2f62616467652f446f63732d6c61746573742d677265656e2e737667)](docs)

# gocloud: Cloud services library.

GoCloud is a golang library which hides differences between different cloud providers' (AWS,GCP,Openstack etc) APIs and allows you to manage different cloud resources through a unified and easy to use API.

![GoCloud Architecture](assets/gocloudarchitecture.png)

## Service Types

**Compute** -- Allows you to manage cloud and virtual servers.

**Compute Storage** -- Allows you to manage Compute storage.

**Container** -- Allows users to install and deploy containers onto container based virtualization platforms.

**Load balancer** -- Allows you to manager Load Balancer service.

**DNS** -- Allows you to manage DNS service.

## Service Providers

### AWS

* EC2 Compute [Link to example](examples/compute/ec2/ec2.md)
* EC2 Storage [Link to example](examples/storage/aws_storage/aws_storage.md)
* Amazon Elastic Container Service (Container) [Link to example](examples/container/aws_container/aws_container.md)
* Elastic Load Balancing [Link to example](examples/loadbalancer/aws_loadbalancer/aws_loadbalancer.md)
* AWS Route53 (DNS) [Link to example](examples/dns/aws_route53/aws_route53.md)

### Google

* Google Compute [Link to example](examples/compute/gce/gce.md)
* Google Compute Storage [Link to example](examples/storage/google_storage/google_storage.md)
* Google Container Service (Container) [Link to example](examples/container/google_container/google_container.md)
* Google Elastic Load Balancing [Link to example](examples/loadbalancer/google_loadbalancer/google_loadbalancer.md)
* Google DNS [Link to example](examples/dns/google_dns/google_dns.md)

Currently, implementations for other cloud providers are being worked on.

## Installation instructions for Linux (Ubuntu)
1. Install golang.  
   ```
   $ sudo apt-get update -y
   $ sudo apt-get install golang -y
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
$ go get github.com/shlokgilda/gocloud
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

6. You are all set to use gocloud! Check out the following YouTube videos for more information and usage examples:
https://youtu.be/4LxsAeoonlY?list=PLOdfztY25UNnxK_0KRRHSngJIyVLDKZxq&t=3

## Development setup

```
$ git clone https://github.com/shlokgilda/gocloud
$ cd gocloud
```

## Unit tests

```
$ cd gocloud
$ go test -v
```
