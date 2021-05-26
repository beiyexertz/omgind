#!/usr/bin/env bash

CURRENT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

PROJECT_DIR=$(dirname "${CURRENT_DIR}")

echo "current project dir: ${PROJECT_DIR}"

# 倒入文件
. "${PROJECT_DIR}/scripts/ctrl.help.sh"
. "${PROJECT_DIR}/scripts/utils.sh"


action=${1-}
target=${2-}
args=("$@")

#echo "${args}"
#echo "${action}"
#echo "${target}"

EXECUTOR=""


function pre_check() {
  check_config_file || return 3
}

function main() {
    if [[ "${action}" == "help" || "${action}" == "h" || "${action}" == "-h" || "${action}" == "--help" ]]; then
      echo ""
    elif [[ "${action}" == "" ]]; then
      echo ""
    else
      pre_check || return 3
      EXECUTOR=$(get_docker_compose_cmd_line)
    fi
  
  case "${action}" in
  start)
    echo ""

    ;;
  stop)
    echo ""

    ;;
  close)
    echo ""

    ;;
  restart)
    echo ""

    ;;
  status)
    echo ""

    ;;
  down)
    echo ""

    ;;
  uninstall)
    echo ""

    ;;

  help)
    usage
    ;;
  --help)
    usage
    ;;
  h)
    usage
    ;;
  -h)
    usage
    ;;
  *)
    echo "No such command: ${action}"
    usage
    ;;
  esac
}

main "$@"