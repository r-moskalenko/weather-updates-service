# Weather Updates Service

It allows you to subscribe to regular weather forecast updates in the selected location.

### Endpoints

- Allows receive weather updates for specific city 
    ```
    /weather?city=<city>
    ```

- Allows you to subscribe for weather updates
    ```
    /subscribe?email=<email>,city=<city>,frequenc=<email>
    ```

- Allows you to confirm email
    ```
    /confirm/:token
    ```

- Unsubscribe from weather updates

    ```
    /unsubscribe/:token
    ```

### Required env variables.

It is better for storing them into .env file.

```
SCHEME - {http,https}
HOST -
SENDGRID_API_KEY - 
FROM_EMAIL - 
MAIL_VERIF_TEMPLATE_ID - 
```

### How to run this project

1. Install docker and docker compowe on your machine
2. Run this command to build image of container with web app:
    ```
    sudo docker build -t go-app .
    ```
3. Start two containers with Rest Api & DB using docker compose.
    ```
    docker compose up
    ```