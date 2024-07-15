# weather-api
ProPlus challange.


sudo docker build -t weather-api .
sudo docker run -p 8080:8080 weather-api

sudo docker run -d --name weather-api -p 8080:8080 weather-api
sudo docker stop weather-api
sudo docker rm weather-api
sudo docker rmi weather-api