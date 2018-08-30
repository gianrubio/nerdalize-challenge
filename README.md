## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ cd api 
	$ buffalo dev

## Creating a new task
```
$ go run cli/main.go task create
Task successfully created
{"id":"00d0054e-416b-4cab-ad43-7bb07f34d561","DockerImage":"gianrubio/task:1","Environment":["BLA=1","X=3"],"Command":"/bin/my_binary","Schedule":"@every 1d","Limits":{"cpu":"200m","memory":"100m"},"Volume":"s3://my-bucket/my-dir/","User":{"id":"5c229a0e-acd3-432d-ab43-666b77de0166","email":"gianrubio@gmail.com","token":"0621f877-24c4-4a42-99e4-c032c69d3015"}}
```

## Listing your tasks
```
$ go run cli/main.go task list
Task successfully created
{"id":"00d0054e-416b-4cab-ad43-7bb07f34d561","DockerImage":"gianrubio/task:1","Environment":["BLA=1","X=3"],"Command":"/bin/my_binary","Schedule":"@every 1d","Limits":{"cpu":"200m","memory":"100m"},"Volume":"s3://my-bucket/my-dir/","User":{"id":"5c229a0e-acd3-432d-ab43-666b77de0166","email":"gianrubio@gmail.com","token":"0621f877-24c4-4a42-99e4-c032c69d3015"}}
```