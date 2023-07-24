# Task-Management-API-
Task Management API using the Go programming 

The assignment involves building a Task Management API using the Go programming language, Gin web framework, and GORM as the ORM library for database interactions. The API will allow users to perform CRUD operations on tasks, including creating, reading, updating, and deleting tasks. The tasks are represented by a "Task" struct with attributes like title, description, due date, and status. The API is tested using tools like curl or Postman to send HTTP requests and verify the functionality of each API endpoint. The assignment covers essential concepts of building RESTful APIs in Go and provides practical experience in handling HTTP requests and database operations. It's a valuable learning opportunity for backend development with Go and related web technologies.


<h1>Workflows:</h1>

### Technology Stack:

GOlang: It is a programming language used for building efficient and scalable applications.

Gin: Gin is a web framework written in Golang that is lightweight and fast. It is used to handle HTTP requests and build RESTful APIs.

Gorm: Gorm is an Object-Relational Mapping (ORM) library for Golang. It allows you to interact with the database using Go structs instead of writing SQL queries directly.

SQLite: SQLite is a lightweight and self-contained database engine that is easy to use and does not require a separate server.



### Step 1 : Set up the project structure:

Create a new Go project.

Install the required dependencies, such as Gin and GORM, using "go get."


### Step 2 : Requirements:

Implement CRUD operations (Create, Read, Update, Delete) for managing tasks.

Each task will have the following properties:

    Title: The title or name of the task.

    Description: A description of the task.

    DueDate: The due date for completing the task.

    Status: The status of the task (e.g., pending, in progress, completed, etc.).

The API should be able to create new tasks, retrieve specific tasks by ID, update tasks, delete tasks, and retrieve a list of all tasks.

Use Gorm to define the database model for the Task entity and handle database interactions.

Use Gin to define API routes and handlers for each CRUD operation.

Use SQLite as the database to store task information.


### Step 3 : Expected API Endpoints:

POST /tasks: Create a new task by sending a JSON payload containing the task details.

GET /tasks/:id: Retrieve a specific task by its ID.

PUT /tasks/:id: Update an existing task by providing the task ID and a JSON payload with updated task details.

DELETE /tasks/:id: Delete a task by its ID.

GET /tasks: Retrieve a list of all tasks.


### Step 4 : Sample Data for Testing

Using sample JSON data to test the API endpoints. For example:

{ 

"title": "Task 1",

 "description": "Description for Task 1", 

"due_date": "2023-07-25",

 "status": "pending" 

}


### Step 5 : Running and Testing the API

Save the code in the main.go file.

Run the API by executing the main.go file.

Use tools like curl or Postman to test the API endpoints by sending HTTP requests.



#Application architecture.
![Screenshot 2023-07-24 043515](https://github.com/Ayush5935/Task-Management-API-/assets/64814485/11648b99-9fb5-4deb-bd75-c88d08d91a5d)


#Working of Model 

![CREATE](https://github.com/Ayush5935/Task-Management-API-/assets/64814485/3a1f10e5-0eca-47e0-abb1-ae5a75c4b91a)
![GET_ALL](https://github.com/Ayush5935/Task-Management-API-/assets/64814485/8f26bc8c-b946-4ed6-bbdb-8f8ba8bcc441)
![UPDATE](https://github.com/Ayush5935/Task-Management-API-/assets/64814485/5a76171d-abd6-4d61-a030-038838176195)
![DELETE](https://github.com/Ayush5935/Task-Management-API-/assets/64814485/58aeb258-0934-4de7-b8aa-70d18c8c2e3e)
