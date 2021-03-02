package com.amaxson.order.service

import com.amaxson.order.exception.OrderNotFoundException
import com.amaxson.order.model.Order
import com.amaxson.order.model.PaymentResult
import com.amaxson.order.repository.OrderRepository
import org.slf4j.Logger
import org.slf4j.LoggerFactory
import org.springframework.stereotype.Service
import java.time.LocalDateTime

@Service
class OrderService(private var repository: OrderRepository) {
  companion object {
    val logger: Logger = LoggerFactory.getLogger(OrderService::class.java)
  }

  fun getOrders(userId: String): List<Order> {
    val orders = repository.findByUserId(userId)
    logger.info("orders: ${orders}")
    return orders
  }

  fun createOrder(order: Order): Order {
    val createdOrder = repository.save(order)
    logger.info("order created: ${createdOrder}")
    return createdOrder
  }

  fun getOrderById(uuid: String): Order {
    val order = repository.findByUuid(uuid)
    if (!order.isPresent) {
      throw OrderNotFoundException()
    }

    return order.get()
  }

  fun updateOrderToPaid(uuid: String, paymentResult: PaymentResult): Order {
    val order = getOrderById(uuid)
    order.isPaid = true
    order.paidAt = LocalDateTime.now()
    order.paymentResult = paymentResult
    return repository.save(order)
  }
}
