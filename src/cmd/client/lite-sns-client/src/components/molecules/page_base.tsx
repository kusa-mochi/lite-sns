import { ReactNode } from "react"
import { useTheme } from "../../providers/themeProvider"
import { css } from "@emotion/css"

type Props = {
    children?: ReactNode
}

export default function PageBase(props: Props) {
    const { children } = props
    const theme = useTheme()

    const pageRootStyle = css`
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
    const pageBodyStyle = css`
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;

        width: 600px;
        height: 100%;
        background-color: ${theme.palette.base.bodyBackgroundColor};
    `

    return (
        <div className={pageRootStyle}>
            <div className={pageBodyStyle}>
                {children}
            </div>
        </div>
    )
}
