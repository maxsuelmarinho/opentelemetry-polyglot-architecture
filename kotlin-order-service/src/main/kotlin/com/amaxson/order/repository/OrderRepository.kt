package com.amaxson.order.repository

import com.amaxson.order.model.Order
import org.springframework.data.jpa.repository.JpaRepository
import java.util.*

interface OrderRepository: JpaRepository<Order, Long> {
  fun findByUserId(userId: String): List<Order>
  fun findByUuid(uuid: String): Optional<Order>
}
