package com.amaxson.order.model

import com.fasterxml.jackson.annotation.JsonProperty
import java.io.Serializable

class ShippingAddress(
  @JsonProperty("address") val address: String,
  @JsonProperty("city") val city: String,
  @JsonProperty("postalCode") val postalCode: String,
  @JsonProperty("country") val country: String
): Serializable
