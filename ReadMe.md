# 🍴 Restaurant API (Go)

A clean and modular **Restaurant API** built with **Go**, **Gorm** and **Gin** , following best practices with **Repositories**, **Services**, **Handlers**, and **Caching**.  
This API provides endpoints for managing users, menus, orders, and order items in a restaurant system. It also provides, authentication, authorization and protected routes

---

## ⚙️ Features

- ✅ **Users**:
  - Create and manage users.
  - Get current logged in users.
- ✅ **Authentication**:
  - JWT Based login system
  - Password hashing & verification
- ✅ **Authorization**:
  - RBAC- Role Based Access Control
  - Middleware to protect endpoints
- ✅ **Menu**: List menu items.
- ✅ **Orders**: Create, fetch, and update orders.
- ✅ **Order Items**: Attach items to an order.
- ✅ **Caching**: Reduce DB load with an in-memory cache layer.
- ✅ **Clean Architecture**: Repositories → Services → Handlers separation → Routers.
- ✅ **Environment Variables** with `.env`.

---
