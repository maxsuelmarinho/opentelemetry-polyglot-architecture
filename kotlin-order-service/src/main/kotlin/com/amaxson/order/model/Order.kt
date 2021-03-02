package com.amaxson.order.model

import com.fasterxml.jackson.annotation.JsonIgnore
import com.fasterxml.jackson.annotation.JsonProperty
import org.hibernate.annotations.CreationTimestamp
import org.hibernate.annotations.Type
import org.hibernate.annotations.UpdateTimestamp
import org.springframework.data.annotation.CreatedDate
import org.springframework.data.annotation.LastModifiedDate
import java.math.BigDecimal
import java.time.LocalDateTime
import java.util.*
import javax.persistence.*
import javax.validation.constraints.NotNull

// TODO: create order response. Remove Json mapping
@Entity
@Table(name = "orders")
class Order(
  @Id
  @JsonIgnore @GeneratedValue(strategy = GenerationType.IDENTITY) @Column(name="id") var id: Long? = null,
  @JsonProperty("_id") @field:NotNull @Column(name="uuid") val uuid: String,
  @OneToMany(mappedBy = "order", cascade = [CascadeType.ALL], fetch = FetchType.EAGER) var orderItems: List<OrderItem> = emptyList(),
  @JsonProperty("user") @field:NotNull @Column(name="user_id") val userId: String,
  @field:NotNull @Column(name="payment_method") val paymentMethod: String,
  @field:NotNull @Column(name = "tax_price") val taxPrice: BigDecimal,
  @field:NotNull @Column(name = "shipping_price") val shippingPrice: BigDecimal,
  @field:NotNull @Column(name = "total_price") val totalPrice: BigDecimal,
  @field:NotNull @Column(name = "is_paid") var isPaid: Boolean,
  @Column(name = "paid_at", columnDefinition = "TIMESTAMP WITH TIME ZONE") var paidAt: LocalDateTime? = null,
  @field:NotNull @Column(name = "is_delivered") val isDelivered: Boolean,
  @Column(name = "delivered_at", columnDefinition = "TIMESTAMP WITH TIME ZONE") val deliveredAt: LocalDateTime? = null,
  @Type(type = "jsonb") @field:NotNull @Column(name="shipping_address") val shippingAddress: ShippingAddress,
  @Type(type = "jsonb") @Column(name="payment_result") var paymentResult: PaymentResult? = null
): Audit()
