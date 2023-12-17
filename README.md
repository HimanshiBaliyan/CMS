# Customer Management System (CMS) Project

## Overview

The Customer Management System (CMS) is a comprehensive project designed to facilitate efficient management of customer data for businesses. This system provides a set of APIs for creating, reading, updating, and deleting customer information. The CMS is built using the Go programming language and the GoFr framework, offering a robust and scalable solution for customer data management.

## Features

- **CRUD Operations:** The system supports standard CRUD operations, allowing users to create, read, update, and delete customer records.

- **Database Integration:** CMS integrates with a MySQL database for persistent storage of customer data. It automates tasks such as connection pooling and retry mechanisms.

- **Unit Tests:** The project maintains a minimum of 60% unit test coverage to ensure the reliability and stability of the codebase.

- **Real-World Scenario:** The CMS project is designed around a real-world scenario, mimicking a customer database where businesses can add, view, update, and remove customer records.

## Use Cases

## 1. Add Customer

### Objective
Allow businesses to add new customer records to the system.

### Steps
1. The business sends a POST request to the `/customer` endpoint with the necessary customer details.
2. The CMS validates the incoming data and inserts the new customer record into the database.
3. A success response is returned to the business, confirming the addition of the customer.

---

## 2. View Customer List

### Objective
Enable businesses to retrieve a list of all existing customers.

### Steps
1. The business sends a GET request to the `/customer` endpoint.
2. The CMS queries the database for the list of customers.
3. The list of customers is returned to the business as a JSON response.

---

## 3. Update Customer Details

### Objective
Allow businesses to update the details of an existing customer.

### Steps
1. The business sends a PUT request to the `/customer/{id}` endpoint, specifying the customer ID to be updated.
2. The CMS retrieves the existing customer details from the database.
3. The system updates the customer information based on the provided data.
4. A success response is returned to the business, confirming the update.

---

## 4. Remove Customer

### Objective
Enable businesses to remove a customer record from the system.

### Steps
1. The business sends a DELETE request to the `/customer/{id}` endpoint, specifying the customer ID to be removed.
2. The CMS checks if the customer exists in the database.
3. If the customer is found, the system deletes the customer record.
4. A success response is returned to the business, confirming the removal of the customer.

---

These use cases provide a comprehensive overview of the essential operations supported by the Customer Management System (CMS). Businesses can seamlessly interact with the system to manage customer data efficiently.

## Getting Started

To get started with the Customer Management System, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/customer-management-system.git
   ```

2. Navigate to the project directory:
   ```bash
   cd customer-management-system
   ```

3. Initialize the go module:
   ```bash
   go mod init github.com/your-username/customer-management-system
   ```

4. Install dependencies:
   ```bash
   go mod tidy
   ```

5. Set up the database:
   - Use the provided Docker command to run a MySQL instance and create the necessary database and table.

6. Run the server:
   ```bash
   go run main.go
   ```

7. Access the API at http://localhost:9000

## Project Structure

The project adheres to a clean and organized structure:

```
.
├── main.go
├── go.mod
├── go.sum
├── configs
│   └── .env
```

...

## Real-World Scenario

The CMS project is designed to simulate a scenario where businesses need to manage customer information. It provides a practical implementation of a Customer Management System, allowing users to add new customers, view the list of existing customers, update customer details, and remove customers from the database.

## Conclusion

The Customer Management System project demonstrates the capabilities of the GoFr framework for building robust and scalable APIs. With its focus on real-world scenarios, database integration, and unit testing, the CMS project serves as a valuable template for businesses seeking to implement a customer data management solution.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Feel free to customize the report further based on the specific features, structure, and goals of your Customer Management System project.
