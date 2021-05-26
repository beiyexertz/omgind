#!/usr/bin/env bash

BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

echo "base ${BASE_DIR}"


function get_docker_compose_cmd_line() {
  ignore_db="$1"
  cmd="docker-compose -f ./compose/docker-compose-app.yml "
  use_ipv6=$(get_config USE_IPV6)
  if [[ "${use_ipv6}" != "1" ]]; then
    cmd="${cmd} -f compose/docker-compose-network.yml "
  else
    cmd="${cmd} -f compose/docker-compose-network_ipv6.yml "
  fi
  services=$(get_docker_compose_services "$ignore_db")
  if [[ "${services}" =~ celery ]]; then
    cmd="${cmd} -f ./compose/docker-compose-task.yml"
  fi
  if [[ "${services}" =~ mysql ]]; then
    cmd="${cmd} -f ./compose/docker-compose-mysql.yml"
  fi
  if [[ "${services}" =~ redis ]]; then
    cmd="${cmd} -f ./compose/docker-compose-redis.yml"
  fi
  if [[ "${services}" =~ lb ]]; then
    cmd="${cmd} -f ./compose/docker-compose-internal.yml -f ./compose/docker-compose-lb.yml"
  else
    cmd="${cmd} -f ./compose/docker-compose-external.yml"
  fi
  if [[ "${services}" =~ xpack ]]; then
    cmd="${cmd} -f ./compose/docker-compose-xpack.yml"
  fi
  if [[ "${services}" =~ omnidb ]]; then
    cmd="${cmd} -f ./compose/docker-compose-omnidb.yml"
  fi
  echo "${cmd}"
}