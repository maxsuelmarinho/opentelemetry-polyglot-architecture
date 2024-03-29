#!/usr/bin/env bash

[ -z "$MIN_HEAP_SIZE" ] && MIN_HEAP_SIZE=40M
[ -z "$MAX_HEAP_SIZE" ] && MAX_HEAP_SIZE=512M
[ -z "$THREADSTACK_SIZE" ] && THREADSTACK_SIZE=228k
[ -z "$JAVA_GC_ARGS" ] && JAVA_GC_ARGS=-XX:MinHeapFreeRatio=20 -XX:MaxHeapFreeRatio=40 -XX:+UseSerialGC -XX:GCTimeRatio=4 -XX:AdaptiveSizePolicyWeight=90

JAVA_OPTS="${JAVA_OPTS} \
      -server \
      -Djava.security.egd=file:/dev/./urandom \
      -XX:CompressedClassSpaceSize=64m \
      -XX:MaxMetaspaceSize=256m"

echo $JAVA_OPTS $MIN_HEAP_SIZE $MAX_HEAP_SIZE $THREADSTACK_SIZE $JAVA_GC_ARGS $JAVA_DIAG_ARGS $JAVA_OPTS_APPEND $PROG_ARGS

java ${JAVA_OPTS} \
 -Xms${MIN_HEAP_SIZE} \
 -Xmx${MAX_HEAP_SIZE} \
 -Xss${THREADSTACK_SIZE} \
 ${JAVA_GC_ARGS} \
 -jar ${APP_HOME}/server.jar
