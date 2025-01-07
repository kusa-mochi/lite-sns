function encodeHTMLForm(data: any) {
    var params = []
    for (const name in data) {
        const val = data[name]
        const param = encodeURIComponent(name) + "=" + encodeURIComponent(val)
        params.push(param)
    }

    return params.join("&").replace(/%20/g, "+")
}

export function callAPI(
    apiPath: string,
    method: string,
    data: any,
    userId: number,
    accessToken: string | null,
    onSuccess: (response: any) => void,
    onFailure?: (response: any) => void,
) {
    const xmlHttpReq = new XMLHttpRequest()
    xmlHttpReq.onreadystatechange = function () {
        const READYSTATE_COMPLETED: number = 4
        const HTTP_STATUS_OK: number = 200
        if (
            this.readyState === READYSTATE_COMPLETED &&
            this.status === HTTP_STATUS_OK
        ) {
            onSuccess(JSON.parse(this.response))
        } else {
            if (onFailure) {
                onFailure(JSON.parse(this.response))
            }
        }
    }
    xmlHttpReq.open(method, apiPath)
    xmlHttpReq.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    xmlHttpReq.setRequestHeader("X-User-Id", userId.toString())
    xmlHttpReq.setRequestHeader("Authorization", accessToken ?? "")
    xmlHttpReq.send(encodeHTMLForm(data))
}
