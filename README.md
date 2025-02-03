A secure and flexible password generator service built with Go, featuring a clean web interface and Docker containerization.

📋 Table of Contents

Features
Technical Stack
Installation
Usage
API Documentation
Deployment
Security
Contributing
✨ Features

🔐 Generate secure passwords with customizable length (8-12 characters)
🔡 Support for special characters
🌐 Clean and responsive web interface
🚀 RESTful API support
📦 Docker ready
🔄 Easy scaling with Docker Compose
🛠 Technical Stack

Backend: Go 1.21+
Frontend: HTML5, CSS3, JavaScript
Container: Docker
Web Server: Native Go HTTP Server
💻 Installation

Local Development

bash
# Clone the repository
git clone https://github.com/Rockphill/password-generator.git

# Navigate to project directory
cd password-generator

# Build and run
go run main.go
Docker

bash
# Build Docker image
docker build -t password-generator:latest .

# Run container
docker run -p 8080:8080 password-generator:latest
🚀 Usage

Web Interface

Access the web interface at http://localhost:8080

Select password length (8-12 characters)
Choose whether to include special characters
Click "Generate Password"
API Endpoint

HTTP
GET /generate?length=10&special=true
Parameters

length	integer	Password length (8-12)	8
special	boolean	Include special characters	false
Example Response

JSON
{
    "password": "Kj#9mP$2nL"
}
📁 Project Structure

Code
password-generator/
├── main.go
├── handlers/
│   └── handlers.go
├── utils/
│   └── password.go
├── templates/
│   └── index.html
└── Dockerfile
🚢 Deployment

Docker Compose

YAML
version: '3.8'
services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
      - ENVIRONMENT=production
    restart: always
🔒 Security Features

Input validation
XSS protection headers
Rate limiting
Secure random number generation
HTTPS support (in production)
🔍 Monitoring

Health checks
Resource monitoring
Request logging
💡 Configuration Options

Environment Variables

PORT	Server port	8080
ENVIRONMENT	Runtime environment	development
Resource Configuration

YAML
services:
  app:
    deploy:
      resources:
        limits:
          cpus: '0.5'
          memory: 256M
        reservations:
          cpus: '0.1'
          memory: 128M
🤝 Contributing

Fork the repository
Create your feature branch (git checkout -b feature/AmazingFeature)
Commit your changes (git commit -m 'Add some AmazingFeature')
Push to the branch (git push origin feature/AmazingFeature)
Open a Pull Request
📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

Last updated: 2025-02-03 16:20:01 UTC
Created by: Rockphill

Note: For production deployment, make sure to:

Set up proper SSL/TLS certificates
Configure appropriate security measures
Set up monitoring and logging
Review and adjust resource limits
For more information, please refer to the documentation.
