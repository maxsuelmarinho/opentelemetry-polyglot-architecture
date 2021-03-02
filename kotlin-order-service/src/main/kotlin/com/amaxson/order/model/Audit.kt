package com.amaxson.order.model

import org.springframework.data.annotation.CreatedDate
import org.springframework.data.annotation.LastModifiedDate
import org.springframework.data.jpa.domain.support.AuditingEntityListener
import java.time.LocalDateTime
import java.util.*
import javax.persistence.*

@MappedSuperclass
@EntityListeners(AuditingEntityListener::class)
abstract class Audit(
  @Column(name="created_at", nullable = false, updatable = false) @CreatedDate var createdAt: LocalDateTime? = null,
  @Column(name="updated_at", nullable = false) @LastModifiedDate var updatedAt: LocalDateTime? = null,
  //@Column(name="deleted_at", nullable = true, updatable = false) var deletedAt: LocalDateTime? = null
)
