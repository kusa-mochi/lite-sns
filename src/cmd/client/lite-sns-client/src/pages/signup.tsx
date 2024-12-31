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
    TooLong,
    ConfirmationNotMatched,
}

export default function Signup() {
    const config = useConfig()
    const MAX_LENGTH_USERNAME: number = 20
    const MAX_LENGTH_EMAILADDRESS: number = 254
    const MIN_LENGTH_PASSWORD: number = 12
    const MAX_LENGTH_PASSWORD: number = 128

    const [nickname, setNickname] = useState("")
    const [isNicknameInvalid, setIsNicknameInvalid] = useState(false)

    const [emailAddress, setEmailAddress] = useState("")
    const [isEmailAddressInvalid, setIsEmailAddressInvalid] = useState(false)

    const [password, setPassword] = useState("")
    const [isPasswordInvalid, setIsPasswordInvalid] = useState(false)

    const [passwordConfirm, setPasswordConfirm] = useState("")
    const [isPasswordConfirmInvalid, setIsPasswordConfirmInvalid] = useState(false)

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
        const hasInputError: boolean = !validateInputs()
        if (hasInputError) return

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
            EmailAddr: emailAddress,
            Nickname: nickname,
            Password: password,
        }))
    }

    function validateNickname(): InputError {
        console.log("validating nickname...")

        if (!nickname) {
            setIsNicknameInvalid(true)
            return InputError.Empty
        }
        if (nickname.length > MAX_LENGTH_USERNAME) {
            setIsNicknameInvalid(true)
            return InputError.TooLong
        }
        if (!nickname.match(/\S/g)) {
            setIsNicknameInvalid(true)
            return InputError.Empty
        }

        setIsNicknameInvalid(false)
        return InputError.None
    }

    function validateEmailAddress(): InputError {
        console.log("validating email address...")

        if (!emailAddress) {
            setIsEmailAddressInvalid(true)
            return InputError.Empty
        }
        if (emailAddress.length > MAX_LENGTH_EMAILADDRESS) {
            setIsEmailAddressInvalid(true)
            return InputError.TooLong
        }
        if (!emailAddress.match(/\S/g)) {
            setIsEmailAddressInvalid(true)
            return InputError.Empty
        }
        if (!emailAddress.match(/^[^\.].*@[a-zA-Z0-9_-]+\.[a-zA-Z0-9\._-]+[^\.]$/)) {
            setIsEmailAddressInvalid(true)
            return InputError.InvalidFormat
        }

        setIsEmailAddressInvalid(false)
        return InputError.None
    }

    function validatePassword(): InputError {
        console.log("validating password...")

        if (!password) {
            setIsPasswordInvalid(true)
            return InputError.Empty
        }
        if (password.length < MIN_LENGTH_PASSWORD) {
            setIsPasswordInvalid(true)
            return InputError.TooShort
        }
        if (password.length > MAX_LENGTH_PASSWORD) {
            setIsPasswordInvalid(true)
            return InputError.TooLong
        }
        if (!password.match(/\S/g)) {
            setIsPasswordInvalid(true)
            return InputError.Empty
        }
        // if use only ASCII characters,
        if (!/^[\x00-\x7F]+$/.test(password)) {
            setIsPasswordInvalid(true)
            return InputError.InvalidCharacter
        }

        setIsPasswordInvalid(false)
        return InputError.None
    }

    function validatePasswordConfirm(): InputError {
        console.log("validating password confirmation...")

        if (!passwordConfirm) {
            setIsPasswordConfirmInvalid(true)
            return InputError.Empty
        }
        if (password !== passwordConfirm) {
            setIsPasswordConfirmInvalid(true)
            return InputError.ConfirmationNotMatched
        }

        setIsPasswordConfirmInvalid(false)
        return InputError.None
    }

    // 戻り値
    //   true:  正常終了
    //   false: 入力値エラー有り
    function validateInputs(): boolean {
        const nicknameResult: InputError = validateNickname()
        const emailAddressResult: InputError = validateEmailAddress()
        const passwordResult: InputError = validatePassword()
        const passwordConfirmResult: InputError = validatePasswordConfirm()

        return (
            nicknameResult === InputError.None &&
            emailAddressResult === InputError.None &&
            passwordResult === InputError.None &&
            passwordConfirmResult === InputError.None
        )
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
    const invalidInputStyle = css`
        ${inputStyle}
        background-color: #ffd6d6;
    `

    return (
        <>
            <div className={formStyle}>
                <div className={labelStyle}>
                    <span className={requireStyle}>*</span>ニックネーム
                </div>
                <div>
                    <input
                        className={isNicknameInvalid ? invalidInputStyle : inputStyle}
                        type="text"
                        value={nickname}
                        onChange={(e) => setNickname(e.target.value)}
                        placeholder={`最大${MAX_LENGTH_USERNAME}文字`}
                        maxLength={MAX_LENGTH_USERNAME}
                        onBlur={() => validateNickname()}
                    />
                </div>
                <div className={labelStyle}>
                    <span className={requireStyle}>*</span>メールアドレス
                </div>
                <div>
                    <input
                        className={isEmailAddressInvalid ? invalidInputStyle : inputStyle}
                        type="email"
                        value={emailAddress}
                        onChange={(e) => setEmailAddress(e.target.value)}
                        placeholder={"例: example@slash-mochi.net"}
                        maxLength={MAX_LENGTH_EMAILADDRESS}
                        onBlur={() => validateEmailAddress()}
                    />
                </div>
                <div className={labelStyle}>
                    <span className={requireStyle}>*</span>登録するパスワード
                </div>
                <div>
                    <input
                        className={isPasswordInvalid ? invalidInputStyle : inputStyle}
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        placeholder={`${MIN_LENGTH_PASSWORD}文字以上`}
                        maxLength={MAX_LENGTH_PASSWORD}
                        onBlur={() => validatePassword()}
                    />
                </div>
                <div className={labelStyle}>
                    <span className={requireStyle}>*</span>登録するパスワード（確認）
                </div>
                <div>
                    <input
                        className={isPasswordConfirmInvalid ? invalidInputStyle : inputStyle}
                        type="password"
                        value={passwordConfirm}
                        onChange={(e) => setPasswordConfirm(e.target.value)}
                        maxLength={MAX_LENGTH_PASSWORD}
                        onBlur={() => validatePasswordConfirm()}
                    />
                </div>
            </div>
            <Button onClick={() => sendEmail()}>認証メールを送信する</Button>
        </>
    )
}
