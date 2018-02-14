# Weather Service

Design

1. Wait for address to get submitted from the page.
2. Call Google geocoding service, basically sending an address and receiving a package back containing lat/lng values.
3. Call DarkSide online API and send a request with lat/lng values.
4. Upone receiving the reply, display output.

Build status

[![CircleCI](https://circleci.com/gh/MagnusTiberius/weatherservice.svg?style=svg)](https://circleci.com/gh/MagnusTiberius/weatherservice)


Please refer to CircleCI https://circleci.com/gh/MagnusTiberius/weatherservice


References:
1. https://developers.google.com/maps/documentation/geocoding/intro
2. https://darksky.net
