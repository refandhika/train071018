# Train #071018

## Specification
- Docker with:
	- Go
	- MySQL

## How To Run
- `docker-compose up`
- Wait `mysql` and `app` up
- Wait `migrate` to finish on migration
- Open `http://localhost:4545` in browser

## Notes
- `migrate` will keep retrying until `mysql` finish loading
- Wait everything loaded before opening the app on the browser
