# go-test
Go Fiber Static Web Server

This is a simple web application built with the Go programming language and the Fiber web framework. It serves a static index.html file to demonstrate the basic functionality of a Fiber web server.
Prerequisites

Before you begin, ensure you have the following installed on your system:

    Go: You need Go version 1.16 or higher. You can download it from the official Go website.

    Fiber: The project uses the Fiber framework, which will be automatically installed when you run the project for the first time.

Getting Started

Follow these steps to get the project up and running.
1. Clone the Repository

First, clone the project from your GitHub repository to your local machine.

git clone https://github.com/your-username/your-repo-name.git
cd your-repo-name

2. Install Dependencies

The go.mod file handles the project's dependencies. Use the following command to download and install the required modules, including Fiber.

go mod tidy

3. Run the Application

Execute the main.go file to start the web server.

go run main.go

You should see a message in your terminal indicating that the server has started.

Fiber app is listening on :3000

4. View the Website

Open your web browser and navigate to http://localhost:3000. You should see the content of your index.html file.
Project Structure

    main.go: This is the main Go file that initializes and starts the Fiber web server. It is configured to serve the index.html file.

    index.html: This is the static HTML file that will be served by the web server.
    
This app runs on localhost:3000
