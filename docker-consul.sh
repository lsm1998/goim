#!/usr/bin/env bash
docker pull consul:1.3.0 && docker run -d --name=dev-consul -p 8500:8500 consul:1.3.0
