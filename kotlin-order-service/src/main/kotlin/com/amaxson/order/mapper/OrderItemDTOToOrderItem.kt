package com.amaxson.order.mapper

import com.amaxson.order.dto.OrderItemDTO
import com.amaxson.order.model.OrderItem
import org.modelmapper.AbstractConverter
import org.springframework.core.convert.converter.Converter
import org.springframework.stereotype.Component
import java.util.*

@Component
class OrderItemDTOToOrderItem: Converter<OrderItemDTO, OrderItem>, AbstractConverter<OrderItemDTO, OrderItem>() {

  override fun convert(source: OrderItemDTO): OrderItem {
    return OrderItem(
      uuid = UUID.randomUUID().toString(),
      name = source.name,
      qty = source.qty,
      image = source.image,
      price = source.price,
      productId = source.product
    )
  }
}
