import { css } from "@emotion/css"

export default function Test() {
    const theme = useTheme("light")
    const testStyle = css`
        color: ${theme?.palette.primary.fontColor};
        background-color: ${theme?.palette.primary.main};
    `
    return (
        <>
            <div className={testStyle}>This is a test component.</div>
            <div>
                <p>Go to <a href="/test2">Test2 Page</a></p>
            </div>
        </>
    )
}
