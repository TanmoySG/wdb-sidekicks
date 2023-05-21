# Script to build all tools (built with Go)
# Usage: sh ./tools/build.sh
parent_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)

for d in $parent_path/*/; do
    if [ -e $d/go.mod ]; then
        tool_name=$(basename $d)
        go build -o tools/bin/${tool_name} $d/${tool_name}.go
    fi
done
