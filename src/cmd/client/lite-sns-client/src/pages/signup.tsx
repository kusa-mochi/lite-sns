import { useState } from "react";
import Button from "../components/molecules/button";
import { useConfig } from "../providers/configProvider";
import { css } from "@emotion/css";
import { callAPI } from "../utils/api_utils";
import Card from "../components/atoms/card";
import Modal from "../components/molecules/modal";
import { useTheme } from "../providers/themeProvider";

enum InputError {
  None = 0,
  InvalidCharacter,
  Empty,
  InvalidFormat,
  TooShort,
  TooLong,
  ConfirmationNotMatched,
}

enum SignupState {
  Input = 0,
  SendingEmail,
  SendedEmail,
  AlreadyRegistered,
  FailedToSendEmail,
}

export default function Signup() {
  const config = useConfig();
  const theme = useTheme();
  const MAX_LENGTH_USERNAME: number = 20;
  const MAX_LENGTH_EMAILADDRESS: number = 254;
  const MIN_LENGTH_PASSWORD: number = 12;
  const MAX_LENGTH_PASSWORD: number = 128;

  const [nickname, setNickname] = useState("");
  const [isNicknameInvalid, setIsNicknameInvalid] = useState(false);

  const [emailAddress, setEmailAddress] = useState("");
  const [isEmailAddressInvalid, setIsEmailAddressInvalid] = useState(false);

  const [password, setPassword] = useState("");
  const [isPasswordInvalid, setIsPasswordInvalid] = useState(false);

  const [passwordConfirm, setPasswordConfirm] = useState("");
  const [isPasswordConfirmInvalid, setIsPasswordConfirmInvalid] =
    useState(false);

  const [signupState, setSignupState] = useState(SignupState.Input);

  function sendEmail() {
    // 入力値チェック
    const hasInputError: boolean = !validateInputs();
    if (hasInputError) return;

    // 入力値をサーバに送信。
    console.log("sending an email...");
    const apiPath: string = `http://${config.appServer.ip}:${config.appServer.port}${config.appServer.apiPrefix}/public/signup`;

    // 「認証メール送信中」メッセージを表示する。
    setSignupState(SignupState.SendingEmail);

    callAPI(
      apiPath,
      "POST",
      {
        EmailAddr: emailAddress,
        Nickname: nickname,
        Password: password,
      },
      -1,
      null,
      (response: any) => {
        console.log("sending email succeeded");
        console.log(response.result);

        // TODO: メール送信の成功／失敗をメッセージに表示する。
        switch (response.result) {
          case "signup fin": // 成功
            setSignupState(SignupState.SendedEmail);
            break;
          case "already registered":
            setSignupState(SignupState.AlreadyRegistered);
            break;
          case "internal server error":
            setSignupState(SignupState.FailedToSendEmail);
            break;
          default:
            setSignupState(SignupState.FailedToSendEmail);
            break;
        }
      }
    );
  }

  function validateNickname(): InputError {
    console.log("validating nickname...");

    if (!nickname) {
      setIsNicknameInvalid(true);
      return InputError.Empty;
    }
    if (nickname.length > MAX_LENGTH_USERNAME) {
      setIsNicknameInvalid(true);
      return InputError.TooLong;
    }
    if (!nickname.match(/\S/g)) {
      setIsNicknameInvalid(true);
      return InputError.Empty;
    }

    setIsNicknameInvalid(false);
    return InputError.None;
  }

  function validateEmailAddress(): InputError {
    console.log("validating email address...");

    if (!emailAddress) {
      setIsEmailAddressInvalid(true);
      return InputError.Empty;
    }
    if (emailAddress.length > MAX_LENGTH_EMAILADDRESS) {
      setIsEmailAddressInvalid(true);
      return InputError.TooLong;
    }
    if (!emailAddress.match(/\S/g)) {
      setIsEmailAddressInvalid(true);
      return InputError.Empty;
    }
    if (!emailAddress.match(/^[^\.].*@[a-zA-Z0-9_-]+\.[a-zA-Z0-9\._-]+$/)) {
      setIsEmailAddressInvalid(true);
      return InputError.InvalidFormat;
    }

    setIsEmailAddressInvalid(false);
    return InputError.None;
  }

  function validatePassword(): InputError {
    console.log("validating password...");

    if (!password) {
      setIsPasswordInvalid(true);
      return InputError.Empty;
    }
    if (password.length < MIN_LENGTH_PASSWORD) {
      setIsPasswordInvalid(true);
      return InputError.TooShort;
    }
    if (password.length > MAX_LENGTH_PASSWORD) {
      setIsPasswordInvalid(true);
      return InputError.TooLong;
    }
    if (!password.match(/^\S+$/g)) {
      setIsPasswordInvalid(true);
      return InputError.Empty;
    }
    // ASCIIコード以外の文字が含まれている場合、
    if (!/^[\x00-\x7F]+$/.test(password)) {
      setIsPasswordInvalid(true);
      return InputError.InvalidCharacter;
    }

    setIsPasswordInvalid(false);
    return InputError.None;
  }

  function validatePasswordConfirm(): InputError {
    console.log("validating password confirmation...");

    if (!passwordConfirm) {
      setIsPasswordConfirmInvalid(true);
      return InputError.Empty;
    }
    if (password !== passwordConfirm) {
      setIsPasswordConfirmInvalid(true);
      return InputError.ConfirmationNotMatched;
    }

    setIsPasswordConfirmInvalid(false);
    return InputError.None;
  }

  // 戻り値
  //   true:  正常終了
  //   false: 入力値エラー有り
  function validateInputs(): boolean {
    const nicknameResult: InputError = validateNickname();
    const emailAddressResult: InputError = validateEmailAddress();
    const passwordResult: InputError = validatePassword();
    const passwordConfirmResult: InputError = validatePasswordConfirm();

    return (
      nicknameResult === InputError.None &&
      emailAddressResult === InputError.None &&
      passwordResult === InputError.None &&
      passwordConfirmResult === InputError.None
    );
  }

  function closePopup() {
    setSignupState(SignupState.Input);
  }

  const formStyle = css`
    display: flex;
    flex-direction: column;
    flex-wrap: nowrap;
    justify-content: flex-start;
    align-items: center;
  `;
  const inputItemStyle = css`
    margin-bottom: 0.5rem;
    width: 100%;
  `;
  const labelStyle = css`
    text-align: left;
  `;
  const requireStyle = css`
    color: red;
  `;
  const inputStyle = css`
    width: 100%;
    background-color: ${theme.palette.base.inputColor};
    color: ${theme.palette.secondary.fontColor};
  `;
  const invalidInputStyle = css`
    ${inputStyle}
    background-color: #ffd6d6;
  `;

  return (
    <>
      <Card topBorder>
        <div className={formStyle}>
          <div className={inputItemStyle}>
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
          </div>
          <div className={inputItemStyle}>
            <div className={labelStyle}>
              <span className={requireStyle}>*</span>メールアドレス
            </div>
            <div>
              <input
                className={
                  isEmailAddressInvalid ? invalidInputStyle : inputStyle
                }
                type="email"
                value={emailAddress}
                onChange={(e) => setEmailAddress(e.target.value)}
                placeholder={"例: example@slash-mochi.net"}
                maxLength={MAX_LENGTH_EMAILADDRESS}
                onBlur={() => validateEmailAddress()}
              />
            </div>
          </div>
          <div className={inputItemStyle}>
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
          </div>
          <div className={inputItemStyle}>
            <div className={labelStyle}>
              <span className={requireStyle}>*</span>登録するパスワード（確認）
            </div>
            <div>
              <input
                className={
                  isPasswordConfirmInvalid ? invalidInputStyle : inputStyle
                }
                type="password"
                value={passwordConfirm}
                onChange={(e) => setPasswordConfirm(e.target.value)}
                maxLength={MAX_LENGTH_PASSWORD}
                onBlur={() => validatePasswordConfirm()}
              />
            </div>
          </div>
        </div>
        <Button onClick={() => sendEmail()}>認証メールを送信する</Button>
      </Card>
      <Modal isOpen={signupState !== SignupState.Input}>
        {signupState === SignupState.SendingEmail && (
          <div>認証メールを送信しています。</div>
        )}
        {signupState === SignupState.SendedEmail && (
          <div>
            <p>認証メールを送信しました。</p>
            <p>
              3分以内にメール記載のリンクをクリックし、ユーザー登録を完了させてください。
            </p>
            <p>このページは閉じてかまいません。</p>
          </div>
        )}
        {signupState === SignupState.AlreadyRegistered && (
          <div>
            <div>メールアドレスは既に登録されています。</div>
            <div>
              <Button onClick={() => closePopup()}>OK</Button>
            </div>
          </div>
        )}
        {signupState === SignupState.FailedToSendEmail && (
          <div>
            <div>認証メールの送信に失敗しました。</div>
            <div>
              <Button onClick={() => closePopup()}>OK</Button>
            </div>
          </div>
        )}
      </Modal>
    </>
  );
}
