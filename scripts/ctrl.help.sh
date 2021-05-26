#!/bin/bash

echo "source help file"

function usage() {
  echo "$(gettext 'omgind Deployment Management Script')"
  echo
  echo "Usage: "
  echo "  ./ctrl.{mode}.sh [COMMAND] [ARGS...]"
  echo "  ./ctrl.{mode}.sh --help"
  echo

}
