// Used with Postgres
@TypeDef(name = "jsonb", typeClass = JsonBinaryType.class)
package com.amaxson.order.configuration;
import com.vladmihalcea.hibernate.type.json.JsonBinaryType;
import org.hibernate.annotations.TypeDef;

// Used with H2
//@TypeDef(name = "jsonb", typeClass = CustomJsonStringType.class)
//package com.amaxson.order.configuration;
//import org.hibernate.annotations.TypeDef;
