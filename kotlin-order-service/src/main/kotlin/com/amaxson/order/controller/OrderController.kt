package com.amaxson.order.controller

import com.amaxson.order.dto.OrderDTO
import com.amaxson.order.dto.PaymentResultDTO
import com.amaxson.order.dto.UpdateOrderDTO
import com.amaxson.order.model.Order
import com.amaxson.order.model.PaymentResult
import com.amaxson.order.service.OrderService
import org.modelmapper.ModelMapper
import org.slf4j.Logger
import org.slf4j.LoggerFactory
import org.springframework.http.HttpStatus
import org.springframework.http.ResponseEntity
import org.springframework.validation.annotation.Validated
import org.springframework.web.bind.annotation.*
import org.springframework.web.servlet.support.ServletUriComponentsBuilder
import java.net.URI
import javax.validation.Valid
import javax.validation.constraints.NotBlank

@RestController
@RequestMapping("/api/orders")
@Validated
class OrderController(private var service: OrderService, private var mapper: ModelMapper) {
  companion object {
    val logger: Logger = LoggerFactory.getLogger(OrderController::class.java)
  }

  @PostMapping
  fun createOrder(@Valid @RequestBody order: OrderDTO): ResponseEntity<Any> {
    logger.info("order: ${order}")

    val createdOrder = service.createOrder(mapper.map(order, Order::class.java))
    val location: URI = ServletUriComponentsBuilder.fromCurrentRequest()
      .path("/{id}")
      .buildAndExpand(createdOrder.uuid).toUri()

    return ResponseEntity.created(location).body(createdOrder)
  }

  @GetMapping
  fun listOrders(@RequestParam user: String): List<Order> {
    logger.debug("list user ${user} orders")
    val orders = service.getOrders(user)
    logger.debug("total: ${orders.size} orders: ${orders}")
    return orders
  }

  @GetMapping("{id}")
  fun getOrderDetails(@PathVariable id: String): ResponseEntity<Order> {
    val order = service.getOrderById(id)
    return ResponseEntity.ok(order)
  }

  @PutMapping("{id}/pay")
  fun updateOrderToPaid(@PathVariable id: String, @Valid @RequestBody dto: PaymentResultDTO): ResponseEntity<Order> {
    val paymentResult = mapper.map(dto, PaymentResult::class.java)
    val updatedOrder = service.updateOrderToPaid(id, paymentResult)
    return ResponseEntity.ok(updatedOrder)
  }
}
