# 🔐 Art Decoder 🎨

## Cyberpunk-Inspired Text Encoding & Decoding Web Application

### 🌟 Project Overview

Art Decoder is a sleek, cyberpunk-themed web application that allows users to encode and decode text with style. Built with Go, HTML, and CSS, this project offers a unique and visually stunning interface for text transformation.

### ✨ Features

- 🎨 **Multiple Themes**
  - Cyberpunk
  - Vaporwave
  - Matrix
- 🔒 **Advanced Encoding/Decoding**
  - Simple and intuitive text transformation
  - Supports both encoding and decoding modes
- 🎉 **Interactive UI**
  - Responsive design
  - Animated glitch effects
  - Theme switching capabilities

### 🚀 Technology Stack

- **Backend**: Go (Golang)
- **Frontend**: HTML5, CSS3
- **Styling**: Custom CSS with dynamic theming
- **Deployment**: Ready for containerization and cloud deployment

### 🛠 Installation

#### Prerequisites

- Go 1.16+
- Git

#### Clone the Repository

```bash
git clone https://gitea.koodsisu.fi/seankipina/art
cd art-decoder
```

#### Change Directories

```bash
cd art-interface
```

#### Run the Application

```bash
# Run with default settings
go run main.go

# Customize port
go run main.go -port=9000

# Enable/Disable debug mode
go run main.go -debug=true
```

### 📝 Notes
To check for the endpoint to see that it is running please use the following command:
```bash
curl -I http://localhost:8080
```
### 🔧 Configuration Flags

| Flag         | Description                  | Default Value |
| ------------ | ---------------------------- | ------------- |
| `-port`      | Server listening port        | 8080          |
| `-animation` | Enable/disable UI animations | true          |
| `-debug`     | Enable debug mode            | false         |
| `-theme`     | Set default theme            | cyberpunk     |

### 📡 API Endpoints

#### Decode Endpoint

- **URL**: `/api/decode`
- **Method**: POST
- **Request Body**:
  ```json
  {
    "input": "Text to decode"
  }
  ```

#### Encode Endpoint

- **URL**: `/api/encode`
- **Method**: POST
- **Request Body**:
  ```json
  {
    "input": "Text to encode"
  }
  ```


### 🎨 Theme Customization

The application supports three unique themes:

1. **Cyberpunk**: Dark, neon-inspired design
2. **Vaporwave**: Retro, pastel color palette
3. **Matrix**: Green, code-like aesthetic

Users can switch themes dynamically using radio buttons.

### 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### 📄 License

Distributed under the MIT License. See `LICENSE` for more information.

### 🌐 Contact

Sean Kipinä - https://www.seankipina.com

Project Link: [https://gitea.koodsisu.fi/seankipina/art](https://gitea.koodsisu.fi/seankipina/art)

---

**Made with 💖 and a lot of caffeine**
