/*
Package plausible implements a client/wrapper for the Plausible Analytics platform.

To start interacting with the API, a client must be created. After that, it's possible to work with multiple sites by requesting site handlers to the client:

  // create a new client with a given token
  client := plausible.NewClient("<your_api_token>")

  // get a site handler from the client
  siteHandler := client.Site("example.com")

  // use client and siteHandler to make queries to the API
*/
package plausible
