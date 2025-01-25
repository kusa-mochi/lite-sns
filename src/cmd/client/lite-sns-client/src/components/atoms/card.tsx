import { ReactNode } from "react"
import { useTheme } from "../../providers/themeProvider"
import { css } from "@emotion/css"

type Props = {
    children?: ReactNode
    topBorder?: boolean
}

export default function Card (props: Props) {
    const { children, topBorder } = props
    const theme = useTheme()

    const cardStyle = css`
        width: 100%;
        background-color: ${theme.palette.base.cardColor};
        padding: 2rem;
        ${topBorder ? `border-top: 1px solid ${theme.palette.base.borderColor};` : ""}
        border-bottom: 1px solid ${theme.palette.base.borderColor};
    `

    return (
        <div className={cardStyle}>
            {children}
        </div>
    )
}
