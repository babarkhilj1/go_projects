# Go OTP Verification with Twilio

A simple implementation of OTP verification using Twilio.

## Setup
 
1. **Set up Twilio**:  
   - Create a Twilio account to obtain:
     - Account SID  
     - Auth Token  
     - Verify Service SID  
2. **Configure Environment Variables**:  
   Create a `.env` file in the project root:
   ```env
   TWILIO_ACCOUNT_SID=<Your_Account_SID>
   TWILIO_AUTHTOKEN=<Your_Auth_Token>
   TWILIO_SERVICES_ID=<Your_Service_SID>
   ```
3. **Install Dependencies**:  
   ```bash
   go mod download
   ```

## Running the Server
Start the server:  
```bash
go run cmd/main.go
```

## API Endpoints

### 1. Send OTP
- **Endpoint**: `POST /otp`
- **Request Body**:
  ```json
  {
    "phoneNumber": "<phone-number-with-country-code>"
  }
  ```
- **Example cURL**:
  ```bash
   Invoke-WebRequest -Uri http://localhost:8000/otp `
     -Method Post `
     -Headers @{"Content-Type"="application/json"} `
     -Body '{"phoneNumber": "+919351146236"}'
  ```
- **Response**:
  ```json
  {
    "status": 202,
    "message": "success",
    "data": "OTP sent successfully"
  }
  ```

### 2. Verify OTP
- **Endpoint**: `POST /verifyOTP`
- **Request Body**:
  ```json
  {
    "user": {
      "phoneNumber": "<phone-number-with-country-code>"
    },
    "code": "<OTP-Code>"
  }
  ```
- **Example cURL**:
  ```bash
  Invoke-WebRequest -Uri http://localhost:8000/verifyOTP `
     -Method Post `
     -Headers @{"Content-Type"="application/json"} `
     -Body '{"user": {"phoneNumber": "+919351146236"}, "code":"632175"}'
  ```
- **Response**:
  ```json
  {
    "status": 202,
    "message": "success",
    "data": "OTP verified successfully"
  }
  ```
