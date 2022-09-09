## 2022/09/09

​		go mod 在使用本地自定义包的时候，会自动去GOPATH下寻找，如果没有的话就会报错，但我们自己定义的包在相对路径里就能寻到，go mod不支持相对路径的import，因此要么将这个路径加入到GOPATH中，要么publish，但因为我们只是开发一个没必要publish出去。

​		这里使用了replace，见main包的go.mod文件。