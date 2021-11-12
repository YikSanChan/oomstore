#!/usr/bin/env bash
SDIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd) && cd "$SDIR" || exit 1
source ./util.sh

init_store
register_features
import_sample > /dev/null

case='featctl list feature works'
expected='Name,Group,Entity,Category,DBValueType,ValueType,Description,OnlineRevision
model,phone,device,batch,varchar(32),string,,<NULL>
price,phone,device,batch,int,int32,,<NULL>
'
actual=$(featctl list feature -o csv)
ignore_time() { cut -d ',' -f 1-8 <<<"$1"; }
assert_eq "$case" "$expected" "$(ignore_time "$actual" | sort)"
