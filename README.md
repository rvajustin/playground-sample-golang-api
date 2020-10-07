# API Playground: an API written in Go

---

## Getting started
Fork & clone the repository locally to work with the code.


## Running locally
Run the following code from the repository root to start the application locally
```bash
go run main.go
```
This will run the application on port 8080 and you can access the endpoints defined in main.v from a web browser or other HTTP utility


## Running in Docker
You'll need to have docker-compose and then the following code from the repository root to start the application locally 
```bash
docker-compose up
```
This will run the application on port 8080 and you can access the endpoints defined in main.v from a web browser or other HTTP utility


## GitHub Workflows
### Push to ACR
This action builds the Docker image whenever any push to the repository is made. Using a number of secrets and an action, [ams0/acr-task-github-action@v1](https://github.com/ams0/acr-task-github-action), this workflow pushes a container image to an Azure Container Registry.

**Secrets**
These secrets should be defined in your repository under Settings/Secrets
```yml
service_principal: a service principal that has contributor access to the container registry
service_principal_password: the password to the aforementioned service principal
tenant: your Azure tenant GUID
registry: your Azure container registry name (given a FQDN of <registry>.azurecr.io, only <registry> should be supplied)
repository: the repository of your image (this will end up being the prefix for the image stored in Azure container registry)
image: the image name (this will end up being the suffix for the image stored in Azure container registry)
```
