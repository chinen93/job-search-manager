# Backend - Job Search Manager

You can use the backend in two ways:

- **Bootstrap with an existing CSV**: Useful if you've already tracked applications in a spreadsheet.
- **Start from scratch**: Add positions manually via the API.

## Load Data from CSV

If you’ve been using a spreadsheet to track applications, start here. This prevents manual entry of past data and gets the system up and running quickly.

### Configuration

Follow the expected format in the `data/positions.csv` file. Add each position in a new row let the helper function populate the server with your applications.

### Populate Server with CSV file

Make sure your CSV is formatted correctly.

```sh
Make run_helper
```

This populates the backend with data from your CSV.

## Server

To start the API server and interact with it via REST:

Use tools like [Postman](https://www.postman.com/) or `curl` to interact with the API.

```sh
Make run
```

## Run Tests

To execute the unit tests:

```sh
Make test
```
