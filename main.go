package main

import (
    "Login_Reg_Go/router"
)




func main() {
    r := router.SetupRouter()
    r.Run(":8080")
}


