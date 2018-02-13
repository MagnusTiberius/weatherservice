# Weather Service

Design

1. wait for address to get submitted from the page.
2. Call Google geocoding service, basically sending an address and receiving a package back containing lat/lng values.
3. Call DarkSide online API and send a request with lat/lng values.
4. Upone receiving the reply, display output.


References:
1. https://developers.google.com/maps/documentation/geocoding/intro
2. https://darksky.net
