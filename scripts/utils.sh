#!/usr/bin/env bash

BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

echo "sourcing: ${BASH_SOURCE[0]}"
echo "mode: ${CMODE}"

. "${BASE_DIR}/const.${CMODE}.sh"



function get_config() {
  cwd=$(pwd)
  key=$1
  value=$(grep "^${key}=" "${CONFIG_FILE}" | awk -F= '{ print $2 }')
  echo "${value}"
}

function get_docker_compose_services() {
  ignore_db="$1"
  services="core"
  services+=" nginx"

  use_external_mysql=$(get_config USE_EXTERNAL_MYSQL)
  if [[ "${use_external_mysql}" != "1" && "${ignore_db}" != "ignore_db" ]]; then
    services+=" mysql"
  fi
  use_external_redis=$(get_config USE_EXTERNAL_REDIS)
  if [[ "${use_external_redis}" != "1" && "${ignore_db}" != "ignore_db" ]]; then
    services+=" redis"
  fi
  use_lb=$(get_config USE_LB)
  if [[ "${use_lb}" == "1" ]]; then
    services+=" lb"
  fi
  services+=" rabbitmq influxdb"

  echo "${services}"
}

function get_docker_compose_cmd_line() {
  ignore_db="$1"
  compose_dir="./scripts/compose.${CMODE}"

  cmd="docker compose -f ${compose_dir}/docker-compose.app.yml "
  use_ipv6=$(get_config USE_IPV6)
  if [[ "${use_ipv6}" != "1" ]]; then
    cmd="${cmd} -f ${compose_dir}/docker-compose.network.ipv4.yml "
  else
    cmd="${cmd} -f ${compose_dir}/docker-compose.network.ipv6.yml "
  fi

  services=$(get_docker_compose_services "$ignore_db")

  if [[ "${services}" =~ mysql ]]; then
    cmd="${cmd} -f ${compose_dir}/docker-compose.mysql.yml"
  fi
  if [[ "${services}" =~ redis ]]; then
    cmd="${cmd} -f ${compose_dir}/docker-compose.redis.yml"
  fi
  if [[ "${services}" =~ lb ]]; then
    cmd="${cmd} -f ${compose_dir}/docker-compose.internal.yml -f ./compose/docker-compose-lb.yml"
  else
#    cmd="${cmd} -f ${compose_dir}/docker-compose.external.yml"
    echo ""
  fi

  echo "${cmd}"
}