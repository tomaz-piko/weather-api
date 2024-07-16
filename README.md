# weather-api
This is a simple weather API, providing info for three croatian cities (Reka, Zadar, Split). The project supports in-memory caching to speedup requests. It also includes a Dockerfile to be used in a docker container.

No matter the installation the api should be accessible at localhost:8080.
The api supports two routes:
- localhost:8080/weather -> Returns weather info for all supported cities
- localhost:8080/weather/{city} -> Returns weather info for a certain city

Current supported cities are (insert in second route above):
1) Reka
2) Zadar
3) Split

<i>The api is not case sensitive. </i>

<b><big>Important notes:</big></b>
If during testing and of these values need to be changes for any reason here is how:
- Caching
    - Caching support is hardcoded. To disable: 
    ```Go
    // In main.go switch the following line to false 
    var enable_cache = true   
    ```
    - Cache items expire in 1 minute and expired items are cleanedup every 5 minutes. To change:
    ```Go
    // In main.go change following lines
    var cache_purge_interval = 5 * time.Minute // How often to purge expired cache items
    var cache_item_expiry = 1 * time.Minute    // How long to before cache items expires
    ```
- Release mode
    - ReleaseMode is deactivated for debugging purposses. To activate:
    ```Go
    // In main.go
    //gin.SetMode(gin.ReleaseMode) // Uncomment this to enable ReleaseMode
    ```

## Prerequisites
- <b>Go</b> programming language version >=1.22 to build and run the app or
- <b>Docker</b> to run the app in a Docker container

## Running the app

Clone the project to desired location and run the following command from the same location.
```bash
    go build -o weather-api && ./weather-api 
```

## Running with Docker

Clone the project to desired location and run the following commands from the same location.
```bash
# Build an image (named weather-api) based on the Dockerfile in current dir
sudo docker build -t weather-api .
# Starts a container (named weather-api) with an exposed 8080 port
sudo docker run -p 8080:8080 weather-api
# Alternative (detach mode) if you want the container to remain running after terminal is closed:
sudo docker run -d --name weather-api -p 8080:8080 weather-api
```

## Docker cleanup
To cleanup the docker container and image after testing the app.
```bash
# If container running in detach mode:
sudo docker stop weather-api # Stops the container
sudo docker rm weather-api # Removes the container
sudo docker rmi weather-api # Removes the image
```
