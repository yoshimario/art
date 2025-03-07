------

# 🔐 Art Decoder 🎨

## Cyberpunk-Inspired Text Encoding & Decoding Web Application

### 🌟 Project Overview

Art Decoder is a sleek, cyberpunk-themed web application that allows users to encode and decode text with style. Built with **Go (Golang)** for the backend and **HTML5/CSS3** for the frontend, this project offers a unique and visually stunning interface for text transformation.

------

## ✨ Features

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

------

## 🚀 Technology Stack

- **Backend**: Go (Golang)
- **Frontend**: HTML5, CSS3
- **Styling**: Custom CSS with dynamic theming
- **Deployment**: Ready for containerization and cloud deployment

------

## 🛠 Installation

### Prerequisites

Before running the Art Decoder application, ensure you have the following installed:

1. **Go (Golang)**: Version 1.16 or higher.

   - Download and install Go from https://golang.org/dl/.

   - Verify the installation by running:

     bash

     ```
     go version
     ```

   - If Go is installed correctly, you'll see something like `go version go1.20.3`. If not, revisit the installation steps.

2. **Git**: To clone the repository.

   - Download and install Git from https://git-scm.com/.

   - Verify the installation by running:

     bash

     ```
     git --version
     ```

   - If Git is installed correctly, you'll see something like `git version 2.40.0`. If not, revisit the installation steps.

------

### Step 1: Clone the Repository

1. Open a terminal (Command Prompt on Windows, Terminal on macOS/Linux).

2. Run the following command to download the Art Decoder project:

   bash

   ```
   git clone https://gitea.koodsisu.fi/seankipina/art
   ```

3. Navigate to the project folder:

   bash

   ```
   cd art-decoder/art-interface
   ```

------

### Step 2: Run the Application

1. Start the server by running:

   bash

   ```
   go run main.go
   ```

   - This will start the Art Decoder server on port `8080` by default.
   - **Do not close this terminal window**—it keeps the server running.

2. Open your web browser and go to:

   ```
   http://localhost:8080
   ```

   - If everything is working, you should see the Art Decoder web interface.

------

### Step 3: Verify the Server is Running (Optional)

To verify that the server is running, you can use a second terminal window:

1. Open a **new terminal window** (don’t close the first one where the server is running).

2. Run the following command to check the server status:

   bash

   ```
   curl -I http://localhost:8080
   ```

   - If the server is running correctly, you’ll see a response like this:

     ```
     HTTP/1.1 200 OK
     ```

   - If you see an error, make sure the server is running in the first terminal window.

------

### Step 4: Customize the Server (Optional)

You can customize how the server runs using command-line flags. Here are some examples:

| Flag         | Description            | Example Command                   |
| :----------- | :--------------------- | :-------------------------------- |
| `-port`      | Change the server port | `go run main.go -port=9000`       |
| `-animation` | Disable animations     | `go run main.go -animation=false` |
| `-debug`     | Enable debug mode      | `go run main.go -debug=true`      |
| `-theme`     | Set the default theme  | `go run main.go -theme=vaporwave` |

For example, to run the server on port `9000` with debug mode enabled:

bash

```
go run main.go -port=9000 -debug=true
```

------

### Step 5: Stop the Server

To stop the server, go back to the terminal where it’s running and press:

- **Ctrl + C** (on Windows, macOS, or Linux).

------

## 📡 API Endpoints

The Art Decoder application provides two main API endpoints for encoding and decoding text.

### Decode Endpoint

- **URL**: `/api/decode`

- **Method**: POST

- **Request Body**:

  json

  ```
  {
    "input": "Text to decode"
  }
  ```

### Encode Endpoint

- **URL**: `/api/encode`

- **Method**: POST

- **Request Body**:

  json

  ```
  {
    "input": "Text to encode"
  }
  ```

------

## 🎨 Theme Customization

The application supports three unique themes:

1. **Cyberpunk**: Dark, neon-inspired design
2. **Vaporwave**: Retro, pastel color palette
3. **Matrix**: Green, code-like aesthetic

You can switch themes dynamically using the radio buttons in the UI.

------

## 🤝 Contributing

We welcome any contributions! If you'd like to contribute to the Art Decoder project, follow these steps:

1. Fork the repository.

2. Create a new branch for your feature:

   bash

   ```
   git checkout -b feature/AmazingFeature
   ```

3. Commit your changes:

   bash

   ```
   git commit -m 'Add some AmazingFeature'
   ```

4. Push your changes to the branch:

   bash

   ```
   git push origin feature/AmazingFeature
   ```

5. Open a Pull Request on the original repository.

------

## 📄 License

This project is distributed under the **MIT License**. See the [LICENSE](https://license/) file for more details.

------

## 🌐 Contact

If you have any questions or feedback, feel free to reach out:

- **Sean Kipinä**: [https://www.seankipina.com](https://www.seankipina.com/)
- **Project Link**: https://gitea.koodsisu.fi/seankipina/art

------

**Made with 💖 and a lot of caffeine**
