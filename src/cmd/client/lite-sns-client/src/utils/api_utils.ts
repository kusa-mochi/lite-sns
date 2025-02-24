function encodeHTMLForm(data: any) {
  var params = [];
  for (const name in data) {
    const val = data[name];
    const param = encodeURIComponent(name) + "=" + encodeURIComponent(val);
    params.push(param);
  }

  return params.join("&").replace(/%20/g, "+");
}

export function callAPI(
  apiPath: string,
  method: string,
  data: any,
  userId: number,
  accessToken: string | null,
  onSuccess: (response: any) => void,
  onFailure?: (response: any) => void,
  onRedirect?: (response: any) => void
) {
  const xmlHttpReq = new XMLHttpRequest();
  xmlHttpReq.onreadystatechange = function () {
    const READYSTATE_COMPLETED: number = 4;
    const HTTP_STATUS_OK: number = 200;
    const HTTP_SEE_OTHER: number = 303;
    console.log(`ready state: ${this.readyState}`);
    console.log(`status: ${this.status}`);
    const res = (() => {
      try {
        const r = JSON.parse(this.response)
        return r
      } catch {
        // this.responseがJSON形式に変換できない場合、resをnullで初期化する。
        return null
      }
    })()

    if (this.status === HTTP_STATUS_OK || this.status === 0) {
      if (this.readyState === READYSTATE_COMPLETED) {
        onSuccess(res);
      }
    } else if (this.status === HTTP_SEE_OTHER) {
      alert("see other")
      if (onRedirect) {
        console.log("redirecting...")
        onRedirect(res);
      }
    } else {
      if (onFailure) {
        onFailure(res);
      }
    }
  };
  xmlHttpReq.open(method, apiPath);
  xmlHttpReq.setRequestHeader(
    "Content-Type",
    "application/x-www-form-urlencoded"
  );
  xmlHttpReq.setRequestHeader("X-User-Id", userId.toString());
  xmlHttpReq.setRequestHeader("Authorization", accessToken ?? "");
  xmlHttpReq.send(encodeHTMLForm(data));
}
