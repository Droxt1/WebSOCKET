# User Actions

1. **POST Request: Signup**

    - Endpoint: http://127.0.0.1:8080/api/v1/user/signup

   Description: This POST request is used to register a new user in the system.

   ```shell
   curl -X POST -H "Content-Type: application/json" -d '{
       "username": "sample_user",
       "email": "sample_user@gmail.com",
       "password": "sample_password123"
   }' http://127.0.0.1:8080/api/v1/user/signup
   ```

2. **POST Request: Login**

    - Endpoint: http://127.0.0.1:8080/api/v1/user/login

   Description: This POST request is used to authenticate a user and log them into the system.

   ```shell
   curl -X POST -H "Content-Type: application/json" -d '{
       "email": "sample_email@gmail.com",
       "password": "sample_password123"
   }' http://127.0.0.1:8080/api/v1/user/login
   ```

3. **GET Request: Logout**

    - Endpoint: http://127.0.0.1:8080/api/v1/user/logout

   Description: This GET request is used to log a user out from the system.

   ```shell
   curl -X GET http://127.0.0.1:8080/api/v1/user/logout
   ```

# WebSocket Actions

#### Create Room

1. **POST Request: Create Room**

    - Endpoint: http://127.0.0.1:8080/api/v1/ws/createRoom

   Description: This POST request is used to create a new room.

   ```shell
   curl -X POST -H "Content-Type: application/json" -d '{
       "id": "Room 1",
       "name": "1"
   }' http://127.0.0.1:8080/api/v1/ws/createRoom
   ```

   Request Body (raw, text):
   ```json
   {
       "id": "Room 1",
       "name": "1"
   }
   ```

#### Join Room

2. **WebSocket Request: Join Room**

    - Endpoint: ws://localhost:8080/api/v1/ws/joinRoom/67?userId=1&username=sample_user

   Description: This WebSocket request is used to join a room with the specified user ID and username.

   WebSocket Connection Example:
   ```shell
   # Example using WebSocket connection
   websocat ws://localhost:8080/api/v1/ws/joinRoom/67?userId=1&username=sample_user
   ```

#### Get Clients by Room

3. **GET Request: Get Clients by Room**

    - Endpoint: http://127.0.0.1:8080/api/v1/ws/get/:1

   Description: This GET request is used to retrieve clients by room.

   ```shell
   curl -X GET http://127.0.0.1:8080/api/v1/ws/get/1
   ```

   Path Variables:
    - 1: The room identifier (replace with the desired room ID).

#### Get All Rooms

4. **Get Rooms**

    GET Request: Get Rooms
        Endpoint: http://127.0.0.1:8080/api/v1/ws/getRooms

    Description: This GET request is used to retrieve a list of available rooms.


    shell
    
    ```shell
    curl -X GET http://127.0.0.1:8080/api/v1/ws/getRooms
    ```

