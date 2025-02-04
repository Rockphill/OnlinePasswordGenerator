# Password Generator Service
![Version](https://img.shields.io/badge/version-1.0.0-blue)
![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?logo=go)
![Last Updated](https://img.shields.io/badge/last%20updated-2025--02--04-brightgreen)
![Author](https://img.shields.io/badge/author-Rockphill-orange)

A secure and flexible password generator service built with Go, featuring a clean web interface and Docker containerization.

## 📋 Table of Contents
- [Features](#features)
- [Technical Stack](#technical-stack)
- [Installation](#installation)
- [Usage](#usage)
- [API Documentation](#api-documentation)
- [Deployment](#deployment)
- [Security](#security)
- [Contributing](#contributing)

## ✨ Features
- 🔐 Generate secure passwords with customizable length (8-12 characters)
- 🔡 Support for special characters
- 🌐 Clean and responsive web interface
- 🚀 RESTful API support
- 📦 Docker ready
- 🔄 Easy scaling with Docker Compose

## 🛠 Technical Stack
- **Backend**: Go 1.21+
- **Frontend**: HTML5, CSS3, JavaScript
- **Container**: Docker
- **Web Server**: Native Go HTTP Server

## 💻 Installation

### Local Development
\```bash
# Clone the repository
git clone https://github.com/Rockphill/password-generator.git

# Navigate to project directory
cd password-generator

# Build and run
go run main.go
\```

### Docker
\```bash
# Build Docker image
docker build -t password-generator:latest .

# Run container
docker run -p 8080:8080 password-generator:latest
\```

## 🚀 Usage

### Web Interface
Access the web interface at `http://localhost:8080`

1. Select password length (8-12 characters)
2. Choose whether to include special characters
3. Click "Generate Password"

### API Endpoint
\```http
GET /generate?length=10&special=true
\```

#### Parameters
| Parameter | Type    | Description                     | Default |
|-----------|---------|--------------------------------|---------|
| length    | integer | Password length (8-12)         | 8       |
| special   | boolean | Include special characters     | false   |

#### Example Response
\```json
{
    "password": "Kj#9mP$2nL"
}
\```

## 📁 Project Structure
\```
password-generator/
├── main.go
├── handlers/
│   └── handlers.go
├── utils/
│   └── password.go
├── templates/
│   └── index.html
└── Dockerfile
\```

## 🚢 Deployment

### Docker Compose
\```yaml
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
\```

## 🔒 Security Features
- Input validation
- XSS protection headers
- Rate limiting
- Secure random number generation
- HTTPS support (in production)

## 🔍 Monitoring
- Health checks
- Resource monitoring
- Request logging

## 💡 Configuration Options

### Environment Variables
| Variable      | Description           | Default     |
|---------------|-----------------------|-------------|
| PORT          | Server port          | 8080        |
| ENVIRONMENT   | Runtime environment  | development |

### Resource Configuration
\```yaml
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
\```

## 🤝 Contributing
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Last updated: 2025-02-04 15:10:31 UTC  
Created by: Rockphill

---

**Note**: For production deployment, make sure to:
- Set up proper SSL/TLS certificates
- Configure appropriate security measures
- Set up monitoring and logging
- Review and adjust resource limits

For more information, please refer to the [documentation](docs/README.md).
