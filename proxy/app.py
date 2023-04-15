from dotenv import load_dotenv  # pip package python-dotenv
import os
#
from flask import Flask, request, Response
import requests  # pip package requests

userAgentHeaderKey = "User-Agent"
proxyApplication = "wdb-proxy"

proxyApplicationMessage = f'forwarded by proxy {proxyApplication}'

load_dotenv()
API_HOST = os.environ.get('FORWARDED_TO')
assert API_HOST, 'Envvar API_HOST is required'

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
        url=request.url.replace(request.host_url, f'{API_HOST}/'),
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


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
