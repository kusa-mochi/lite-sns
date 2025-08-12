import { MouseEvent, ReactNode, useEffect, useState } from "react"
import { useTheme } from "../../providers/themeProvider"
import { css } from "@emotion/css"

type Props = {
    children?: ReactNode
    active?: boolean
    disabled?: boolean
    width?: string
    height?: string
    onClick?: (e: MouseEvent) => void
}

export default function MenuButton (props: Props) {
    const { children, active, disabled, width, height, onClick } = props
    const theme = useTheme()
    const [activeState, setActiveState] = useState(false)
    const [fColor, setFColor] = useState(theme.palette.secondary.fontColor)
    const [fSize, setFSize] = useState(1)

    useEffect(() => {
        setActiveState(active === undefined ? false : active)
    }, [])

    useEffect(() => {
        setActiveState(active === undefined ? false : active)
    }, [active])

    const buttonStyle = css`
        background-color: transparent;
        color: ${fColor};
        box-shadow: none;
        border: none;
        border-radius: 0;
        ${width === undefined ? "" : `width: ${width};`}
        ${height === undefined ? "" : `height: ${height};`}
        font-size: ${fSize}rem;
        padding: ${fSize/3}rem;
        cursor: ${disabled === true ? "default" : "pointer"};
        opacity: ${disabled === undefined || disabled === false ? "1" : "0.4"};
        outline: none;

        &:hover {
            filter: ${disabled === true ? "none" : "brightness(90%)"};
        }

        filter: ${activeState && disabled === false ? "brightness(80%)" : "none"};
    `

    return (
        <>
            <button
                onClick={onClick ? onClick : () => {}}
                className={buttonStyle}
                disabled={disabled}
            >
                {children}
            </button>
        </>
    )
}
