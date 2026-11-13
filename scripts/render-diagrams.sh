#!/bin/sh
set -eu
: "${PLANTUML_JAR:?Set PLANTUML_JAR to the PlantUML 1.2025.2 jar path}"
java -Djava.awt.headless=true -jar "$PLANTUML_JAR" -charset UTF-8 -tpng docs/diagrams/*.puml
