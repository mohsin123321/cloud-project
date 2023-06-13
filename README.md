# Agrio 
[<img src="https://cordis.europa.eu/docs/news/images/2020-02/413531.jpg">](agrio)



The aim of the project is to create a scalable web application in Golang that can receive data from an external simulator generating IoT data. To ensure smooth and efficient delivery of the application to production, we need to implement Continuous Integration and Continuous Deployment (CI/CD) using AWS CodePipeline. Our goal is to automate the deployment process and reduce the time required for manual deployments. We will also need to ensure that the application can be scaled up easily to handle increasing traffic as more users begin to use it. Our focus is to leverage AWS services such as ECS,EC2, CodePipeline, and CodeBuild to automate the deployment process and achieve scalability.
Additionally, we will be using MongoDB as our database to store the received IoT data. Our aim is to create a seamless and reliable pipeline that can handle the deployment and scaling of the application while ensuring the data is stored securely and reliably in our database.
 
 ## Technologies 
- Programming language: Golang 1.19
- Database : MongoDB
- Container : Docker
- Simulator link 


## Commands
1- run following commands to run the simulator 

`docker-compose pull`

`docker-compose up`

2- after running simulator dockerfile, go to `http://localhost:8090/` 

3- create a target session and name it as you want and in address, enter the http://localhost:8080/device

4- run the Go application with `go run main.go` command

  # AWS-Deployment-Documentation
you can find a step-by-step walkthrough of implementing CI/CD pipeline for a Golang application on AWS using Elastic Beanstalk, CodePipeline, and CodeBuild in the following link :

[AWS-Deployment-Documentation](https://github.com/sarahrajabazdeh/AWS-Deployment-Documentation/blob/main/README.md)
