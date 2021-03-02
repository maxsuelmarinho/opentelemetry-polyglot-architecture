package com.amaxson.order.dto

import java.math.BigDecimal
import javax.validation.constraints.Min
import javax.validation.constraints.NotBlank

data class OrderItemDTO(
  @field:NotBlank
  val name: String,
  @field:Min(value = 1)
  val qty: Int,
  @field:NotBlank
  val image: String,
  val price: BigDecimal,
  @field:NotBlank
  val product: String,
)
