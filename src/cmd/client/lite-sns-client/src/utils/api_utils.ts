export function encodeHTMLForm(data: any) {
    var params = []
    for (const name in data) {
        const val = data[name]
        const param = encodeURIComponent(name) + "=" + encodeURIComponent(val)
        params.push(param)
    }

    return params.join("&").replace(/%20/g, "+")
}
