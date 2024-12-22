import Button from "../components/molecules/button"

export default function Signup() {

    function EncodeHTMLForm(data: any) {
        var params = []
        for (const name in data) {
            const val = data[name]
            const param = encodeURIComponent(name) + "=" + encodeURIComponent(val)
            params.push(param)
        }

        return params.join("&").replace(/%20/g, "+")
    }

    function SendEmail() {
        console.log("sending an email...")
        const xmlHttpReq = new XMLHttpRequest()
        xmlHttpReq.onreadystatechange = function () {
            const READYSTATE_COMPLETED: number = 4
            const HTTP_STATUS_OK: number = 200
            if (
                this.readyState === READYSTATE_COMPLETED &&
                this.status === HTTP_STATUS_OK
            ) {
                console.log("sending email succeeded")
            }
        }
        xmlHttpReq.open("POST", "http://localhost:12381/lite-sns/api/v1/signup")
        xmlHttpReq.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
        xmlHttpReq.send(EncodeHTMLForm({
            EmailAddr: "whoatemyapplepie@gmail.com",
            TestParam: "hogeohge",
        }))
    }

    return (
        <>
            <Button onClick={() => SendEmail()}>Send a email</Button>
        </>
    )
}
