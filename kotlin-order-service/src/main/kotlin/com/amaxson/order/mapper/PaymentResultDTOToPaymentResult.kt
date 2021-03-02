package com.amaxson.order.mapper

import com.amaxson.order.dto.PaymentResultDTO
import com.amaxson.order.dto.ShippingAddressDTO
import com.amaxson.order.model.PaymentResult
import com.amaxson.order.model.ShippingAddress
import org.modelmapper.AbstractConverter
import org.springframework.core.convert.converter.Converter
import org.springframework.stereotype.Component

@Component
class PaymentResultDTOToPaymentResult: Converter<PaymentResultDTO, PaymentResult>, AbstractConverter<PaymentResultDTO, PaymentResult>() {
  override fun convert(source: PaymentResultDTO): PaymentResult {
    return PaymentResult(
      id = source.id,
      status = source.status,
      updateTime = source.updateTime,
      emailAddress = source.payer["email_address"]?.asText() ?: ""
    )
  }
}
