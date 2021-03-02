package com.amaxson.order.dto

data class ShippingAddressDTO(
  val address: String,
  val city: String,
  val postalCode: String,
  val country: String,
)
