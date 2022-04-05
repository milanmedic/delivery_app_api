##### Git Hooks
Instructions for Git Hooks located in .githooks folder in the ```package.json``` file.

##### Dependencies
    * Go Gin: https://github.com/gin-gonic/gin

##### Helpers

Config for caching Git credentials: ```git config --global credential.helper "cache --timeout=86400"```

##### TODO
* Registration (DONE)
* Login (DONE)
    * Forbid Login if user not verified - Do after verification logic
* OAuth Registration - DONE
    * Since Github OAuth provides only limited information:
        * Name
        * Surname
        * Username
        * Email
        * City
        Disallow the ordering of items unless the User has provided all the information missing: (NOT DONE)
            * Password
            * Date of Birth
            * Address:
                * Street
                * Street Number
                * Street Num Postfix

* OAuth Login - DONE
    * Forbid Login if user not verified - Do after verification logic (NOT DONE)
* Verify Customer (Admin) (NOT DONE)
* Verify Deliverer (Admin) (NOT DONE)
* Send Verification Email Confirmation (Admin) (NOT DONE)
* Get Profile (Customer, Admin, Deliverer) (NOT DONE)
* Get Current Order (Customer) - Verified Only (NOT DONE)
* Place New Order (Customer) - Verified Only (NOT DONE)
* Previous Orders (Customer) - Verified Only (NOT DONE)
* All Orders (Admin) (NOT DONE)
* Add Article (Admin) (NOT DONE)

##### EDGE CASES
* Since GitHub Customer registration doesn't provide all the needed field for the address, an address will be created with missing fields.
  Need to figure out a way to eliminate this edge-case.