import { ChangeEvent, useEffect, useState } from "react"
import Button from "../components/molecules/button"
import { css } from "@emotion/css"
import { useConfig } from "../providers/configProvider"
import { encodeHTMLForm } from "../utils/api_utils"

export default function Signin() {
    const config = useConfig()
    const MAX_LENGTH_EMAILADDRESS: number = 254
    const MAX_LENGTH_PASSWORD: number = 128

    const [emailAddress, setEmailAddress] = useState("")
    const [password, setPassword] = useState("")
    const [isSigninEnabled, setIsSigninEnabled] = useState(false)

    function validateEmailAddress(): boolean {
        console.log("validating email address...")

        return (
            emailAddress.length > 0 &&
            emailAddress.length <= MAX_LENGTH_EMAILADDRESS &&
            emailAddress.match(/^\S+$/g) !== null &&
            emailAddress.match(/^[^\.].*@[a-zA-Z0-9_-]+\.[a-zA-Z0-9\._-]+$/) !== null
        )
    }

    function validatePassword(): boolean {
        console.log("validating password...")

        return (
            password.length > 0 &&
            password.length <= MAX_LENGTH_PASSWORD &&
            password.match(/^\S+$/g) !== null
        )
    }

    function validate() {
        const isValidEmailAddr: boolean = validateEmailAddress()
        const isValidPassword: boolean = validatePassword()

        setIsSigninEnabled(isValidEmailAddr && isValidPassword)
    }

    function onEmailAddrChanged(e: ChangeEvent<HTMLInputElement>) {
        setEmailAddress(e.target.value)
    }

    function onPasswordChanged(e: ChangeEvent<HTMLInputElement>) {
        setPassword(e.target.value)
    }

    function signin() {
        console.log("signing in...")
        const apiPath: string = `http://${config.appServer.ip}:${config.appServer.port}${config.appServer.apiPrefix}/public/signin`
        console.log(`api path: ${apiPath}`)
        const xmlHttpReq = new XMLHttpRequest()
        xmlHttpReq.onreadystatechange = function () {
            const READYSTATE_COMPLETED: number = 4
            const HTTP_STATUS_OK: number = 200
            if (
                this.readyState === READYSTATE_COMPLETED &&
                this.status === HTTP_STATUS_OK
            ) {
                console.log("sign in succeeded")
                
                const res = JSON.parse(this.response)
                console.log(res.token)  // APIサーバーが発行したアクセストークン

                // TODO: アクセストークンをコンテキストに保存し、任意のページで使えるようにする。

                location.replace("/timeline")
            }
        }
        xmlHttpReq.open("POST", apiPath)
        xmlHttpReq.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
        xmlHttpReq.send(encodeHTMLForm({
            EmailAddr: emailAddress,
            Password: password,
        }))
    }

    useEffect(() => {
        validate()
    }, [emailAddress, password])

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
                    <span className={requireStyle}>*</span>メールアドレス
                </div>
                <div>
                    <input
                        className={inputStyle}
                        type="email"
                        value={emailAddress}
                        onChange={onEmailAddrChanged}
                        placeholder={"例: example@slash-mochi.net"}
                        maxLength={MAX_LENGTH_EMAILADDRESS}
                    />
                </div>
                <div className={labelStyle}>
                    <span className={requireStyle}>*</span>パスワード
                </div>
                <div>
                    <input
                        className={inputStyle}
                        type="password"
                        value={password}
                        onChange={onPasswordChanged}
                        maxLength={MAX_LENGTH_PASSWORD}
                    />
                </div>
            </div>
            <Button disabled={!isSigninEnabled} onClick={() => signin()}>サインイン</Button>
            <p>Go to <a href="/test2">Test2 Page</a></p>
        </>
    )
}
