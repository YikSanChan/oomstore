#!/usr/bin/env bash
SDIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd) && cd "$SDIR" || exit 1
source ./util.sh

init_store
register_features
import_sample

case='featctl update feature works'
featctl update feature price --description "new description"
expected='
Name:          price
Group:         phone
Entity:        device
Category:      batch
DBValueType:   int
ValueType:     int32
Description:   new description
Revision:
DataTable:
CreateTime:
ModifyTime:
'
actual=$(featctl describe feature price)
ignore() { grep -Ev '^(CreateTime|ModifyTime|Revision|DataTable)' <<<"$1"; }
assert_eq "$case" "$(ignore "$expected")" "$(ignore "$actual")"
