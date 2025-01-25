import { css } from "@emotion/css";
import Button from "../components/molecules/button";
import { useNavigate } from "react-router";
import { useTheme } from "../providers/themeProvider";
import Card from "../components/atoms/card";

export default function Home() {
    const navigate = useNavigate()
    const theme = useTheme()

    const appRootStyle = css`
        position: relative;
        color: ${theme.palette.secondary.fontColor};
        background-color: ${theme.palette.base.backgroundColor};
        width: 100%;
        height: 100%;

        display: flex;
        flex-direction: column;
        flex-wrap: nowrap;
        justify-content: flex-start;
        align-items: center;
    `
    const appBodyStyle = css`
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;

        width: 600px;
        height: 100%;
        background-color: ${theme.palette.base.bodyBackgroundColor};
    `

    const actionAreaStyle = css`
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: center;
        align-items: center;
        align-content: center;
    `

    const signupButtonContainerStyle = css`
        margin: 8px;
    `

    const signinButtonContainerStyle = css`
        margin: 8px;
    `

    return (
        <div className={appRootStyle}>
            <div className={appBodyStyle}>
                <div>
                    Go to <a href="/test">Test Page</a>
                </div>
                <div>
                    Lite SNS
                </div>
                <Card topBorder>
                    <div className={actionAreaStyle}>
                        <div className={signupButtonContainerStyle}>
                            <Button onClick={() => {navigate("/signup")}} width="80px">Sign up</Button>
                        </div>
                        <div className={signinButtonContainerStyle}>
                            <Button onClick={() => {navigate("/signin")}} width="80px">Sign in</Button>
                        </div>
                    </div>
                </Card>
            </div>
        </div>
    )
}
