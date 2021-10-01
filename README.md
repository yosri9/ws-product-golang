
## Work Sample for Product Role, Golang Variant



configuration:

1- go to env/env.go file    and configure database constants and
    routes constant

2- by default global limit rating is 2 request per hour , go to 
    middleware/global_middleware/stat_middleware.go
    and change variables ratingTimeLimiter ,requestNumberLimiter  to values you want 

3- By default limit rating per person is 15 requests per minute, you can access 
    middleware/per_client_middleware/store_config.go file 
    and  change statTokens ,statInterval       

4- you can type: "go mod init server" in the terminal with path src/server 
    to import packages and then type  go mod tidy





