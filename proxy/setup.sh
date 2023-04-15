SCRIPT_DIR=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)

python3 -m venv venv
. venv/bin/activate
pip3 install -r requirements.txt
