import { css } from "@emotion/css"
import { ReactNode, useEffect, useState } from "react"
import { useTheme } from "../../providers/themeProvider"

type Props = {
    children?: ReactNode
    enabled?: boolean
    primary?: boolean
    secondary?: boolean
    onClick?: () => void
}

export default function Button (props: Props) {
    const theme = useTheme()
    const { children, enabled, primary, secondary, onClick } = props
    const [bgColor, setBgColor] = useState(theme.palette.secondary.main)
    const [fColor, setFColor] = useState(theme.palette.secondary.fontColor)
    const [fSize, setFSize] = useState(1)

    useEffect(() => {
        if (secondary) {
            setBgColor(theme.palette.secondary.main)
            setFColor(theme.palette.secondary.fontColor)
        }
        if (primary) {
            setBgColor(theme.palette.primary.main)
            setFColor(theme.palette.primary.fontColor)
        }
        setFSize(theme.typography.fontSize)
    }, [])

    useEffect(() => {
        if (secondary) {
            setBgColor(theme.palette.secondary.main)
            setFColor(theme.palette.secondary.fontColor)
        }
        if (primary) {
            setBgColor(theme.palette.primary.main)
            setFColor(theme.palette.primary.fontColor)
        }
        setFSize(theme.typography.fontSize)
    }, [primary, secondary])

    const buttonStyle = css`
        background-color: ${bgColor};
        color: ${fColor};
        box-shadow: rgba(0, 0, 0, 0.2) 0px 3px 1px -2px, rgba(0, 0, 0, 0.14) 0px 2px 2px 0px, rgba(0, 0, 0, 0.12) 0px 1px 5px 0px;
        font-size: ${fSize}rem;
        padding: ${fSize/3}rem;
        cursor: ${enabled === false ? "default" : "pointer"};
        opacity: ${enabled === undefined || enabled === true ? "1" : "0.4"};

        &:hover {
            filter: ${enabled === false ? "none" : "brightness(90%)"};
        }
    `

    return (
        <>
            <div
                onClick={onClick ? () => onClick() : () => {}}
                className={buttonStyle}
            >
                {children}
            </div>
        </>
    )
}
