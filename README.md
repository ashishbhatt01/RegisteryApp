# RegisteryApp
Registery App for adding, subtracting, reading values

It starts the application with port 4000. Port can be changed in main.go file.

Sample request for addition:

Request Type: POST, 
Request Url: http://localhost:4000/registry/add,
Body:{
    "value": 50
}

Sample request for subtraction:

Request Type: POST,
Request Url: http://localhost:4000/registry/subs,
Body:{
    "value": 20
}

Sample request to get current value

Request Type: GET,
Request Url: http://localhost:4000/registry/value
