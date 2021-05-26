#!/bin/bash

action=${1-}
target=${2-}
args=("$@")

echo "${args}"
echo "${action}"
echo "${target}"

EXECUTOR=""

# 倒入文件
. ./scripts/ctrl.help.sh

function pre_check() {
  check_config_file || return 3
}

function main() {
    echo " ------- main "
    if [[ "${action}" == "help" || "${action}" == "h" || "${action}" == "-h" || "${action}" == "--help" ]]; then
      echo ""
    elif [[ "${action}" == "" ]]; then
      echo ""
    else
      pre_check || return 3
      EXECUTOR=$(get_docker_compose_cmd_line)
    fi
  
  case "${action}" in


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