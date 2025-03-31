# 💰 **Expense Tracker**  
A simple command-line expense tracker that stores data in the filesystem.  

## 🚀 **Features**  

🔗 [Project Roadmap](https://roadmap.sh/projects/expense-tracker)  

- ✅ **Add an expense** ➝ Description, amount, and category.  
- ❌ **Update an expense** (🚧 Not yet implemented).  
- ✅ **Delete an expense** by ID.  
- ✅ **View all expenses** in a list format.  
- ✅ **View a summary** of all expenses (💵 Total amount spent).  
- ❌ **View a summary for a specific month** (🚧 Not yet implemented).  

---

## ⚙️ **Usage**  

### 🔨 **Build the Program**  

```sh
./build.sh
```

### ✏️ **Adding an Expense**

```sh
./bin/expense-tracker add --description="Lunch" --amount 150.00 --category="Food"
```

### 📋 **List All Expenses**

```sh
./bin/expense-tracker list
```

### 📊 **Summary Total Expenses**

```sh
./bin/expense-tracker summary
```

### 🗑️ **Delete an Expense**

```sh
./bin/expense-tracker delete --id 1
```
