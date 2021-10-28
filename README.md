# React.js E-commerce example

```shell
$ npx create-react-app frontend
$ npm i react-bootstrap
$ npm i react-router-dom react-router-bootstrap
```

## Getting Started

1. Run the services:
    ```shell
    $ docker-compose up -d
    ```
### Services

- [Amaxson Frontend in React.js](http://localhost:3000)
- [BFF in Node.js](http://localhost:8000)
- [Product Service in Golang](http://localhost:8090/health)
- [Order Service in Kotlin](http://localhost:8080/actuator/health)
- [Legacy Back-end in Node.js](http://localhost:5000)
- [Jaeger Query UI](http://localhost:16686)
- [Zipkin](http://localhost:9411)
- [Prometheus](http://localhost:9090)
- [Kibaba](http://localhost:5601)

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
9. [Get user's orders](#get-user-orders)
10. [Update order to paid](#update-order-to-paid)

Método | Caminho | Descrição
--- | --- | ---
GET | /api/products | List products
GET | /api/products/{productId} | Get product details
GET | /api/products/top | Get top products
POST | /api/users/login | User login
GET | /api/users/profile | User profile
UPDATE | /api/users/profile | Update user profile
POST | /api/users | Register user
POST | /api/orders | Place order
GET | /api/orders/{orderId} | Get order by ID
GET | /api/orders/myorders | Get user's orders
PUT | /api/orders/{orderId}/pay | Update order to paid
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
$ curl -i -H 'Content-Type: application/json' -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjYwMjI4NmUwMGI0MjNmMDc3ZjlhMDYyNiIsImlhdCI6MTYxMzAxODQ2MCwiZXhwIjoxNjE1NjEwNDYwfQ.RuHxDw5yS7XXI-eUHTFciQCz6NlZZNca8JEM_wtO8_M' -X POST -d '
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
}' http://localhost:8080/api/orders
```

### Get order by id

```shell
$ curl -X GET http://localhost:8080/api/orders/ea2b2588-4291-450a-9825-73592f8e31b9
```

### Get user's orders

```shell
# bff
$ curl -X GET -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjYwMjI4NmUwMGI0MjNmMDc3ZjlhMDYyNiIsImlhdCI6MTYxNDY1Mzg1MSwiZXhwIjoxNjE3MjQ1ODUxfQ.SdH3XnZUkkGMSaG_zY1Pc9P0ZrJeL2lO3JtkEWdYmsk' 'http://localhost:8000/api/orders/myorders'

# order-service
$ curl -X GET 'http://localhost:8080/api/orders/?userId=abcde'
```

### Update order to paid

```shell
$ curl -X PUT -d '
{
  "create_time": "2021-03-01T01:51:32Z",
  "update_time": "2021-03-01T01:53:47Z",
  "id": "4XH33102DY969251L",
  "intent": "CAPTURE",
  "status": "COMPLETED",
  "payer": {
    "email_address": "sb-siiyk5126117@personal.example.com",
    "payer_id": "3C67369JYE83N",
    "address": {
      "country_code": "US"
    },
    "name": {
      "given_name": "John",
      "surname": "Doe"
    }
  },
  "purchase_units": [
    {
      "reference_id": "default",
      "soft_descriptor": "PAYPAL *JOHNDOESTES",
      "amount": {
        "value": "689.99",
        "currency_code": "USD"
      },
      "payee": {
        "email_address": "sb-uztjc5128853@business.example.com",
        "merchant_id": "86SRWLNR5PWZA"
      },
      "shipping": {
        "name": {
          "full_name": "John Doe"
        },
        "address": {
          "address_line_1": "1 Main St",
          "admin_area_2": "San Jose",
          "admin_area_1": "CA",
          "postal_code": "95131",
          "country_code": "US"
        }
      },
      "payments": {
        "captures": [
          {
            "status": "PENDING",
            "id": "2K0965801L644061D",
            "final_capture": true,
            "create_time": "2021-03-01T01:53:47Z",
            "update_time": "2021-03-01T01:53:47Z",
            "amount": {
              "value": "689.99",
              "currency_code": "USD"
            },
            "seller_protection": {
              "status": "ELIGIBLE",
              "dispute_categories": [
                "ITEM_NOT_RECEIVED",
                "UNAUTHORIZED_TRANSACTION"
              ]
            },
            "status_details": {
              "reason": "RECEIVING_PREFERENCE_MANDATES_MANUAL_ACTION"
            },
            "links": [
              {
                "href": "https://api.sandbox.paypal.com/v2/payments/captures/2K0965801L644061D",
                "rel": "self",
                "method": "GET",
                "title": "GET"
              },
              {
                "href": "https://api.sandbox.paypal.com/v2/payments/captures/2K0965801L644061D/refund",
                "rel": "refund",
                "method": "POST",
                "title": "POST"
              },
              {
                "href": "https://api.sandbox.paypal.com/v2/checkout/orders/4XH33102DY969251L",
                "rel": "up",
                "method": "GET",
                "title": "GET"
              }
            ]
          }
        ]
      }
    }
  ],
  "links": [
    {
      "href": "https://api.sandbox.paypal.com/v2/checkout/orders/4XH33102DY969251L",
      "rel": "self",
      "method": "GET",
      "title": "GET"
    }
  ]
}' http://localhosts:8080/api/orders/603c48051de00d0ea056ab03/pay
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
- https://snyk.io/blog/10-best-practices-to-containerize-nodejs-web-applications-with-docker/
