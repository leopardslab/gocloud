![GoCloud Logo](assets/logo.png)

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/0fce581810a6420aaca4ba6757c54529)](https://www.codacy.com/app/dilantha111/gocloud?utm_source=github.com&utm_medium=referral&utm_content=dilantha111/gocloud&utm_campaign=Badge_Grade)
[![Build Status](https://travis-ci.org/dilantha111/gocloud.svg?branch=master)](https://travis-ci.org/dilantha111/gocloud)
[![Gitter](https://img.shields.io/badge/chat-on%20gitter-ff006f.svg?style=flat-square)](https://gitter.im/cloudlibz/gocloud)

# gocloud

Cloud services library

GoCloud is a golang library which hides differences between different cloud providers' (AWS,GCP,Openstack etc) APIs and allows you to manage different cloud resources through a unified and easy to use API.

# Service Types

**Compute** --Allows you to manage cloud and virtual servers.

**Compute Storage** --Allows you to manage Compute storage.

**Container** --Allows users to install and deploy containers onto container based virtualization platforms.

**Load balancer** --Allows you to manager Load Balancer service.

**DNS** --Allows you to manage DNS service.

## Service Providers

### AWS

* EC2 (Compute)
* EC2 Storage
* Amazon Elastic Container Service (Container)
* Elastic Load Balancing
* AWS Route53

### Google

* Google Compute
* Google Compute Storage
* Google Container Service (Container)
* Google Elastic Load Balancing
* Google DNS

Currently, implementaions for other cloud providers are being worked on.
