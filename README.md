# Agrio 
[<img src="https://cordis.europa.eu/docs/news/images/2020-02/413531.jpg">](agrio)



he aim of our project is to obtain information on the status of the crops in real-time by using IoT devices to collect data from sensors.You can gain knowledge to improve your harvesting decisions by using predictive analytics.The trend analysis aids farmers in forecasting impending weather and crop harvesting.
 in this case we will have rational farm management plans to save both time and money.
 
 ## Technologies 
- Programming language: Golang 1.19
- Database : MongoDB
- Container : Docker
- Simulator link 

## Architecture 
the project has made of three main components.
1- Simulator : it is a prepared directory in which we are using it for generating data.
for now we are using simulator to generate the data, later on we will connect to real sensors. 
2- Golang applicaition : here we have the core logic of our applicaiton, after validating the token in middleware, we will create the devices.
3- MongoDb: after creating devices, we will save the data into mongodb because the data that is coming from the simulator are based on different sensors and different data types.
![Untitled Diagram drawio](https://user-images.githubusercontent.com/33392969/195940133-8c75f367-cbd0-4f31-ba50-ded56d4567f5.png)

## Commands
