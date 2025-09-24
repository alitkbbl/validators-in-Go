# User Validation API

A simple RESTful service built with the **Gin** framework in **Go** that validates user data based on custom business rules.

---

## 📦 Features

- Validates user information using structured rules
- Custom validation for:
  - Iranian National ID (کد ملی)
  - Iranian phone number (must start with `09`)
  - Birth date (age between 7 and 100 years)
  - Unique username and email
- Supports multipart/form-data requests
- Fully tested with unit tests

---

# User Validation API

A simple RESTful service built with the **Gin** framework in **Go** that validates user data based on custom business rules.

---

## 📦 Features

- Validates user information using structured rules
- Custom validation for:
  - Iranian National ID (کد ملی)
  - Iranian phone number (must start with `09`)
  - Birth date (age between 7 and 100 years)
  - Unique username and email
- Supports multipart/form-data requests
- Fully tested with unit tests

---

## 🛠️ Installation & Setup

### Prerequisites
- Go 1.21 or higher

### Steps
1. Clone the repository:
```bash
git clone https://github.com/your-username/your-repo-name.git
cd your-repo-name
```
2. Install dependencies:

bash
```
go mod download

```

Run the application:
bash
```
go run main.go validators.go types.go
```
The server will start at http://localhost:8080

---

### 🚀 Usage Example

Using curl:
```bash
curl -X POST http://localhost:8080/user/validate \
  -F "first_name=ali" \
  -F "last_name=alavi" \
  -F "username=alialavi" \
  -F "email=alavi@quera.org" \
  -F "phone_number=09123456789" \
  -F "birth_date=2000/01/01" \
  -F "national_id=7421368515"
```
