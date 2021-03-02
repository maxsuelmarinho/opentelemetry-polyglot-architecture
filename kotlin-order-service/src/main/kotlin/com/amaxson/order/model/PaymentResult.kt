package com.amaxson.order.model

import com.fasterxml.jackson.annotation.JsonProperty
import java.io.Serializable

class PaymentResult(
    @JsonProperty("id") val id: String,
    @JsonProperty("status") val status: String,
    @JsonProperty("update_time") val updateTime: String,
    @JsonProperty("email_address") val emailAddress: String
): Serializable

