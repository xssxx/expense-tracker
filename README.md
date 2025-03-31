# **Expense Tracker**

A simple command-line expense tracker that stores data in the filesystem.

## **Features**

https://roadmap.sh/projects/expense-tracker

- ✅ **Add an expense** with a description, amount, and category.
- ❌ **Update an expense** (Not yet implemented).
- ✅ **Delete an expense** by ID.
- ✅ **View all expenses** in a list format.
- ✅ **View a summary** of all expenses (total amount spent).
- ❌ **View a summary of expenses for a specific month** (Not yet implemented).

## **Usage**

### **Build Program**

```sh
./build.sh
```

### **Adding an Expense**

```sh
./bin/expense-tracker add --description="Lunch" --amount 150.00 --category="Food"
```

### **List All Expenses**

```sh
./bin/expense-tracker list
```

### **Summary Total Expenses**

```sh
./bin/expense-tracker summary
```

### **Delete an Expense**

```sh
./bin/expense-tracker delete --id 1
```
