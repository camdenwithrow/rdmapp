### rdmapp - Roadmap Feature Management

---

#### Overview
rdmapp is a  web application built to manage feature requests and customer feedback using a streamlined roadmap board. It leverages Golang for the backend, HTMX, and TailwindCSS with DaisyUI for a modern, responsive UI.

---

#### Installation

To install and run rdmapp locally, follow these steps:

1. **Clone the repository:**
   ```bash
   git clone github.com/camdenwithrow/rdmapp.git
   cd rdmapp
   ```

2. **Install npm packages:**
   ```bash
   cd ui
   npm install
   cd ..
   ```

3. **Initialize Go modules:**
   ```bash
   go mod tidy
   ```

4. **Run the application:**
   ```bash
   air
   ```

   The `air` command will start the application server, enabling live reload for development.

---

#### Technologies Used

- **Backend:** Golang
- **Frontend:** HTMX, TailwindCSS, DaisyUI
- **Development:** Air (for live reload)

---

#### Usage

Once the application is running, you can access rdmapp in your web browser at `http://localhost:PORT`.

---

#### Acknowledgments

- Special thanks to the creators of Golang, HTMX, TailwindCSS, DaisyUI, and Air for their fantastic tools and libraries.

---
