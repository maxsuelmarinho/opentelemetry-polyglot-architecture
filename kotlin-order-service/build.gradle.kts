import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
	id("org.springframework.boot") version "2.4.3"
	id("io.spring.dependency-management") version "1.0.11.RELEASE"
	kotlin("jvm") version "1.4.30"
	kotlin("plugin.spring") version "1.4.30"
  kotlin("plugin.jpa") version "1.4.30"
}

group = "com.amaxson"
version = "0.1.0-SNAPSHOT"
java.sourceCompatibility = JavaVersion.VERSION_1_8

repositories {
	mavenCentral()
}

dependencies {
	implementation("org.springframework.boot:spring-boot-starter")
  implementation("org.springframework.boot:spring-boot-starter-web")
  implementation("org.springframework.boot:spring-boot-starter-actuator")
  implementation("org.springframework.boot:spring-boot-starter-data-jpa")
  implementation("org.springframework.boot:spring-boot-starter-validation")
  implementation("org.springframework.boot:spring-boot-starter-hateoas")
  runtimeOnly("org.springframework.boot:spring-boot-devtools")
  implementation("org.modelmapper:modelmapper:2.3.3")
  implementation("org.flywaydb:flyway-core:5.2.4")
  implementation("com.fasterxml.jackson.module:jackson-module-kotlin")
  implementation("com.vladmihalcea:hibernate-types-52:2.4.1")
  implementation("org.postgresql:postgresql:42.2.5")
  implementation("com.h2database:h2")
	implementation("org.jetbrains.kotlin:kotlin-reflect")
	implementation("org.jetbrains.kotlin:kotlin-stdlib-jdk8")
  implementation("io.opentelemetry:opentelemetry-api:1.7.0")
  implementation("io.opentelemetry:opentelemetry-api-metrics:1.7.0-alpha")
  runtimeOnly("io.opentelemetry.instrumentation:opentelemetry-oshi:1.7.0-alpha")
  implementation("io.opentelemetry:opentelemetry-extension-annotations:1.7.0")
  implementation("io.opentelemetry:opentelemetry-exporter-logging:1.7.0")
  runtimeOnly("com.github.oshi:oshi-core:5.3.1")
  implementation("io.grpc:grpc-bom:1.34.1")
  implementation("io.micrometer:micrometer-core")
  implementation("io.micrometer:micrometer-registry-prometheus")
	testImplementation("org.springframework.boot:spring-boot-starter-test")
}

tasks.withType<KotlinCompile> {
	kotlinOptions {
		freeCompilerArgs = listOf("-Xjsr305=strict")
		jvmTarget = "1.8"
	}
}

tasks.withType<Test> {
	useJUnitPlatform()
}
