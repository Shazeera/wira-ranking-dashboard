The Wira Ranking Dashboard is a web application designed to display player rankings for the game WIRA. The system consists of a backend built with Golang and a frontend developed using Vue.js, both hosted with Nginx.

This project includes the following features:

Backend: RESTful API for managing player data, utilizing PostgreSQL for storage.
Frontend: Interactive and responsive user interface to view rankings.
Hosting: Configured using Nginx to serve both the frontend and backend.

How to Run the Project
Prerequisites
-Install PostgreSQL, Golang, Node.js, and npm on your system.
-Install Nginx if you want to host the project.

Steps to Run
Backend

-Navigate to the backend folder:
  cd backend
  
-Install dependencies:
  go mod tidy

-Update the main.go file with your PostgreSQL credentials:
  dsn := "host=localhost user=postgres password=your_password dbname=your_db port=5432 sslmode=disable"

-Run database migrations and seed data:
  go run main.go

-Start the backend server:
  go run main.go

The backend will run at http://localhost:8080

Frontend

-Navigate to the frontend folder:
  cd ../frontend
  
-Install dependencies:
  npm install

-Build the frontend:
  npm run build

-Serve the frontend using Nginx or npm run serve (for development):
  npm run serve
  
The frontend will run at http://localhost:8081

Hosting with Nginx

-Install Nginx:
  sudo apt install nginx

-Configure Nginx to serve the frontend and backend:
  server {
    listen 80;
    server_name your-ip-address;

    location / {
        root /path/to/frontend/dist;
        index index.html;
    }

    location /api/ {
        proxy_pass http://localhost:8080/;
    }
}

Replace your-ip-address with your local IP or public IP and /path/to/ with the path to your project.

-Restart Nginx:

  
Accessing the Project
-Frontend: http://your-ip-address
-Backend API: http://your-ip-address/api/players



