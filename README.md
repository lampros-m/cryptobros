
### **1. Create this variable in cli.go file in package config (check that the filepath is included in .gitignore file):**
> RestartNetworkManagerCLICommand = `echo <your super user password here>" | sudo -S service NetworkManager restart`

### **2. The configuration of program execution is up to config.go file under config package - Separate configuration instances are created at main.go files for each CMD**

### **3. Make sure that your network manager applies IPv6 automatically when connecting to a network**

### **4. To fetch data for current day:**
`make build run-create`

### **5. To query data for a specific day (the query is build in main.go of this command - documentation is included):**
`make build run-query`

The results are stored in /results in JSON format
