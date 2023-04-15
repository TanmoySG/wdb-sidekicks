import os
from flask import Flask, request, Response
from dotenv import load_dotenv
import requests

userAgentHeaderKey = "User-Agent"

load_dotenv()

# configurations loaded from environment variables
FORWARD_TO = os.environ.get('FORWARD_TO')
PROXY_APPLICATION = os.environ.get('PROXY_APPLICATION')
PROXY_PORT = os.environ.get('PROXY_PORT')

proxyApplicationMessage = f'forwarded by proxy {PROXY_APPLICATION}'

app = Flask(__name__)


@app.route('/',
           defaults={'path': ''},
           methods=["POST", "GET", "DELETE", "PATCH"])
@app.route('/<path:path>', methods=["POST", "GET", "DELETE", "PATCH"])
def redirect_to_API_HOST(path):
    requestHeaders = {}

    for headerKey, headerValue in request.headers:
        if headerKey == userAgentHeaderKey:
            requestHeaders[
                headerKey] = f'{headerValue}; {proxyApplicationMessage}'
            continue
        requestHeaders[headerKey] = headerValue

    if userAgentHeaderKey not in requestHeaders:
        requestHeaders[userAgentHeaderKey] = proxyApplicationMessage

    res = requests.request(
        method=request.method,
        url=request.url.replace(request.host_url, f'{FORWARD_TO}/'),
        headers=requestHeaders,
        data=request.get_data(),
        allow_redirects=False,
    )

    excluded_headers = [
        'content-encoding', 'content-length', 'transfer-encoding', 'connection'
    ]
    headers = [(k, v) for k, v in res.raw.headers.items()
               if k.lower() not in excluded_headers]

    response = Response(res.content, res.status_code, headers)
    return response


if __name__ == "__main__":
    app.run(port=PROXY_PORT)
