package main

import
(
   "github.com/cloudlibz/gocloud/gocloud"
   "fmt"
)
func main(){

  amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)
  fmt.Println(amazoncloud)
}
