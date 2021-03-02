package com.amaxson.order.configuration

import com.fasterxml.jackson.databind.JsonNode
import com.vladmihalcea.hibernate.type.util.ObjectMapperWrapper
import org.h2.util.StringUtils
import java.lang.reflect.Type

class CustomObjectMapperWrapper: ObjectMapperWrapper() {

  override fun <T : Any?> fromString(string: String?, clazz: Class<T>?): T {
    val value = super.fromString(string, clazz)
    return value
  }

  override fun toString(value: Any?): String {
    val string = super.toString(value)
    return StringUtils.convertBytesToHex(string.toByteArray(Charsets.UTF_8))
  }

  override fun toJsonNode(value: String?): JsonNode {
    val string = String(StringUtils.convertHexToBytes(value), Charsets.UTF_8)
    return super.toJsonNode(string)
  }

  override fun <T : Any?> fromString(value: String?, type: Type?): T {
    val string = String(StringUtils.convertHexToBytes(value), Charsets.UTF_8)
      .replace("\u0000", "")
    return super.fromString(string, type)
  }

}
