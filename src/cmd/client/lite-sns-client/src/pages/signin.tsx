import { ChangeEvent, useEffect, useState } from "react"
import Button from "../components/molecules/button"
import { css } from "@emotion/css"
import { useConfig } from "../providers/configProvider"
import { callAPI } from "../utils/api_utils"
import { Auth, setAuthType, useSetAuth } from "../providers/authProvider"
import { useNavigate } from "react-router"
import Card from "../components/atoms/card"

export default function Signin() {
    const config = useConfig()
    const MAX_LENGTH_EMAILADDRESS: number = 254
    const MAX_LENGTH_PASSWORD: number = 128

    const navigate = useNavigate()

    const [emailAddress, setEmailAddress] = useState("")
    const [password, setPassword] = useState("")
    const [isSigninEnabled, setIsSigninEnabled] = useState(false)

    // アクセストークンとユーザーID（固有の数値）をコンテキストに保存し、任意のページで使えるようにするためのフック。
    const setAuth = useSetAuth()

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
        callAPI(
            apiPath,
            "POST",
            {
                EmailAddr: emailAddress,
                Password: password,
            },
            -1,
            null,
            (response: any) => {
                console.log(`token:   ${response.token}`)      // APIサーバーが発行したアクセストークン
                console.log(`user id: ${response.user_id}`)    // ユーザーID（固有の数値）

                const au: Auth = {
                    userId: response.user_id,
                    tokenString: response.token,
                }
                setAuth({ type: setAuthType, payload: au })

                navigate("/timeline")
            }
        )
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
        <Card topBorder>
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
        </Card>
    )
}
