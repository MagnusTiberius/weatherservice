# Weather Service

**Demo**
1. Weather data is at http://35.225.142.51:8090/address/Dallas%20TX/255589200
2. User interface is at http://35.193.30.249:8088/

note: you will need to downloand a Chrome plugin from CORS.
see https://chrome.google.com/webstore/detail/allow-control-allow-origi/nlfbmbojpeacfghkpbjhddihlkkiljbi?hl=en

**Design**

1. Wait for address and/or date-time to get submitted into the page.
2. Call Google geocoding service, basically sending an address and receiving a package back containing lat/lng values.
3. Call DarkSide online API and send a request with lat/lng and(or) date-time values.
4. Upon receiving the reply, display output.

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
