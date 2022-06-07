# Introduction
During the vtubestudio integration project, there is a need to test out if the frontend is integrating properly with the backend through websocket.

However, it feels like a hassel to ensure that all developers have access to VTubeStudio before they can start testing.

Hence, this project is born. A mock of the core components we want to test.

This will be served at the default addr of vtuberstudio, ws://localhost:8001

# Features
We will mock a few features here.
- plugin registration
- plugin authentication
- hotkey trigger
- expected error responses

# Dependencies
|Dependency|Reason|
|-|-|
|github.com/gorilla/websocket|Websocket library to simulate ws connection to vtubestudio|
|github.com/gin-gonic/gin|Webserver for serving the websocket|

# Development trace
This section will hold information of how our development has evolved as each commit goes in.

## 