
### **1. Create this variable in cli.go file in package config (check that the filepath is included in .gitignore file):**
> RestartNetworkManagerCLICommand = `echo <your super user password here>" | sudo -S service NetworkManager restart`

### **2. The configuration of program execution is up to config.go file under config package - Separate configuration instances are created at main.go files for each CMD**