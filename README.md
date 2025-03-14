# money-transfer-system

```markdown
# Simple Concurrent Money Transfer System

This Go-based system allows users to transfer money between each other, ensuring atomic updates to balances to prevent race conditions and preventing overdrafts. The system is designed with concurrency safety in mind and provides a simple HTTP API to facilitate money transfers between users.

## Requirements
- **Go version**: Go 1.18 or higher
- **Concurrency-safe data structures** to manage user balances.
- **Error Handling**: Covers basic error cases like invalid users, insufficient funds, and transfers to the sender's own account.

## System Overview
The system allows the following functionality:
- Transfer money between users (`POST /transfer`).
- Retrieve the current balance of a user (`GET /balance`).

The initial account balances are as follows:
- Mark: $100
- Jane: $50
- Adam: $0

### Example API Usage

1. **Check User Balance**:  
   Retrieve the current balance of a user.
   ```bash
   curl -X GET http://localhost:8080/balance?user=Mark
   ```
   **Response**:
   ```json
   {"balance": 70, "user": "Mark"}
   ```

2. **Transfer Money**:  
   Initiate a money transfer between two users.
   ```bash
   curl -X POST "http://localhost:8080/transfer" -H "Content-Type: application/json" -d '{
       "from": "Mark",
       "to": "Jane",
       "amount": 30
   }'
   ```
   **Response**:
   ```json
   {"message": "Transfer successful"}
   ```

## Setting Up the System

### Prerequisites
1. **Go**: Ensure that Go is installed on your machine. You can download it from [here](https://go.dev/dl/).
2. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/money-transfer-system.git
   cd money-transfer-system
   ```

### Install Dependencies
Run the following command to install any necessary Go dependencies:
```bash
go mod tidy
```

### Run the Server
Start the Go application with the following command:
```bash
go run main.go
```
The server will start on `http://localhost:8080`.

### Run Tests
Ensure that everything is working by running the tests:
```bash
go test -v ./test
```

## Design and Concurrency Strategy

The system uses a **map** to store user balances, where the user name is the key and the balance is the value. To ensure atomicity and avoid race conditions, we use a **sync.Mutex** to lock the map whenever a balance is being updated. This ensures that only one transfer is processed at a time, avoiding conflicts or inconsistencies when multiple transfers are initiated concurrently.

### Locking Strategy
1. **Mutex per user**: Each user has a mutex associated with them to ensure that balance updates for each user are done atomically. The lock is held during the transfer process to ensure no other transfer operations can modify the user’s balance until the current operation completes.
2. **Global mutex**: A global mutex (`sync.Mutex`) is used to ensure no two transfers can happen simultaneously.

### Example Code Structure

- **main.go**: The entry point of the program. Contains the HTTP server and handlers for API endpoints.
- **transfer.go**: Contains the logic for money transfer, including validation and balance updates.
- **user.go**: Defines the user structure, holding the user's balance and their associated mutex.

### Error Handling
The system handles various error cases:
- **Invalid User**: If the user does not exist, the system returns an appropriate error message.
- **Insufficient Funds**: If the user does not have enough money to transfer, the system returns an error.
- **Transfer to Self**: The system prevents users from transferring money to their own account.


## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```



