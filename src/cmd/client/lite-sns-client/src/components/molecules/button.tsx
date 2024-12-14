import { css } from "@emotion/css"
import { MouseEvent, ReactNode, useEffect, useState } from "react"
import { useTheme } from "../../providers/themeProvider"

type Props = {
    children?: ReactNode
    active?: boolean
    enabled?: boolean
    focused?: boolean
    primary?: boolean
    secondary?: boolean
    onClick?: (e: MouseEvent) => void
}

export default function Button (props: Props) {
    const theme = useTheme()
    const { children, active, enabled, focused, primary, secondary, onClick } = props
    const [activeState, setActiveState] = useState(false)
    const [bgColor, setBgColor] = useState(theme.palette.secondary.main)
    const [fColor, setFColor] = useState(theme.palette.secondary.fontColor)
    const [outlineColor, setOutlineColor] = useState(theme.palette.secondary.outlineColor)
    const [fSize, setFSize] = useState(1)

    useEffect(() => {
        setActiveState(active === undefined ? false : active)
        if (secondary) {
            setBgColor(theme.palette.secondary.main)
            setFColor(theme.palette.secondary.fontColor)
            setOutlineColor(theme.palette.secondary.outlineColor)
        }
        if (primary) {
            setBgColor(theme.palette.primary.main)
            setFColor(theme.palette.primary.fontColor)
            setOutlineColor(theme.palette.primary.outlineColor)
        }
        setFSize(theme.typography.fontSize)
    }, [])

    useEffect(() => {
        setActiveState(active === undefined ? false : active)
        if (secondary) {
            setBgColor(theme.palette.secondary.main)
            setFColor(theme.palette.secondary.fontColor)
            setOutlineColor(theme.palette.secondary.outlineColor)
        }
        if (primary) {
            setBgColor(theme.palette.primary.main)
            setFColor(theme.palette.primary.fontColor)
            setOutlineColor(theme.palette.primary.outlineColor)
        }
        setFSize(theme.typography.fontSize)
    }, [active, primary, secondary])

    const buttonStyle = css`
        background-color: ${bgColor};
        color: ${fColor};
        box-shadow: rgba(0, 0, 0, 0.2) 0px 3px 1px -2px, rgba(0, 0, 0, 0.14) 0px 2px 2px 0px, rgba(0, 0, 0, 0.12) 0px 1px 5px 0px;
        font-size: ${fSize}rem;
        padding: ${fSize/3}rem;
        cursor: ${enabled === false ? "default" : "pointer"};
        opacity: ${enabled === undefined || enabled === true ? "1" : "0.4"};
        outline-color: ${outlineColor};
        outline-style: solid;
        outline-width: ${focused === true ? "midium" : "0"};

        &:hover {
            filter: ${enabled === false ? "none" : "brightness(90%)"};
        }
        
        filter: ${activeState && enabled !== false ? "brightness(80%)" : "none"};
    `

    return (
        <>
            <div
                onClick={onClick ? onClick : () => {}}
                className={buttonStyle}
            >
                {children}
            </div>
        </>
    )
}
