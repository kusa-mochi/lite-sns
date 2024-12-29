import { useState } from "react"
import Button from "../components/molecules/button"
import { useConfig } from "../providers/configProvider"
import { css } from "@emotion/css"

enum InputError {
    None = 0,
    InvalidCharacter,
    Empty,
    InvalidFormat,
    TooShort,
}

export default function Signup() {
    const config = useConfig()

    const [nickname, setNickname] = useState("")
    const [emailAddress, setEmailAddress] = useState("")
    const [password, setPassword] = useState("")
    const [passwordConfirm, setPasswordConfirm] = useState("")

    function encodeHTMLForm(data: any) {
        var params = []
        for (const name in data) {
            const val = data[name]
            const param = encodeURIComponent(name) + "=" + encodeURIComponent(val)
            params.push(param)
        }

        return params.join("&").replace(/%20/g, "+")
    }

    function sendEmail() {
        // 入力値チェック
        validateInputs()

        // 入力値をサーバに送信。
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
        xmlHttpReq.open("POST", `http://${config.appServer.ip}:${config.appServer.port}${config.appServer.apiPrefix}/signup`)
        xmlHttpReq.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
        xmlHttpReq.send(encodeHTMLForm({
            EmailAddr: "whoatemyapplepie@gmail.com",
            TestParam: "hogeohge",
        }))
    }

    function validateNickname() {
        console.log("validating nickname...")
    }

    function validateEmailAddress() {
        console.log("validating email address...")
    }

    function validatePassword() {
        console.log("validating password...")
    }

    function validatePasswordConfirm() {
        console.log("validating password confirmation...")
    }

    function validateInputs() {
        validateNickname()
        validateEmailAddress()
        validatePassword()
        validatePasswordConfirm()
    }

    const formStyle = css`
        display: grid;
        grid-template-columns: 216px 260px;
    `
    const labelStyle = css`
        text-align: left;
    `
    const requireStyle = css`
        color: red;
    `
    const inputStyle = css`
        width: 100%;
    `

    return (
        <>
            <div className={formStyle}>
                <div className={labelStyle}>
                    <span className={requireStyle}>*</span>ニックネーム
                </div>
                <div>
                    <input
                        className={inputStyle}
                        type="text"
                        value={nickname}
                        onChange={(e) => setNickname(e.target.value)}
                        placeholder="最大16文字"
                        maxLength={16}
                        onBlur={() => validateNickname()}
                    />
                </div>
                <div className={labelStyle}>
                    <span className={requireStyle}>*</span>メールアドレス
                </div>
                <div>
                    <input
                        className={inputStyle}
                        type="email"
                        value={emailAddress}
                        onChange={(e) => setEmailAddress(e.target.value)}
                        maxLength={254}
                        onBlur={() => validateEmailAddress()}
                    />
                </div>
                <div className={labelStyle}>
                    <span className={requireStyle}>*</span>登録するパスワード
                </div>
                <div>
                    <input
                        className={inputStyle}
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        maxLength={128}
                        onBlur={() => validatePassword()}
                    />
                </div>
                <div className={labelStyle}>
                    <span className={requireStyle}>*</span>登録するパスワード（確認）
                </div>
                <div>
                    <input
                        className={inputStyle}
                        type="password"
                        value={passwordConfirm}
                        onChange={(e) => setPasswordConfirm(e.target.value)}
                        maxLength={128}
                        onBlur={() => validatePasswordConfirm()}
                    />
                </div>
            </div>
            <Button onClick={() => sendEmail()}>認証メールを送信する</Button>
        </>
    )
}
