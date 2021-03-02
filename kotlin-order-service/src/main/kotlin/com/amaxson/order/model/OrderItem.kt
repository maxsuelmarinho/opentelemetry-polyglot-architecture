package com.amaxson.order.model

import com.fasterxml.jackson.annotation.JsonIgnore
import com.fasterxml.jackson.annotation.JsonProperty
import java.math.BigDecimal
import java.time.LocalDateTime
import javax.persistence.*
import javax.validation.constraints.NotNull

// TODO: create order item response. Remove Json mapping
@Entity
@Table(name = "order_items")
class OrderItem(
  @JsonIgnore @Id @GeneratedValue(strategy = GenerationType.IDENTITY) @Column(name="id") var id: Long? = null,
  @JsonIgnore @ManyToOne @JoinColumn(name="order_id") @NotNull var order: Order? = null,
  @NotNull @Column(name="name") val name: String,
  @JsonProperty("_id") @NotNull @Column(name="uuid") val uuid: String,
  @NotNull @Column(name="qty") val qty: Int,
  @NotNull @Column(name="image") val image: String,
  @NotNull @Column(name = "price") val price: BigDecimal,
  @JsonProperty("product") @NotNull @Column(name="product_id") val productId: String
): Audit()
