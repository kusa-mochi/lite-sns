# lite-sns

## How to debug

### Make config files

You must make config files and set params. Their format is as in debug dir's.

- src/cmd/app_server/conf/release/app_server.json
- src/cmd/client/lite-sns-client/src/conf/release/lite-sns-client/json

### Run a devenvs

1. Run the dev container using command `docker compose up -d` in the dir `./.lite-sns-docker`.
1. Attach a "lite-sns-devenvs" container using VSCode Docker extensions.
1. Open `/project` dir in attaching VSCode window.

### Initialize DB tables

1. Open a terminal in the VSCode window and run command:
    ```
    cd /project/src/cmd/db_initializer
    go run main.go
    ```
    You can close this terminal.

1. You can login to lite-sns using these test user data after starting the backend server and the frontend:
    
    |No.|email address|password|
    |---:|:---|:---|
    |1|lite-sns_dev-admin@slash-mochi.net|123412341234|
    |2|lite-sns_dev-user1@slash-mochi.net|abcdefghijklmnop|
    |3|lite-sns_dev-user2@slash-mochi.net|this_is_a_password|
    |4|lite-sns_dev-user3@slash-mochi.net|this-is-a-password-too|

### Run a backend server

1. Open a terminal in the VSCode window and run command:
    ```
    cd /project/src/cmd/app_server
    go run main.go
    ```
    Now the backend server is built and starts running.

### Run a frontend

1. Open a terminal in the VSCode window and run command:
    ```
    cd /project/src/cmd/client/lite-sns-client
    npm run dev
    ```
1. In the docker host, open frontend URL using Web browser. The URL is determined by params in the config file `/project/src/cmd/clinet/lite-sns-client/src/conf/release/lite-sns-client.json`.
    - You must make config files before debugging. See "Make config files".

## How to read PlantUML diagrams

### Run a PlantUML server

1. Install the PlantUML extension.
1. Set the PlantUML server address to `http://localhost:30080/`
1. Run the PlantUML server using command `docker compose -f docker-compose.doc.yml up -d` in the dir `./.lite-sns-docker`.
1. Open a PlantUML document (./doc/*.puml).
1. Put a cursor to PlantUML code and press Alt+F4.
