# React.js E-commerce example

```shell
$ npx create-react-app frontend
$ npm i react-bootstrap
$ npm i react-router-dom react-router-bootstrap
```

## REST API

### Available endpoints

1. [List products](#list-products)
2. [Get product details](#get-product-details)
3. [User login](#user-login)
4. [User profile](#user-profile)
5. [Update user profile](#update-user-profile)
6. [Register user](#register-user)
7. [Place Order](#place-order)
8. [Get order by ID](#get-order-by-id)

Método | Caminho | Descrição
--- | --- | ---
GET | /api/products | List products
GET | /api/products/{productId} | Get product details
POST | /api/users/login | User login
GET | /api/users/profile | User profile
UPDATE | /api/users/profile | Update user profile
POST | /api/users | Register user
POST | /api/orders | Place order
GET | /api/orders/{orderId} | Get order by ID
POST | /api/product/{productId}/reviews | Create product review

### List products

```shell
$ curl -i -X GET http://localhost:5000/api/products
```

### Get product details

```shell
$ curl -i -X GET http://localhost:5000/api/products/602286e00b423f077f9a062c
```

### User Login

```shell
$ curl -i -H 'Content-Type: application/json' -X POST http://localhost:5000/api/users/login -d '{"email":"john@example.com","password":"123456"}'
```

### User Profile

```shell
$ curl -i -H 'Content-Type: application/json' -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjYwMjI4NmUwMGI0MjNmMDc3ZjlhMDYyNiIsImlhdCI6MTYxMzAxODQ2MCwiZXhwIjoxNjE1NjEwNDYwfQ.RuHxDw5yS7XXI-eUHTFciQCz6NlZZNca8JEM_wtO8_M' -X GET http://localhost:5000/api/users/profile
```

### Update User Profile

```shell
$ curl -i -H 'Content-Type: application/json' -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjYwMjI4NmUwMGI0MjNmMDc3ZjlhMDYyNiIsImlhdCI6MTYxMzAxODQ2MCwiZXhwIjoxNjE1NjEwNDYwfQ.RuHxDw5yS7XXI-eUHTFciQCz6NlZZNca8JEM_wtO8_M' -X POST http://localhost:5000/api/users/profile -d '
{
  "name": "Luke",
  "email": "luke@example.com",
  "password": "123456"
}'
```

### Register user

```shell
$ curl -i -H 'Content-Type: application/json' http://localhost:5000/api/users -d '
{
  "name": "Luke",
  "email": "luke@example.com",
  "password": "123456"
}'
```

### Place Order

```shell
$ curl -i -H 'Content-Type: application/json' -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjYwMjI4NmUwMGI0MjNmMDc3ZjlhMDYyNiIsImlhdCI6MTYxMzAxODQ2MCwiZXhwIjoxNjE1NjEwNDYwfQ.RuHxDw5yS7XXI-eUHTFciQCz6NlZZNca8JEM_wtO8_M' -X POST http://localhost:5000/api/orders -d '
{
  "orderItems": [
    {
      "product": "602286e00b423f077f9a0628",
      "name": "Airpods Wireless Bluetooth Headphones",
      "image": "/images/airpods.jpg",
      "price": 89.99,
      "countInStock": 10,
      "qty": 2
    }
  ],
  "shippingAddress": {
    "address": "Paulo Barreto",
    "city": "Rio de Janeiro",
    "postalCode": "22280010",
    "country": "Brazil"
  },
  "paymentMethod": "PayPal",
  "itemsPrice": "179.98",
  "shippingPrice": "0.00",
  "taxPrice": "27.00",
  "totalPrice": "206.98"
}'
```

### Create product review

```shell
$ curl -i -H 'Content-Type: application/json' -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjYwMjI4NmUwMGI0MjNmMDc3ZjlhMDYyNiIsImlhdCI6MTYxMzAxODQ2MCwiZXhwIjoxNjE1NjEwNDYwfQ.RuHxDw5yS7XXI-eUHTFciQCz6NlZZNca8JEM_wtO8_M' -X POST http://localhost:5000/api/products/602286e00b423f077f9a062c/reviews -d '
{
  "rating": 5,
  "comment": "Awesome headphones"
}'
```

## Bugs

- Cart Screen:
    - When remove the the last added item and then refresh the page, the item came back into the list

## References

- https://jwt.io
- https://bootswatch.com/
- https://react-bootstrap.github.io/getting-started/introduction/
- https://cdnjs.com/libraries?q=font