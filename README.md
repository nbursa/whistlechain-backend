# **WhistleChain Backend**
📢 **Anonymous Whistleblowing Platform Using Blockchain**  

WhistleChain is a **secure and anonymous** reporting platform designed for **journalists, employees, and citizens** to report unethical activities. It leverages **blockchain** for **tamper-proof records** and ensures **data integrity and confidentiality**.

---

## **🚀 Features**
- **🔒 Secure Report Submission** – Encrypts and stores reports securely.
- **📜 Immutable Record-Keeping** – Uses **blockchain hashing** to ensure report integrity.
- **🛡️ Anonymity & Privacy** – Ensures whistleblowers' safety.
- **📦 PostgreSQL Storage** – Structured data persistence.
- **⚡ Built with Go & Fiber** – High-performance backend.

---

## **🛠 Tech Stack**
- **Backend:** Go, Fiber 🚀  
- **Database:** PostgreSQL  
- **Blockchain:** Cryptographic Hashing  
- **Cloud Deployment:** Heroku / Vercel (Planned)  

---

## **📄 API Endpoints**
| Method | Endpoint         | Description  |
|--------|-----------------|--------------|
| **GET**  | `/` | Check if the API is running |
| **POST** | `/report` | Submit a new encrypted report |
| **GET**  | `/reports` | Get all submitted reports |
| **GET**  | `/report/:id` | Get a specific report by ID |

---

## **🔧 Setup & Run Locally**
### **1️⃣ Clone the Repository**
```sh
git clone https://github.com/nbursa/whistlechain-backend.git
cd whistlechain-backend
```
### **2️⃣ Install Dependencies**
```sh
go mod tidy
```
### **3️⃣ Set Up Environment Variables**
Create a `.env` file and add:
```
DATABASE_URL=postgres://your_user:your_password@localhost:5432/your_db
PORT=3000
```
### **4️⃣ Run the Server**
```sh
go run main.go
```
The API will be available at **`http://localhost:3000`** 🚀.

---

## **📌 Future Roadmap**
- ✅ Implement cryptographic signatures for authenticity.
- ✅ Add JWT-based authentication for company admins.
- 🚧 Deploy to cloud.
- 🚧 Build a front-end for easy report submission.
- 🚧 Extend blockchain integration.

---

## **🤝 Contributing**
Want to help improve **WhistleChain**?  
- **Fork & Star** this repo ⭐
- Open **issues & PRs** with improvements.

---

## **📜 License**
**MIT License** – Free to use & modify.

---

## **💡 Author**
👨‍💻 **Nenad Bursać**  
🚀 GitHub: [nbursa](https://github.com/nbursa)  
💼 LinkedIn: [Nenad Bursać](https://www.linkedin.com/in/nenadbursac/)  
