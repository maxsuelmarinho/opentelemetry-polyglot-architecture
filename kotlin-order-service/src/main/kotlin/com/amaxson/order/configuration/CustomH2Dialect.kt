package com.amaxson.order.configuration

import org.hibernate.dialect.H2Dialect
import org.slf4j.Logger
import org.slf4j.LoggerFactory
import java.sql.Types

class CustomH2Dialect: H2Dialect() {
  companion object {
    val logger: Logger = LoggerFactory.getLogger(CustomH2Dialect::class.java)
  }
  init {
    logger.info("Registering 'jsonb' type")
    registerColumnType(Types.OTHER, "jsonb")
  }
}
