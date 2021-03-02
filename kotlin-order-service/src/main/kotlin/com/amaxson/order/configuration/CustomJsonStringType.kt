package com.amaxson.order.configuration

import com.vladmihalcea.hibernate.type.json.JsonStringType

class CustomJsonStringType : JsonStringType(CustomObjectMapperWrapper()) {

  override fun getName(): String {
    return "jsonb"
  }
}
