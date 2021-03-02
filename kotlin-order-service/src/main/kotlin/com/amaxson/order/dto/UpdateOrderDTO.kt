package com.amaxson.order.dto

import javax.validation.Valid
import javax.validation.constraints.NotBlank

data class UpdateOrderDTO(
  var orderId: String,
  @field:Valid
  val paymentResult: PaymentResultDTO
)
