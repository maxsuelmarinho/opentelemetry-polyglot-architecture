package com.amaxson.order.dto

import java.math.BigDecimal
import javax.validation.Valid
import javax.validation.constraints.NotBlank
import javax.validation.constraints.NotEmpty
import javax.validation.constraints.Size

data class OrderDTO(
  @field:NotBlank
  val userId: String,
  @field:NotEmpty
  @field:Valid
  val orderItems: List<@Valid OrderItemDTO> = mutableListOf(),
  val shippingAddress: ShippingAddressDTO,
  @field:NotBlank
  val paymentMethod: String,
  val taxPrice: BigDecimal,
  val shippingPrice: BigDecimal,
  val totalPrice: BigDecimal
)
