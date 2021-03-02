package com.amaxson.order.mapper

import com.amaxson.order.dto.OrderDTO
import com.amaxson.order.dto.OrderItemDTO
import com.amaxson.order.model.Order
import com.amaxson.order.model.OrderItem
import org.modelmapper.AbstractConverter
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Component
import java.util.*

@Component
class OrderDTOToOrder: AbstractConverter<OrderDTO, Order>() {
  @Autowired
  lateinit var orderItemDTOToOrderItem: OrderItemDTOToOrderItem
  @Autowired
  lateinit var shippingAddressDTOToShippingAddress: ShippingAddressDTOToShippingAddress

  override fun convert(source: OrderDTO): Order {

    val order = Order(
      uuid = UUID.randomUUID().toString(),
      isPaid = false,
      isDelivered = false,
      paymentMethod = source.paymentMethod,
      taxPrice = source.taxPrice,
      shippingPrice = source.shippingPrice,
      totalPrice = source.totalPrice,
      userId = source.userId,
      shippingAddress = shippingAddressDTOToShippingAddress.convert(source.shippingAddress)
    )

    val orderItems: List<OrderItem> = source.orderItems.map {
      val orderItem: OrderItem = orderItemDTOToOrderItem.convert(it)
      orderItem.order = order
      orderItem
    }
    order.orderItems = orderItems

    return order
  }
}
