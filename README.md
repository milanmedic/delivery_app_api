##### Git Hooks
Instructions for Git Hooks located in .githooks folder in the ```package.json``` file.

##### Dependencies
    * Go Gin: https://github.com/gin-gonic/gin

##### Helpers

Config for caching Git credentials: ```git config --global credential.helper "cache --timeout=86400"```

##### EDGE CASES
* Since GitHub Customer registration doesn't provide all the needed field for the address, an address will be created with missing fields.
  Need to figure out a way to eliminate this edge-case.