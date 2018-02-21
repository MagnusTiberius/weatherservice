# Weather Service

**Demo**
1. Weather data is at http://35.225.142.51:8090/address/Dallas%20TX/255589200
                      http://35.225.142.51:8090/address/Seattle%20WA
2. User interface is at http://35.193.30.249:8088/

note: you will need to download a Chrome plugin from CORS. Once the plugin is installed, please enable CORS so that the UI can receive data from the sibling microservice. Please download it at  https://chrome.google.com/webstore/detail/allow-control-allow-origi/nlfbmbojpeacfghkpbjhddihlkkiljbi?hl=en

**Design**

1. Wait for address and/or date-time to get submitted into the page.
2. Pass address and/or datetime to microservice. See https://github.com/MagnusTiberius/weatherservicedarksky for details.
3. Upon receiving the reply, display output to a textarea field in json format.

**Build status**

[![CircleCI](https://circleci.com/gh/MagnusTiberius/weatherservice.svg?style=svg)](https://circleci.com/gh/MagnusTiberius/weatherservice)

Please see .circleci/config.yml for details about the CI build.

**Docker Container**

Each build will create a docker container which is packaged and sent to a Google Cloud project repo.


**Google Cloud**
1. A config.yml script sets up a gcloud environment during build which then creates and pushes the docker container to the repo.
2. A create pod command is entered in a cluster console
   ```
   kubectl run web-weather --image=us.gcr.io/weatherservice-195512/weatherserviceweb --port 8088
   ```
3. A create service command is then used to expose it.
   ```
   kubectl expose deployment web-weather --type=LoadBalancer --port 8088 --target-port 8088
   ```

**Dependent Microservices**

1. https://github.com/MagnusTiberius/weatherservicedarksky

**References:**
1. https://developers.google.com/maps/documentation/geocoding/intro
2. https://darksky.net
