package com.amaxson.order.mapper

import com.amaxson.order.dto.ShippingAddressDTO
import com.amaxson.order.model.ShippingAddress
import org.modelmapper.AbstractConverter
import org.springframework.core.convert.converter.Converter
import org.springframework.stereotype.Component

@Component
class ShippingAddressDTOToShippingAddress: Converter<ShippingAddressDTO, ShippingAddress>, AbstractConverter<ShippingAddressDTO, ShippingAddress>() {
  override fun convert(source: ShippingAddressDTO): ShippingAddress {
    return ShippingAddress(
      address = source.address,
      city = source.city,
      postalCode = source.postalCode,
      country = source.country
    )
  }
}
