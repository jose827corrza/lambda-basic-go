# LAMBDA BASIC GO

Simple Code To implement in CI/CD flow into AWS Lambda, using as main langugage Golang.

## Important Info

The field called handler, must be set the same way of directory as we set the -o in the build command.

Example.

the command shown below indicates that the binary will be placed into the build folder. Hence then handler  option in Lambda must say something like.

> GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o build/bootstrap cmd/main.go

> build/bootstrap

## Steps to setup the code to implement in Lambda manually

1. Run the command that will create the binary.

> GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o build/bootstrap cmd/main.go

2. Generate the zip containing the artifact generated in the previous step.

> zip myFunction.zip build/bootstrap

3. Upload the .zip file in the already created lambda.

## Steps to setup the code automatically, using AWS Pipeline

Send a commit into the main branch of the repository.
