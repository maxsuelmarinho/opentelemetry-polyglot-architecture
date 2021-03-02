package com.amaxson.order.configuration

import com.amaxson.order.dto.OrderDTO
import com.amaxson.order.dto.OrderItemDTO
import com.amaxson.order.dto.PaymentResultDTO
import com.amaxson.order.dto.ShippingAddressDTO
import com.amaxson.order.model.Order
import com.amaxson.order.model.OrderItem
import com.amaxson.order.model.PaymentResult
import com.amaxson.order.model.ShippingAddress
import org.modelmapper.Converter
import org.modelmapper.ModelMapper
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.data.jpa.repository.config.EnableJpaAuditing

@Configuration
@EnableJpaAuditing
class OrderConfiguration {

  @Autowired
  lateinit var orderDTOToOrder: Converter<OrderDTO, Order>
  @Autowired
  lateinit var orderItemDTOToOrderItem: Converter<OrderItemDTO, OrderItem>
  @Autowired
  lateinit var shippingAddressDTOToShippingAddress: Converter<ShippingAddressDTO, ShippingAddress>
  @Autowired
  lateinit var paymentResultDTOToPaymentResult: Converter<PaymentResultDTO, PaymentResult>

  @Bean
  fun modelMapper(): ModelMapper {
    val modelMapper = ModelMapper()
    modelMapper.addConverter(orderDTOToOrder)
    modelMapper.addConverter(orderItemDTOToOrderItem)
    modelMapper.addConverter(shippingAddressDTOToShippingAddress)
    modelMapper.addConverter(paymentResultDTOToPaymentResult)
    return modelMapper
  }
}
