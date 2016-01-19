# trudiedo-api

API server for Trudie Do

initially based on http://thenewstack.io/make-a-restful-json-api-go/

Run this as a docker container

## Running container

```
docker run -it --rm -p 8080:8080  vallard/trudiedo-api /bin/bash
```
To run in demon mode: 

```
docker run -p 8080:8080 -d --name trudieAPI vallard/trudiedo-api
```
