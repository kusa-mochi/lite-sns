import { css } from "@emotion/css"
import { ReactNode, useEffect, useState } from "react"
import { useTheme } from "../../providers/themeProvider"

type Props = {
    children?: ReactNode
    primary?: boolean
    secondary?: boolean
    onClick?: () => void
}

export default function Button (props: Props) {
    const theme = useTheme()
    const { children, primary, secondary, onClick } = props
    const [bgColor, setBgColor] = useState("")
    const [fontColor, setFontColor] = useState("")

    useEffect(() => {
        // defult colors
        setBgColor(theme.palette.secondary.main)
        setFontColor(theme.palette.secondary.fontColor)

        if (secondary) {
            setBgColor(theme.palette.secondary.main)
            setFontColor(theme.palette.secondary.fontColor)
        }
        if (primary) {
            setBgColor(theme.palette.primary.main)
            setFontColor(theme.palette.primary.fontColor)
        }
    }, [])

    useEffect(() => {
        // defult colors
        setBgColor(theme.palette.secondary.main)
        setFontColor(theme.palette.secondary.fontColor)
        
        if (secondary) {
            setBgColor(theme.palette.secondary.main)
            setFontColor(theme.palette.secondary.fontColor)
        }
        if (primary) {
            setBgColor(theme.palette.primary.main)
            setFontColor(theme.palette.primary.fontColor)
        }
    }, [primary, secondary])

    const buttonStyle = css`
        background-color: ${bgColor};
        color: ${fontColor};
        box-shadow: rgba(0, 0, 0, 0.2) 0px 3px 1px -2px, rgba(0, 0, 0, 0.14) 0px 2px 2px 0px, rgba(0, 0, 0, 0.12) 0px 1px 5px 0px;
    `
    console.log(`bg:${bgColor}, color:${fontColor}`)

    return (
        <div
            onClick={onClick ? () => onClick() : () => {}}
            className={buttonStyle}
        >
            {children}
        </div>
    )
}
