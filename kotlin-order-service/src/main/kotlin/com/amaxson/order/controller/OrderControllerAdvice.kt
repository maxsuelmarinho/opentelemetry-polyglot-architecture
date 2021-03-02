package com.amaxson.order.controller

import com.amaxson.order.exception.OrderNotFoundException
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.ExceptionHandler
import org.springframework.web.bind.annotation.RestControllerAdvice
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler

@RestControllerAdvice
class OrderControllerAdvice: ResponseEntityExceptionHandler() {

  @ExceptionHandler(OrderNotFoundException::class)
  fun handleUserNotFound(e: OrderNotFoundException) = ResponseEntity.notFound().build<Any>()
}
