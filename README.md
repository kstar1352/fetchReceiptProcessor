# Kailas Shekar
## Steps To Run Application


## Directions

- Run the following commands to create a docker image
 ```sh
docker build --tag kailasshekar/receipt-processor:latest .
docker run --publish 3000:3000 kailasshekar/receipt-processor:latest
```
- Now we can run the first two commands in the requests.http file
- For the first command we need to grab the response id and input that in the {id} of the second command
- Once we replace the id {id} in the 2nd command we can run the 2nd command to get the number of points

## Done!