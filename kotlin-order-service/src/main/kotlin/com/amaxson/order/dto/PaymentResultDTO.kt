package com.amaxson.order.dto

import com.fasterxml.jackson.annotation.JsonAlias
import com.fasterxml.jackson.annotation.JsonProperty
import com.fasterxml.jackson.annotation.JsonUnwrapped
import com.fasterxml.jackson.databind.JsonNode
import com.fasterxml.jackson.databind.annotation.JsonDeserialize
import javax.validation.constraints.NotBlank
import javax.validation.constraints.NotNull

data class PaymentResultDTO(
  @field:NotBlank
  @JsonProperty("id") val id: String,

  @field:NotBlank
  @JsonProperty("status") val status: String,

  @field:NotBlank
  @JsonProperty("update_time") val updateTime: String,

  @field:NotNull
  @JsonProperty("payer")
  val payer: Map<String, JsonNode>
)
