#!/bin/sh
docker-compose up --build -d
docker exec -it gotoys_go_1 bash
