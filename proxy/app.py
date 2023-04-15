import calendar
import logging
import os
import time
from flask import Flask, request, Response
from dotenv import load_dotenv
import requests

logging.basicConfig(filename="record.log", level=logging.INFO)
logging.getLogger("werkzeug").setLevel("WARNING")  # to filter out werkzeug logs

# configurations loaded from environment variables
load_dotenv()
FORWARD_TO = os.environ.get("FORWARD_TO_HOST")
PROXY_APPLICATION = os.environ.get("PROXY_APPLICATION")
PROXY_PORT = os.environ.get("PROXY_PORT")
REMOVE_HOST_PATH = os.environ.get("REMOVE_HOST_PATH")

user_agent_header_key = "User-Agent"
proxy_application_user_agent = f"forwarded by proxy {PROXY_APPLICATION}"

app = Flask(__name__)


@app.route("/", defaults={"path": ""}, methods=["POST", "GET", "DELETE", "PATCH"])
@app.route("/<path:path>", methods=["POST", "GET", "DELETE", "PATCH"])
def redirect_to_API_HOST(path):
    updated_request_headers = {}

    for header_key, header_value in request.headers:
        if header_key == user_agent_header_key:
            updated_request_headers[
                header_key
            ] = f"{header_value}; {proxy_application_user_agent}"
            continue
        updated_request_headers[header_key] = header_value

    if user_agent_header_key not in updated_request_headers:
        updated_request_headers[user_agent_header_key] = proxy_application_user_agent

    host_url_to_remove = f"{request.host_url}{REMOVE_HOST_PATH}"
    forward_url = request.url.replace(host_url_to_remove, f"{FORWARD_TO}/")

    res = requests.request(
        method=request.method,
        url=forward_url,
        headers=updated_request_headers,
        data=request.get_data(),
        allow_redirects=False,
    )

    excluded_headers = [
        "content-encoding",
        "content-length",
        "transfer-encoding",
        "connection",
    ]
    headers = [
        (k, v) for k, v in res.raw.headers.items() if k.lower() not in excluded_headers
    ]

    log_event(method=request.method, url=request.url, forwarded_to=forward_url)

    return Response(res.content, res.status_code, headers)


def log_event(method, url, forwarded_to):
    current_GMT = time.gmtime()
    time_stamp = calendar.timegm(current_GMT)
    app.logger.info(
        {
            "method": method,
            "request": url,
            "forwarded_to": forwarded_to,
            "timestamp": time_stamp,
        }
    )


if __name__ == "__main__":
    app.run(port=PROXY_PORT)
